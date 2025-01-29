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
	//gorm.Model
	ID             uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID         uuid.UUID      `gorm:"type:uuid;not null" json:"user_id" validate:"required"`
	RoomID         uuid.UUID      `gorm:"type:uuid;not null" json:"room_id" validate:"required"`
	CheckInDate    time.Time      `gorm:"not null" json:"check_in_date" validate:"required,gtefield=CreatedAt"`
	CheckOutDate   time.Time      `gorm:"not null" json:"check_out_date" validate:"required,gtefield=CheckInDate"`
	NumGuests      int            `gorm:"not null;default:1" json:"num_guests" validate:"required,min=1,max=10"`
	SpecialRequest string         `gorm:"type:text" json:"special_request" validate:"max=500"`
	TotalPrice     float64        `gorm:"type:decimal(10,2);not null" json:"total_price" validate:"required,min=0"`
	Status         Status         `gorm:"type:varchar(20);not null;default:PENDING" json:"status" validate:"required,oneof=PENDING CONFIRMED CANCELLED COMPLETED"`
	CreatedAt      time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type CreateReservationRequest struct {
	RoomID     *string `json:"room_id"`
	RoomNumber *int    `json:"room_number"`

	CheckInDate  string `json:"check_in_date"`
	CheckoutDate string `json:"check_out_date"`

	//CheckInDate  time.Time `json:"check_in_date"`
	//CheckoutDate time.Time `json:"check_out_date"`

	NumGuests      int     `json:"num_guests"`
	SpecialRequest *string `json:"special_request"`
}