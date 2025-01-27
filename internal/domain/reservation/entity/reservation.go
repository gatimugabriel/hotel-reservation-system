package entity

import (
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
	ID              uuid.UUID
	UserID          uuid.UUID
	RoomID          uuid.UUID
	CheckInDate     time.Time
	CheckOutDate    time.Time
	NumGuests       int
	SpecialRequests string
	TotalPrice      float64
	Status          Status
	CreatedAt       time.Time
	UpdatedAt       time.Time
}