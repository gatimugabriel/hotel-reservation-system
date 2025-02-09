package entity

import (
	"github.com/google/uuid"
	"time"
)

// RoomType : represents a category of rooms with similar characteristics
type RoomType struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name         string    `gorm:"type:citext;not null;unique" json:"name" validate:"required,min=2,max=100"`
	Description  string    `gorm:"type:citext" json:"description" validate:"max=500"`
	BasePrice    float64   `gorm:"type:decimal(10,2);not null" json:"base_price" validate:"required,min=0"`
	MaxOccupancy int       `gorm:"not null" json:"max_occupancy" validate:"required,min=1,max=10"`
	NumBeds      int       `gorm:"not null" json:"num_beds" validate:"required,min=1,max=5"`
	SquareMeters float64   `gorm:"type:decimal(6,2);not null" json:"square_meters" validate:"required,min=1"`
	Status       string    `gorm:"type:citext;not null;default:'ACTIVE'" json:"status" validate:"oneof=ACTIVE INACTIVE"`

	BedTypeID uuid.UUID `gorm:"type:uuid;not null;index" json:"bed_type_id" validate:"required"`
	Bed       BedType   `gorm:"foreignKey:BedTypeID" json:"bed_type"`

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}