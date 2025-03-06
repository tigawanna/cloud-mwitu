package controller

import (
    "net/http"
    "time"

    "github.com/go-fuego/fuego"
    "github.com/go-fuego/fuego/option"
    "github.com/tigawanna/cloud-mwitu/internal/services"
)

// SessionCookieName is the name of the cookie that stores the session ID
const SessionCookieName = "cloud_mwitu_session"

// AuthResources handles authentication-related endpoints
type AuthResources struct {
    AuthService services.AuthService
}

// LoginRequest represents login credentials
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// RegisterRequest represents registration data
type RegisterRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// UserResponse is the sanitized user data for responses
type UserResponse struct {
    ID        uint      `json:"id"`
    Username  string    `json:"username"`
    CreatedAt time.Time `json:"created_at"`
    LastLogin time.Time `json:"last_login,omitempty"`
    IsActive  bool      `json:"is_active"`
}

// Routes registers all authentication routes
func (rs AuthResources) Routes(s *fuego.Server) {
    authGroup := fuego.Group(s, "/auth")

    fuego.Post(authGroup, "/register", rs.registerUser,
        option.DefaultStatusCode(201),
        option.Description("Register a new user account"),
    )

    fuego.Post(authGroup, "/login", rs.loginUser,
        option.Description("Login with username and password"),
    )

    fuego.Post(authGroup, "/logout", rs.logoutUser,
        option.Description("Logout current user"),
    )

    fuego.Get(authGroup, "/profile", rs.getUserProfile,
        option.Description("Get current user profile"),
    )

    fuego.Post(authGroup, "/refresh", rs.refreshSession,
        option.Description("Refresh the current session"),
    )
}

// registerUser handles user registration
func (rs AuthResources) registerUser(c fuego.ContextWithBody[RegisterRequest]) (UserResponse, error) {
    req, err := c.Body()
    if err != nil {
        return UserResponse{}, fuego.BadRequestError{
            Title:  "Invalid Request",
            Detail: "Could not parse request body",
            Err:    err,
        }
    }

    // Validate input
    if req.Username == "" || req.Password == "" {
        return UserResponse{}, fuego.BadRequestError{
            Title:  "Validation Error",
            Detail: "Username and password are required",
        }
    }

    if len(req.Password) < 8 {
        return UserResponse{}, fuego.BadRequestError{
            Title:  "Validation Error",
            Detail: "Password must be at least 8 characters",
        }
    }

    // Register the user
    user, err := rs.AuthService.RegisterUser(req.Username, req.Password)
    if err != nil {
        if err == services.ErrUserExists {
            return UserResponse{}, fuego.ConflictError{
                Title:  "Username Taken",
                Detail: "This username is already in use",
                Err:    err,
            }
        }
        return UserResponse{}, fuego.InternalServerError{
            Title:  "Registration Failed",
            Detail: "Could not complete registration",
            Err:    err,
        }
    }

    // Convert to sanitized response
    response := UserResponse{
        ID:        user.ID,
        Username:  user.Username,
        CreatedAt: user.CreatedAt,
        IsActive:  user.IsActive,
    }

    return response, nil
}

// loginUser handles user login
func (rs AuthResources) loginUser(c fuego.ContextWithBody[LoginRequest]) (map[string]string, error) {
    req, err := c.Body()
    if err != nil {
        return nil, fuego.BadRequestError{
            Title:  "Invalid Request",
            Detail: "Could not parse request body",
            Err:    err,
        }
    }

    // Validate input
    if req.Username == "" || req.Password == "" {
        return nil, fuego.BadRequestError{
            Title:  "Validation Error",
            Detail: "Username and password are required",
        }
    }

    // Extract client information
    userAgent := c.Request().UserAgent()
    ipAddress := c.Request().RemoteAddr

    // Attempt login
    session, err := rs.AuthService.Login(req.Username, req.Password, userAgent, ipAddress)
    if err != nil {
        if err == services.ErrInvalidCredentials {
            return nil, fuego.UnauthorizedError{
                Title:  "Authentication Failed",
                Detail: "Invalid username or password",
                Err:    err,
            }
        }
        return nil, fuego.InternalServerError{
            Title:  "Login Failed",
            Detail: "Could not process login request",
            Err:    err,
        }
    }

    // Set session cookie
    http.SetCookie(c.Response(), &http.Cookie{
        Name:     SessionCookieName,
        Value:    session.ID,
        Path:     "/",
        Expires:  session.ExpiresAt,
        HttpOnly: true,
        SameSite: http.SameSiteStrictMode,
        Secure:   false, // Set to true in production with HTTPS
    })

    return map[string]string{
        "message": "Login successful",
    }, nil
}

