package repository

import (
	"context"
	"fmt"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/payment/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/database"
	"github.com/google/uuid"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *entity.Payment) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Payment, error)
	Update(ctx context.Context, payment *entity.Payment) error
}

type PaymentRepositoryImpl struct {
	db *database.Service
}

func NewPaymentRepository(db *database.Service) *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{db: db}
}

func (repo *PaymentRepositoryImpl) Create(ctx context.Context, payment *entity.Payment) error {
	if err := repo.db.WithContext(ctx).Create(payment).Error; err != nil {
		return fmt.Errorf("failed to create payment: %w", err)
	}
	return nil
}

func (repo *PaymentRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*entity.Payment, error) {
	var payment entity.Payment
	if err := repo.db.WithContext(ctx).First(&payment, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get payment: %w", err)
	}
	return &payment, nil
}

func (repo *PaymentRepositoryImpl) Update(ctx context.Context, payment *entity.Payment) error {
	if err := repo.db.WithContext(ctx).Save(payment).Error; err != nil {
		return fmt.Errorf("failed to update payment: %w", err)
	}
	return nil
}