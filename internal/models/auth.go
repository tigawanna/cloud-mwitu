package models

import (
    "time"

    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

// User represents a user account in the system
type User struct {
    ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
    Username     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
    PasswordHash string    `gorm:"type:varchar(255);not null" json:"-"` // Never expose in JSON
    LastLogin    time.Time `gorm:"" json:"last_login,omitempty"`
    IsActive     bool      `gorm:"default:true" json:"is_active"`
    Sessions     []Session `gorm:"foreignKey:UserID" json:"-"` // One-to-many relationship
}

// Session represents an authentication session
type Session struct {
    ID           string    `gorm:"primaryKey;type:varchar(64)" json:"id"`
    UserID       uint      `gorm:"index;not null" json:"user_id"`
    CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
    ExpiresAt    time.Time `gorm:"not null" json:"expires_at"`
    LastActivity time.Time `gorm:"autoUpdateTime" json:"last_activity"`
    UserAgent    string    `gorm:"type:varchar(255)" json:"user_agent"`
    IPAddress    string    `gorm:"type:varchar(45)" json:"ip_address"` // IPv6 can be up to 45 chars
    IsValid      bool      `gorm:"default:true" json:"is_valid"`
}

// BeforeSave hook for User to hash the password before saving
func (u *User) BeforeSave(tx *gorm.DB) error {
    // Only hash the password if it has been changed
    if u.PasswordHash != "" && len(u.PasswordHash) < 60 {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
        if err != nil {
            return err
        }
        u.PasswordHash = string(hashedPassword)
    }
    return nil
}

// CheckPassword compares a password against the hashed password
func (u *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
    return err == nil
}

// LoginRequest represents a user login request
type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

// RegisterRequest represents a user registration request
type RegisterRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required,min=8"`
}
