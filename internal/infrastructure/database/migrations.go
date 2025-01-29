package database

import (
	roomEntity "github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	"gorm.io/gorm"
)

import (
	hotelEntity "github.com/gatimugabriel/hotel-reservation-system/internal/domain/hotel/entity"
	reservationEntity "github.com/gatimugabriel/hotel-reservation-system/internal/domain/reservation/entity"
	userEntity "github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/entity"
)

// RunMigrations performs auto-migration for all models
func RunMigrations(db *gorm.DB) error {
	// Enable uuid-ossp extension for UUID support
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "citext";`) // allow case insensitive

	// Auto-migrate all models
	return db.AutoMigrate(
		&userEntity.User{},
		&hotelEntity.Hotel{},
		&roomEntity.RoomType{},
		&roomEntity.Room{},
		&reservationEntity.Reservation{},
	)
}