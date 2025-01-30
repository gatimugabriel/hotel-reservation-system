package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/reservation/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// ReservationRepository repository for Bookings
type ReservationRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Reservation, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.Reservation, error)
	GetByRoomID(ctx context.Context, roomID uuid.UUID) ([]*entity.Reservation, error)
	GetByDateRange(ctx context.Context, checkIn, checkOut time.Time) ([]*entity.Reservation, error)

	Create(ctx context.Context, reservation *entity.Reservation) error
	Update(ctx context.Context, reservation *entity.Reservation) (*entity.Reservation, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type ReservationRepositoryImpl struct {
	db *database.Service
}

func NewReservationRepository(db *database.Service) *ReservationRepositoryImpl {
	return &ReservationRepositoryImpl{db: db}
}

func (repo *ReservationRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*entity.Reservation, error) {
	var reservation entity.Reservation
	if err := repo.db.WithContext(ctx).
		Preload("Payment").
		Preload("Room.Hotel").
		Preload("Room").
		Preload("Room.RoomType").
		Preload("User").
		Where("id = ?", id).First(&reservation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("reservation not found")
		}
		return nil, err
	}
	return &reservation, nil
}

func (repo *ReservationRepositoryImpl) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.Reservation, error) {
	var reservations []*entity.Reservation
	if err := repo.db.WithContext(ctx).
		Preload("Room").
		Preload("User").
		Preload("Room.RoomType").
		Preload("Room.Hotel").
		Where("user_id = ?", userID).Find(&reservations).Error; err != nil {
		return nil, fmt.Errorf("failed to get reservations by user ID: %w", err)
	}
	return reservations, nil
}

func (repo *ReservationRepositoryImpl) GetByRoomID(ctx context.Context, roomID uuid.UUID) ([]*entity.Reservation, error) {
	var reservations []*entity.Reservation
	if err := repo.db.WithContext(ctx).Where("room_id = ?", roomID).Find(&reservations).Error; err != nil {
		return nil, fmt.Errorf("failed to get reservations by room ID: %w", err)
	}
	return reservations, nil
}

func (repo *ReservationRepositoryImpl) GetByDateRange(ctx context.Context, checkIn, checkOut time.Time) ([]*entity.Reservation, error) {
	var reservations []*entity.Reservation
	if err := repo.db.DB.WithContext(ctx).
		Where("(check_in_date BETWEEN ? AND ?) OR (check_out_date BETWEEN ? AND ?)",
			checkIn, checkOut, checkIn, checkOut).
		Find(&reservations).Error; err != nil {
		return nil, fmt.Errorf("failed to get reservations by date range: %w", err)
	}
	return reservations, nil
}

func (repo *ReservationRepositoryImpl) Create(ctx context.Context, reservation *entity.Reservation) error {
	if err := repo.db.WithContext(ctx).Create(reservation).Error; err != nil {
		return fmt.Errorf("failed to create reservation: %w", err)
	}
	return nil
}

func (repo *ReservationRepositoryImpl) Update(ctx context.Context, reservation *entity.Reservation) (*entity.Reservation, error) {
	updatedReservation := reservation

	if err := repo.db.WithContext(ctx).Save(&reservation).Error; err != nil {
		return nil, fmt.Errorf("failed to update reservation: %w", err)
	}
	return updatedReservation, nil
}

func (repo *ReservationRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	if err := repo.db.WithContext(ctx).Delete(&entity.Reservation{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete reservation: %w", err)
	}
	return nil
}