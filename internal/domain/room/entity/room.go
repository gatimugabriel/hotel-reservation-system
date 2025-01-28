package entity

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/hotel/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// RoomType : represents a category of rooms with similar characteristics
type RoomType struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name         string         `gorm:"type:varchar(100);not null;unique" json:"name" validate:"required,min=2,max=100"`
	Description  string         `gorm:"type:text" json:"description" validate:"max=500"`
	BasePrice    float64        `gorm:"type:decimal(10,2);not null" json:"base_price" validate:"required,min=0"`
	MaxOccupancy int            `gorm:"not null" json:"max_occupancy" validate:"required,min=1,max=10"`
	NumBeds      int            `gorm:"not null" json:"num_beds" validate:"required,min=1,max=5"`
	BedType      string         `gorm:"type:varchar(50);not null" json:"bed_type" validate:"required,oneof=SINGLE DOUBLE QUEEN KING"`
	SquareMeters float64        `gorm:"type:decimal(6,2);not null" json:"square_meters" validate:"required,min=1"`
	Rooms        []Room         `gorm:"foreignKey:RoomTypeID" json:"rooms,omitempty"`
	CreatedAt    time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// Room : represents an individual room in a hotel
type Room struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	RoomNumber    string         `gorm:"type:varchar(50);not null;uniqueIndex:idx_hotel_room_number,priority:2" json:"room_number" validate:"required,min=1,max=50"`
	HotelID       uuid.UUID      `gorm:"type:uuid;not null;uniqueIndex:idx_hotel_room_number,priority:1;index" json:"hotel_id" validate:"required"`
	RoomTypeID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"room_type_id" validate:"required"`
	FloorNumber   int            `gorm:"type:int;default=0" json:"floor_number" validate:"required,min=0"`
	IsAvailable   bool           `gorm:"default:true" json:"is_available"`
	IsMaintenance bool           `gorm:"default:false" json:"is_maintenance"`
	Hotel         entity.Hotel   `gorm:"foreignKey:HotelID" json:"hotel,omitempty"`
	RoomType      RoomType       `gorm:"foreignKey:RoomTypeID" json:"room_type,omitempty"`
	CreatedAt     time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}