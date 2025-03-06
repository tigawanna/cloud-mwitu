package services

import (
    "errors"
    "time"

    "github.com/google/uuid"
    "github.com/tigawanna/cloud-mwitu/internal/models"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

// Common errors
var (
    ErrInvalidCredentials = errors.New("invalid username or password")
    ErrUserExists         = errors.New("user already exists")
    ErrInvalidSession     = errors.New("invalid session")
    ErrSessionExpired     = errors.New("session expired")
    ErrUserNotFound       = errors.New("user not found")
)

// AuthServiceImpl implements the AuthService interface
type AuthServiceImpl struct {
    db            *gorm.DB
    sessionConfig SessionConfig
}

// NewAuthService creates a new authentication service
func NewAuthService(db *gorm.DB, config *SessionConfig) AuthService {
    // Use default config if none provided
    cfg := DefaultSessionConfig()
    if config != nil {
        cfg = *config
    }

    return &AuthServiceImpl{
        db:            db,
        sessionConfig: cfg,
    }
}

// RegisterUser creates a new user account
func (s *AuthServiceImpl) RegisterUser(username, password string) (*models.User, error) {
    // Check if user already exists
    var existingUser models.User
    result := s.db.Where("username = ?", username).First(&existingUser)
    if result.RowsAffected > 0 {
        return nil, ErrUserExists
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    // Create new user
    user := &models.User{
        Username:     username,
        PasswordHash: string(hashedPassword),
        IsActive:     true,
    }

    if err := s.db.Create(user).Error; err != nil {
        return nil, err
    }

    return user, nil
}

// GetUserByID retrieves a user by ID
func (s *AuthServiceImpl) GetUserByID(id uint) (*models.User, error) {
    var user models.User
    if err := s.db.First(&user, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrUserNotFound
        }
        return nil, err
    }
    return &user, nil
}

// GetUserByUsername retrieves a user by username
func (s *AuthServiceImpl) GetUserByUsername(username string) (*models.User, error) {
    var user models.User
    if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrUserNotFound
        }
        return nil, err
    }
    return &user, nil
}

// UpdateUser updates a user's information
func (s *AuthServiceImpl) UpdateUser(user *models.User) error {
    return s.db.Save(user).Error
}

// DeleteUser deletes a user account
func (s *AuthServiceImpl) DeleteUser(id uint) error {
    return s.db.Delete(&models.User{}, id).Error
}

// Login authenticates a user and creates a new session
func (s *AuthServiceImpl) Login(username, password, userAgent, ipAddress string) (*models.Session, error) {
    user, err := s.GetUserByUsername(username)
    if err != nil {
        return nil, ErrInvalidCredentials
    }

    // Check password
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return nil, ErrInvalidCredentials
    }

    // Update last login time
    user.LastLogin = time.Now()
    s.db.Save(user)

    // Create new session
    sessionID := uuid.New().String()
    session := &models.Session{
        ID:           sessionID,
        UserID:       user.ID,
        CreatedAt:    time.Now(),
        ExpiresAt:    time.Now().Add(s.sessionConfig.SessionDuration),
        LastActivity: time.Now(),
        UserAgent:    userAgent,
        IPAddress:    ipAddress,
        IsValid:      true,
    }

    if err := s.db.Create(session).Error; err != nil {
        return nil, err
    }

    return session, nil
}

// Logout invalidates a user session
func (s *AuthServiceImpl) Logout(sessionID string) error {
    result := s.db.Model(&models.Session{}).
        Where("id = ?", sessionID).
        Update("is_valid", false)

    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return ErrInvalidSession
    }

    return nil
}

// ValidateSession checks if a session is valid and not expired
func (s *AuthServiceImpl) ValidateSession(sessionID string) (*models.Session, error) {
    var session models.Session
    if err := s.db.Where("id = ? AND is_valid = ?", sessionID, true).First(&session).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrInvalidSession
        }
        return nil, err
    }

    // Check expiration
    if time.Now().After(session.ExpiresAt) {
        s.db.Model(&session).Update("is_valid", false)
        return nil, ErrSessionExpired
    }

    // Update last activity
    s.db.Model(&session).Update("last_activity", time.Now())
    
    // Update the struct as well to reflect DB changes
    session.LastActivity = time.Now()
    
    return &session, nil
}

// RefreshSession extends the expiration time of a session
func (s *AuthServiceImpl) RefreshSession(sessionID string) (*models.Session, error) {
    session, err := s.ValidateSession(sessionID)
    if err != nil {
        return nil, err
    }

    // Extend expiration time
    session.ExpiresAt = time.Now().Add(s.sessionConfig.SessionDuration)
    if err := s.db.Save(session).Error; err != nil {
        return nil, err
    }

    return session, nil
}

// GetSessionByID retrieves a session by ID
func (s *AuthServiceImpl) GetSessionByID(sessionID string) (*models.Session, error) {
    var session models.Session
    if err := s.db.Where("id = ?", sessionID).First(&session).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrInvalidSession
        }
        return nil, err
    }
    return &session, nil
}

// GetUserFromSession retrieves the user associated with a session
func (s *AuthServiceImpl) GetUserFromSession(sessionID string) (*models.User, error) {
    session, err := s.ValidateSession(sessionID)
    if err != nil {
        return nil, err
    }

    return s.GetUserByID(session.UserID)
}

// InvalidateAllUserSessions invalidates all sessions for a specific user
func (s *AuthServiceImpl) InvalidateAllUserSessions(userID uint) error {
    return s.db.Model(&models.Session{}).
        Where("user_id = ?", userID).
        Update("is_valid", false).
        Error
}

// CleanExpiredSessions removes expired sessions from the database
func (s *AuthServiceImpl) CleanExpiredSessions() error {
    return s.db.Where("expires_at < ? OR is_valid = ?", time.Now(), false).
        Delete(&models.Session{}).
        Error
}
