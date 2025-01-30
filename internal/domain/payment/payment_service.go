package payment

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/payment/entity"
	"github.com/google/uuid"
)

type Service interface {
	ProcessPayment(payment *entity.Payment) error
	RefundPayment(paymentID uuid.UUID) error
	GetPaymentHistory(userID uuid.UUID) ([]*entity.Payment, error)
	ValidatePaymentMethod(method entity.Method) error
}