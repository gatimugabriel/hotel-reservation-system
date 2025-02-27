package entity

import (
	"encoding/json"
	paymentEntity "github.com/gatimugabriel/hotel-reservation-system/internal/domain/payment/entity"
	roomEntity "github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	userEntity "github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/entity"
	"github.com/google/uuid"
	"time"
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

// Reservation represents the booking of a room
type Reservation struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CheckInDate    time.Time `gorm:"not null;uniqueIndex:idx_room_dates,priority:2" json:"check_in_date" validate:"required,gtefield=CreatedAt"`
	CheckOutDate   time.Time `gorm:"not null;uniqueIndex:idx_room_dates,priority:3" json:"check_out_date" validate:"required,gtefield=CheckInDate"`
	NumGuests      int       `gorm:"not null;default:1" json:"num_guests" validate:"required,min=1,max=10"`
	SpecialRequest string    `gorm:"type:text" json:"special_request" validate:"max=500"`
	TotalPrice     float64   `gorm:"type:decimal(10,2);not null" json:"total_price" validate:"required,min=0"`
	Status         Status    `gorm:"type:varchar(20);not null;default:PENDING" json:"status" validate:"required,oneof=PENDING CONFIRMED CANCELLED COMPLETED"`

	RoomID uuid.UUID       `gorm:"type:uuid;not null;uniqueIndex:idx_room_dates,priority:1" json:"room_id" validate:"required"`
	Room   roomEntity.Room `gorm:"foreignKey:RoomID;references:ID" json:"room"`

	UserID uuid.UUID       `gorm:"type:uuid;not null;index" json:"user_id" validate:"required"`
	User   userEntity.User `gorm:"foreignKey:UserID;references:ID" json:"user"`

	PaymentID uuid.UUID             `gorm:"type:uuid;index" json:"payment_id"`
	Payment   paymentEntity.Payment `gorm:"foreignKey:PaymentID;references:ID" json:"payment,omitempty"`

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// CreateReservationRequest represents the data object used when user need to create a reservation
type CreateReservationRequest struct {
	RoomID     *string `json:"room_id"`
	RoomNumber *int    `json:"room_number"`

	CheckInDate  string `json:"check_in_date" validate:"required"`
	CheckoutDate string `json:"check_out_date" validate:"required"`

	NumGuests      int                  `json:"num_guests" validate:"required"`
	SpecialRequest *string              `json:"special_request"`
	PaymentMethod  paymentEntity.Method `json:"payment_method,omitempty" validate:"required,oneof=CREDIT_CARD DEBIT_CARD PAYPAL BANK_TRANSFER CRYPTO CASH"`
	PaymentDetails json.RawMessage      `json:"payment_details,omitempty" validate:"required"`
}