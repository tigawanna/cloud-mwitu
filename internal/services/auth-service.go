package services

import (
    "time"

    "github.com/tigawanna/cloud-mwitu/internal/models"
)

// AuthService interface defines authentication-related operations
type AuthService interface {
    // User account management
    RegisterUser(username, password string) (*models.User, error)
    GetUserByID(id uint) (*models.User, error)
    GetUserByUsername(username string) (*models.User, error)
    UpdateUser(user *models.User) error
    DeleteUser(id uint) error

    // Authentication operations
    Login(username, password, userAgent, ipAddress string) (*models.Session, error)
    Logout(sessionID string) error
    ValidateSession(sessionID string) (*models.Session, error)
    RefreshSession(sessionID string) (*models.Session, error)
    
    // Session management
    GetSessionByID(sessionID string) (*models.Session, error)
    GetUserFromSession(sessionID string) (*models.User, error)
    InvalidateAllUserSessions(userID uint) error
    CleanExpiredSessions() error
}

// SessionConfig holds configuration for sessions
type SessionConfig struct {
    CookieName      string
    CookiePath      string
    SessionDuration time.Duration
    SecureCookie    bool
    SameSitePolicy  string
}

// DefaultSessionConfig returns the default session configuration
func DefaultSessionConfig() SessionConfig {
    return SessionConfig{
        CookieName:      "cloud_mwitu_session",
        CookiePath:      "/",
        SessionDuration: 7 * 24 * time.Hour, // 1 week
        SecureCookie:    false,              // Set to true in production with HTTPS
        SameSitePolicy:  "strict",
    }
}
