package entity

import (
	"encoding/json"
	userEntity "github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/entity"
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
	MethodCash         Method = "CASH"
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
	ID             uuid.UUID       `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Amount         float64         `gorm:"type:decimal(10,2);not null" json:"amount" validate:"required,min=0"`
	Currency       string          `gorm:"type:varchar(3);not null" json:"currency" validate:"required,len=3"`
	PaymentMethod  Method          `gorm:"type:varchar(20);not null" json:"payment_method" validate:"required,oneof=CREDIT_CARD DEBIT_CARD PAYPAL BANK_TRANSFER CRYPTO"`
	PaymentStatus  Status          `gorm:"type:varchar(20);not null;default:PENDING" json:"payment_status" validate:"required,oneof=PENDING SUCCESS FAILED REFUNDED"`
	TransactionID  string          `gorm:"type:varchar(100);not null;uniqueIndex" json:"transaction_id" validate:"required"`
	PaymentDetails json.RawMessage `gorm:"type:jsonb" json:"payment_details" validate:"required"`

	UserID uuid.UUID       `gorm:"type:uuid;not null;index" json:"user_id" validate:"required"`
	User   userEntity.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`

	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}