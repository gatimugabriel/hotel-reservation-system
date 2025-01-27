package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Hotel : represents the hotel entity
type Hotel struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name          string         `gorm:"type:varchar(255);not null" json:"name" validate:"required,min=3,max=255"`
	Description   string         `gorm:"type:text" json:"description" validate:"max=1000"`
	Address       string         `gorm:"type:varchar(255);not null" json:"address" validate:"required,min=5,max=255"`
	City          string         `gorm:"type:varchar(100);not null;index" json:"city" validate:"required,min=2,max=100"`
	Country       string         `gorm:"type:varchar(100);not null" json:"country" validate:"required,min=2,max=100"`
	Latitude      float64        `gorm:"type:decimal(10,8)" json:"latitude" validate:"latitude"`
	Longitude     float64        `gorm:"type:decimal(11,8)" json:"longitude" validate:"longitude"`
	ContactNumber string         `gorm:"type:varchar(20)" json:"contact_number" validate:"required,e164"`
	Email         string         `gorm:"type:varchar(255)" json:"email" validate:"required,email"`
	OwnerID       uuid.UUID      `gorm:"type:uuid;not null" json:"owner_id"`
	ManagerID     uuid.UUID      `gorm:"type:uuid;not null" json:"manager_id"`
	IsActive      bool           `gorm:"default:true" json:"is_active"`
	CreatedAt     time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}