// logoutUser handles user logout
func (rs AuthResources) logoutUser(c fuego.ContextNoBody) (map[string]string, error) {
    // Get the session cookie
    cookie, err := c.Request().Cookie(SessionCookieName)
    if err != nil {
        return map[string]string{
            "message": "Already logged out",
        }, nil
    }

    // Invalidate the session
    sessionID := cookie.Value
    if err := rs.AuthService.Logout(sessionID); err != nil {
        return nil, fuego.InternalServerError{
            Title:  "Logout Failed",
            Detail: "Could not process logout request",
            Err:    err,
        }
    }

    // Clear the session cookie
    http.SetCookie(c.Response(), &http.Cookie{
        Name:     SessionCookieName,
        Value:    "",
        Path:     "/",
        MaxAge:   -1,
        HttpOnly: true,
        SameSite: http.SameSiteStrictMode,
    })

    return map[string]string{
        "message": "Logout successful",
    }, nil
}

// getUserProfile gets the current user's profile
func (rs AuthResources) getUserProfile(c fuego.ContextNoBody) (UserResponse, error) {
    // Get the session cookie
    cookie, err := c.Request().Cookie(SessionCookieName)
    if err != nil {
        return UserResponse{}, fuego.UnauthorizedError{
            Title:  "Not Authenticated",
            Detail: "Authentication required to access this resource",
            Err:    err,
        }
    }

    // Validate session and get user
    user, err := rs.AuthService.GetUserFromSession(cookie.Value)
    if err != nil {
        if err == services.ErrInvalidSession || err == services.ErrSessionExpired {
            return UserResponse{}, fuego.UnauthorizedError{
                Title:  "Invalid Session",
                Detail: "Your session has expired or is invalid",
                Err:    err,
            }
        }
        return UserResponse{}, fuego.InternalServerError{
            Title:  "Profile Retrieval Failed",
            Detail: "Could not retrieve user profile",
            Err:    err,
        }
    }

    // Convert to sanitized response
    response := UserResponse{
        ID:        user.ID,
        Username:  user.Username,
        CreatedAt: user.CreatedAt,
        LastLogin: user.LastLogin,
        IsActive:  user.IsActive,
    }

    return response, nil
}

// refreshSession extends the current session
func (rs AuthResources) refreshSession(c fuego.ContextNoBody) (map[string]string, error) {
    // Get the session cookie
    cookie, err := c.Request().Cookie(SessionCookieName)
    if err != nil {
        return nil, fuego.UnauthorizedError{
            Title:  "Not Authenticated",
            Detail: "Authentication required to refresh session",
            Err:    err,
        }
    }

    // Refresh the session
    sessionID := cookie.Value
    session, err := rs.AuthService.RefreshSession(sessionID)
    if err != nil {
        if err == services.ErrInvalidSession || err == services.ErrSessionExpired {
            return nil, fuego.UnauthorizedError{
                Title:  "Invalid Session",
                Detail: "Your session has expired or is invalid",
                Err:    err,
            }
        }
        return nil, fuego.InternalServerError{
            Title:  "Session Refresh Failed",
            Detail: "Could not refresh your session",
            Err:    err,
        }
    }

    // Update the cookie
    http.SetCookie(c.Response(), &http.Cookie{
        Name:     SessionCookieName,
        Value:    session.ID,
        Path:     "/",
        Expires:  session.ExpiresAt,
        HttpOnly: true,
        SameSite: http.SameSiteStrictMode,
        Secure:   false, // Set to true in production with HTTPS
    })

    return map[string]string{
        "message": "Session refreshed",
    }, nil
}
