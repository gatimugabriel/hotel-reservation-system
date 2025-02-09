package entity

import (
	"github.com/google/uuid"
	"time"
)

type Room struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	RoomNumber  int        `gorm:"type:int;not null;unique" json:"room_number" validate:"required,min=1"`
	FloorNumber int        `gorm:"type:int;default=0" json:"floor_number" validate:"required,min=0"`
	Status      RoomStatus `gorm:"type:varchar(20);not null;default:'available'" json:"status" validate:"required,oneof=available unavailable under_maintenance"`

	RoomTypeID uuid.UUID `gorm:"type:uuid;not null;index" json:"room_type_id" validate:"required"`
	RoomType   RoomType  `gorm:"foreignKey:RoomTypeID" json:"room_type"`

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type RoomStatus string

const (
	Available        RoomStatus = "available"
	Unavailable      RoomStatus = "unavailable"
	UnderMaintenance RoomStatus = "under_maintenance"
)