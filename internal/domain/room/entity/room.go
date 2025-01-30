package entity

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/hotel/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Room struct {
	ID               uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	RoomNumber       int       `gorm:"type:int;not null;uniqueIndex:idx_hotel_room_number,priority:2" json:"room_number" validate:"required,min=1"`
	FloorNumber      int       `gorm:"type:int;default=0" json:"floor_number" validate:"required,min=0"`
	IsAvailable      bool      `gorm:"default:true" json:"is_available"`
	UnderMaintenance bool      `gorm:"default:false" json:"under_maintenance"`
	AvailableFrom    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"available_from"`

	RoomTypeID uuid.UUID `gorm:"type:uuid;not null;index" json:"room_type_id" validate:"required"`
	RoomType   RoomType  `gorm:"foreignKey:RoomTypeID" json:"room_type"`

	HotelID uuid.UUID    `gorm:"type:uuid;not null;uniqueIndex:idx_hotel_room_number,priority:1;index" json:"hotel_id" validate:"required"`
	Hotel   entity.Hotel `gorm:"foreignKey:HotelID" json:"hotel_details"`

	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}