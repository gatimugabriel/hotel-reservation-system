package database

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/hotel/entity"
	roomEntity "github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	userEntity "github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/entity"
	"gorm.io/gorm"
)

// RunMigrations performs auto-migration for all models
func RunMigrations(db *gorm.DB) error {
	// Enable uuid-ossp extension for UUID support
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	// Auto-migrate all models
	return db.AutoMigrate(
		&userEntity.User{},
		&entity.Hotel{},
		&roomEntity.RoomType{},
		&roomEntity.Room{},
	)
}