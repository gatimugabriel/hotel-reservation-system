package payment

import (
	"github.com/google/uuid"
)

type Service interface {
	ProcessPayment(payment *Payment) error
	RefundPayment(paymentID uuid.UUID) error
	GetPaymentHistory(userID uuid.UUID) ([]*Payment, error)
	ValidatePaymentMethod(method Method) error
}