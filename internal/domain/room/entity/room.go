package entity

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/hotel/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

//
//// Room : represents an individual room model in a hotel
//type Room struct {
//	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
//	RoomNumber    string         `gorm:"type:varchar(50);not null;uniqueIndex:idx_hotel_room_number,priority:2" json:"room_number" validate:"required,min=1,max=50"`
//	HotelID       uuid.UUID      `gorm:"type:uuid;not null;uniqueIndex:idx_hotel_room_number,priority:1;index" json:"hotel_id" validate:"required"`
//	RoomTypeID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"room_type_id" validate:"required"`
//	RoomTypeName  string         `gorm:"-" json:"room_type_name"`
//	HotelName     string         `gorm:"-" json:"hotel_name"`
//	FloorNumber   int            `gorm:"type:int;default=0" json:"floor_number" validate:"required,min=0"`
//	IsAvailable   bool           `gorm:"default:true" json:"is_available"`
//	IsMaintenance bool           `gorm:"default:false" json:"is_maintenance"`
//	Hotel         entity.Hotel   `gorm:"foreignKey:HotelID" json:"hotel,omitempty"`
//	RoomType      RoomType       `gorm:"foreignKey:RoomTypeID" json:"room_type,omitempty"`
//	CreatedAt     time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
//	UpdatedAt     time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
//	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
//}

// HotelLimited reps a custom struct for the limited hotel fields
type HotelLimited struct {
	Name          string `json:"name"`
	ContactNumber string `json:"contact_number"`
}

// RoomTypeLimited reps a custom struct for room type without ID and timestamps
type RoomTypeLimited struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	BasePrice    float64 `json:"base_price"`
	MaxOccupancy int     `json:"max_occupancy"`
	NumBeds      int     `json:"num_beds"`
	BedType      string  `json:"bed_type"`
	SquareMeters float64 `json:"square_meters"`
}

type Room struct {
	ID            uuid.UUID       `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	RoomNumber    string          `gorm:"type:varchar(50);not null;uniqueIndex:idx_hotel_room_number,priority:2" json:"room_number" validate:"required,min=1,max=50"`
	HotelID       uuid.UUID       `gorm:"type:uuid;not null;uniqueIndex:idx_hotel_room_number,priority:1;index" json:"hotel_id" validate:"required"`
	RoomTypeID    uuid.UUID       `gorm:"type:uuid;not null;index" json:"room_type_id" validate:"required"`
	FloorNumber   int             `gorm:"type:int;default=0" json:"floor_number" validate:"required,min=0"`
	IsAvailable   bool            `gorm:"default:true" json:"is_available"`
	IsMaintenance bool            `gorm:"default:false" json:"is_maintenance"`
	Hotel         entity.Hotel    `gorm:"foreignKey:HotelID" json:"-"`
	HotelInfo     HotelLimited    `gorm:"-" json:"hotel"`
	RoomType      RoomType        `gorm:"foreignKey:RoomTypeID" json:"-"`
	RoomTypeInfo  RoomTypeLimited `gorm:"-" json:"room_type"`
	CreatedAt     time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
}