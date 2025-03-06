package models

import "time"


type CaddyFileHistory struct {
    ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
    Content      string    `gorm:"type:text" json:"content"`
    UpdatedBlock string    `gorm:"type:text" json:"updatedBlock"`
    User         string    `gorm:"" json:"user,omitempty"` // Optional: track who made the change
    Reason       string    `gorm:"" json:"reason,omitempty"` // Optional: track why the change was made
}

// TableName overrides the table name used by GORM
func (CaddyFileHistory) TableName() string {
    return "caddy_file_histories" // customize the table name
}

type SystemDFileHistory struct {
    ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
    ServiceName  string    `gorm:"type:varchar(100);index" json:"serviceName"`
    Content      string    `gorm:"type:text" json:"content"`
    UpdatedBlock string    `gorm:"type:text" json:"updatedBlock"`
    UnitType     string    `gorm:"type:varchar(20)" json:"unitType"` // service, socket, timer, etc.
    User         string    `gorm:"" json:"user,omitempty"`
    Reason       string    `gorm:"" json:"reason,omitempty"`
    Status       string    `gorm:"type:varchar(20)" json:"status,omitempty"` // active, inactive, modified
    FilePath     string    `gorm:"type:varchar(255)" json:"filePath"`
}

// TableName overrides the table name used by GORM
func (SystemDFileHistory) TableName() string {
    return "systemd_file_histories"
}
