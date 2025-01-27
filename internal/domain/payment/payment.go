package payment

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Method : payment methods supported by the system
type Method string

const (
	MethodCreditCard   Method = "CREDIT_CARD"
	MethodDebitCard    Method = "DEBIT_CARD"
	MethodPayPal       Method = "PAYPAL"
	MethodBankTransfer Method = "BANK_TRANSFER"
	MethodCrypto       Method = "CRYPTO"
)

// Status : state of a payment
type Status string

const (
	StatusPending  Status = "PENDING"
	StatusSuccess  Status = "SUCCESS"
	StatusFailed   Status = "FAILED"
	StatusRefunded Status = "REFUNDED"
)

// Payment : represents a payment transaction
type Payment struct {
	ID             uuid.UUID
	ReservationID  uuid.UUID
	UserID         uuid.UUID
	Amount         float64
	Currency       string
	PaymentMethod  Method
	PaymentStatus  Status
	TransactionID  string
	PaymentDetails json.RawMessage
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Repository interface {
	GetByID(id uuid.UUID) (*Payment, error)
	GetByReservationID(reservationID uuid.UUID) ([]*Payment, error)
	GetByUserID(userID uuid.UUID) ([]*Payment, error)
	Create(payment *Payment) error
	Update(payment *Payment) error
}