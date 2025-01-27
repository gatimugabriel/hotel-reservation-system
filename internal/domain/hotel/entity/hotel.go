package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Hotel : represents the hotel entity
type Hotel struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name          string         `gorm:"type:varchar(255);not null"`
	Description   string         `gorm:"type:text"`
	Address       string         `gorm:"type:varchar(255);not null"`
	City          string         `gorm:"type:varchar(100);not null;index"`
	Country       string         `gorm:"type:varchar(100);not null"`
	Latitude      float64        `gorm:"type:decimal(10,8)"`
	Longitude     float64        `gorm:"type:decimal(11,8)"`
	ContactNumber string         `gorm:"type:varchar(20)"`
	Email         string         `gorm:"type:varchar(255)"`
	OwnerID       uuid.UUID      `gorm:"type:uuid;not null"`
	ManagerID     uuid.UUID      `gorm:"type:uuid;not null"`
	IsActive      bool           `gorm:"default:true"`
	CreatedAt     time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}