package entity

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

// Status represents the state of a reservation
type Status string

const (
	StatusPending    Status = "PENDING"
	StatusConfirmed  Status = "CONFIRMED"
	StatusInProgress Status = "IN_PROGRESS"
	StatusCompleted  Status = "COMPLETED"
	StatusCancelled  Status = "CANCELLED"
)

// Reservation represents a booking of a room
type Reservation struct {
	ID              uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID          uuid.UUID      `gorm:"type:uuid;not null" json:"user_id" validate:"required"`
	RoomID          uuid.UUID      `gorm:"type:uuid;not null" json:"room_id" validate:"required"`
	CheckInDate     time.Time      `gorm:"not null" json:"check_in_date" validate:"required,gtefield=CreatedAt"`
	CheckOutDate    time.Time      `gorm:"not null" json:"check_out_date" validate:"required,gtefield=CheckInDate"`
	NumGuests       int            `gorm:"not null" json:"num_guests" validate:"required,min=1,max=10"`
	SpecialRequests string         `gorm:"type:text" json:"special_requests" validate:"max=500"`
	TotalPrice      float64        `gorm:"type:decimal(10,2);not null" json:"total_price" validate:"required,min=0"`
	Status          Status         `gorm:"type:varchar(20);not null" json:"status" validate:"required,oneof=PENDING CONFIRMED CANCELLED COMPLETED"`
	CreatedAt       time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}