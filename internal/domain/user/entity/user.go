package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Role string

const (
	RoleGuest   Role = "GUEST"
	RoleStaff   Role = "STAFF"
	RoleManager Role = "MANAGER"
	RoleAdmin   Role = "ADMIN"
	RoleOwner   Role = "OWNER"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	FirstName    string    `gorm:"type:varchar(255);not null" json:"first_name" validate:"required,min=3,max=20"`
	LastName     string    `gorm:"type:varchar(255);not null" json:"last_name" validate:"required,min=3,max=20"`
	Email        string    `gorm:"type:varchar(100);not null;unique" json:"email" validate:"required,email"`
	Phone        string    `gorm:"type:varchar(20):not null;unique" json:"phone" validate:"required,e164"`
	PasswordHash string    `gorm:"type:varchar(100);not null" json:"password" validate:"required,min=8,passwd"`
	Role         Role      `gorm:"type:varchar(10);not null" json:"role" validate:"oneof=GUEST STAFF MANAGER ADMIN OWNER"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	IsVerified   bool      `gorm:"default:false" json:"is_verified"`
	//AuthType      authType  `gorm:"type:varchar(20)" json:"auth_type" validate:"oneof=BASIC GOOGLE"`
	//GoogleID      string    `gorm:"type:varchar(20)" json:"google_id,omitempty"`
	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenData struct {
	UserID string
	Role   string
}
