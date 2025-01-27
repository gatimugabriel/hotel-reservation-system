package entity

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/hotel/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RoomType : represents a category of rooms with similar characteristics
type RoomType struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name         string         `gorm:"type:varchar(100);not null;unique"`
	Description  string         `gorm:"type:text"`
	BasePrice    float64        `gorm:"type:decimal(10,2);not null"`
	MaxOccupancy int            `gorm:"not null"`
	NumBeds      int            `gorm:"not null"`
	BedType      string         `gorm:"type:varchar(50);not null"`
	SquareMeters float64        `gorm:"type:decimal(6,2);not null"`
	Rooms        []Room         `gorm:"foreignKey:RoomTypeID"`
	CreatedAt    time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// Room : represents an individual room in a hotel
type Room struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	RoomNumber    string         `gorm:"type:varchar(50);not null;uniqueIndex:idx_hotel_room_number,priority:2"`
	HotelID       uuid.UUID      `gorm:"type:uuid;not null;uniqueIndex:idx_hotel_room_number,priority:1;index"`
	RoomTypeID    uuid.UUID      `gorm:"type:uuid;not null;index"`
	FloorNumber   int            `gorm:"not null"`
	IsAvailable   bool           `gorm:"default:true"`
	IsMaintenance bool           `gorm:"default:false"`
	Hotel         entity.Hotel   `gorm:"foreignKey:HotelID"`
	RoomType      RoomType       `gorm:"foreignKey:RoomTypeID"`
	CreatedAt     time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}