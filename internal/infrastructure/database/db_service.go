package database

import (
	"context"
	"fmt"
	"github.com/gatimugabriel/hotel-reservation-system/internal/config"
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Service provides database operations
type Service struct {
	DB *gorm.DB
}

var (
	dbInstance      *Service
	once            sync.Once
	databaseInitErr error
)

// NewDatabaseService creates and configures a new database connection
func NewDatabaseService(cfg *config.Config) (*Service, error) {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
			cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port)

		gormConfig := &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}

		var db *gorm.DB
		var connectionError error

		// Retry connection with backoff
		maxRetries := 5
		for i := 0; i < maxRetries; i++ {
			db, connectionError = gorm.Open(postgres.Open(dsn), gormConfig)
			if connectionError == nil {
				break
			}
			log.Printf("Database connection attempt %d failed: %v", i+1, connectionError)
			time.Sleep(time.Second * time.Duration(i+1))
		}

		if connectionError != nil {
			databaseInitErr = fmt.Errorf("failed to connect to database after %d retries: %v", maxRetries, connectionError)
			return
		}

		// connection pool
		sqlDB, err := db.DB()
		if err != nil {
			databaseInitErr = fmt.Errorf("failed to get database connection pool: %w", err)
			return
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
		sqlDB.SetConnMaxIdleTime(time.Minute * 30)

		// Run migrations
		if migrationErr := RunMigrations(db); migrationErr != nil {
			databaseInitErr = fmt.Errorf("database migration failed: %w", migrationErr)
			return
		}
		log.Println("Database migration completed successfully")

		dbInstance = &Service{DB: db}
	})

	if databaseInitErr != nil {
		return nil, databaseInitErr
	}

	return dbInstance, nil
}

// WithContext returns a new DB instance with context
func (s *Service) WithContext(ctx context.Context) *gorm.DB {
	return s.DB.WithContext(ctx)
}

// Transaction executes operations within a transaction
func (s *Service) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return s.DB.WithContext(ctx).Transaction(fn)
}