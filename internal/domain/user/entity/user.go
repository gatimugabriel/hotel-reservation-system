package entity

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/constants"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID            `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	FirstName    string               `gorm:"type:varchar(255);not null" json:"first_name" validate:"required,min=3,max=20"`
	LastName     string               `gorm:"type:varchar(255);not null" json:"last_name" validate:"required,min=3,max=20"`
	Email        string               `gorm:"type:varchar(100);not null;unique" json:"email" validate:"required,email"`
	Phone        string               `gorm:"type:varchar(20);not null;unique" json:"phone" validate:"required,e164"`
	PasswordHash string               `gorm:"type:varchar(100);not null" json:"password" validate:"required,min=8,passwd"`
	Role         constants.Role       `gorm:"type:varchar(20);not null" json:"role" validate:"oneof=GUEST STAFF MANAGER ADMIN PROPERTYOWNER"`
	IsVerified   bool                 `gorm:"default:false" json:"is_verified"`
	Status       constants.UserStatus `gorm:"type:varchar(20);default:ACTIVE" json:"status" validate:"oneof=ACTIVE INACTIVE"`

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}