package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/tigawanna/cloud-mwitu/internal/configs"
	"github.com/tigawanna/cloud-mwitu/internal/controllers"
	"github.com/tigawanna/cloud-mwitu/internal/services"
)

// ErrorResponse represents a structured error response
type ErrorResponse struct {
    Status  int    `json:"status"`
    Title   string `json:"title"`
    Message string `json:"message"`
    Error   string `json:"error,omitempty"`
}

// AuthMiddleware creates middleware that enforces authentication
func AuthMiddleware(authService services.AuthService) func(next http.Handler) http.Handler {
    excludeList := []string{"/swagger/", "/caddy/", "/api/auth/login", "/api/auth/register"}
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Exclude endpoints
            for _, v := range excludeList {
                if strings.Contains(r.URL.Path, v) {
                    next.ServeHTTP(w, r)
                    return
                }
            }
            
            // Get session cookie
            cookie, err := r.Cookie(controller.SessionCookieName)
            if err != nil {
                sendUnauthorizedResponse(w, "No valid session found", err.Error())
                return
            }

            sessionID := cookie.Value
            
            // Validate the session
            _, err = authService.ValidateSession(sessionID)
            if err != nil {
                sendUnauthorizedResponse(w, "Session invalid or expired", err.Error())
                return
            }

            // Call the next handler
            next.ServeHTTP(w, r)
        })
    }
}

// sendUnauthorizedResponse sends a structured JSON error response
func sendUnauthorizedResponse(w http.ResponseWriter, message string, errDetails string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusUnauthorized)
    
    response := ErrorResponse{
        Status:  http.StatusUnauthorized,
        Title:   "Unauthorized",
        Message: message,
    }
    
    // Only include error details in non-production environments
    env := configs.Env{}
    if errDetails != "" && env.Debug {
        response.Error = errDetails
    }
    
    json.NewEncoder(w).Encode(response)
}
