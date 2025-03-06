package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/tigawanna/cloud-mwitu/internal/configs"
	// "github.com/tigawanna/cloud-mwitu/internal/models"
	"github.com/tigawanna/cloud-mwitu/internal/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DbInstance holds the database connection
type DbInstance struct {
    Db *gorm.DB
}

// Database is a singleton instance
var (
    Database DbInstance
    once     sync.Once
)

// Config holds database configuration
type Config struct {
    DSN      string // Data Source Name
    Debug    bool   // Enable debug mode
    InMemory bool   // Use in-memory database
}

// Initialize initializes the database connection with the given config
func Initialize(cfg Config) *gorm.DB {
    once.Do(func() {
        // Determine the DSN
        dsn := cfg.DSN
        if dsn == "" {
            dsn = "file:mwitu.db?cache=shared"
        }
        if cfg.InMemory {
            dsn = "file::memory:?cache=shared"
        }
        
        // Configure logger
        logLevel := logger.Silent
        if cfg.Debug {
            logLevel = logger.Info
        }
        
        gormConfig := &gorm.Config{
            Logger: logger.Default.LogMode(logLevel),
        }
        
        // Open database connection
        db, err := gorm.Open(sqlite.Open(dsn), gormConfig)
        if err != nil {
            log.Fatalf("Failed to connect to database: %v", err)
        }
        
        log.Println("Connected to the database")
        
        // Store in singleton
        Database = DbInstance{Db: db}
    })
    
    return Database.Db
}

// GetDB returns the database instance, initializing with default config if not already initialized
func GetDB() *gorm.DB {
    if Database.Db == nil {
        return Initialize(Config{
            DSN:   "file:mwitu.db?cache=shared",
            Debug: true,
        })
    }
    return Database.Db
}

// Close closes the database connection (SQLite doesn't actually need this, but good practice)
func Close() error {
    if Database.Db == nil {
        return nil
    }
    
    sqlDB, err := Database.Db.DB()
    if err != nil {
        return err
    }
    return sqlDB.Close()
}

// AutoMigrate creates or updates database tables based on models
func AutoMigrate(db *gorm.DB, models ...interface{}) error {
    if err := db.AutoMigrate(models...); err != nil {
        return fmt.Errorf("database migration failed: %w", err)
    }
    log.Println("Database migrations completed successfully")
    return nil
}

func ensureDataDirectory() {
    dataDir := filepath.Dir(configs.GetEnv().SQLitePath)
    if err := os.MkdirAll(dataDir, 0755); err != nil {
        log.Fatalf("Failed to create data directory: %v", err)
    }
}
// InitDB initializes the database and runs migrations
func InitDB() *gorm.DB {
	ensureDataDirectory()
    // Get database configuration
    config := Config{
        DSN:     configs.GetEnv().DatabaseDSN,
        Debug:   configs.GetEnv().Debug,
        InMemory: false,
    }
    
    // Initialize database
    database := Initialize(config)
    
    // Register models and run migrations
    // err := AutoMigrate(database, 
    //     &models.User{},
    //     &models.Session{},
    //     &models.CaddyFileHistory{},
    //     &models.SystemDFileHistory{},
    //     // Add other models here
    // )
    
    // if err != nil {
    //     log.Fatalf("Database migration failed: %v", err)
    // }
    
    // Start a background task to clean expired sessions
    go cleanExpiredSessions(database)
    
    return database
}

// cleanExpiredSessions periodically removes expired sessions
func cleanExpiredSessions(database *gorm.DB) {
    ticker := time.NewTicker(24 * time.Hour) // Run once per day
    for range ticker.C {
        authService := services.NewAuthService(database, nil)
        if err := authService.CleanExpiredSessions(); err != nil {
            log.Printf("Error cleaning expired sessions: %v", err)
        }
    }
}
