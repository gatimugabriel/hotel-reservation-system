package services

import (
	"context"
	"fmt"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/reservation/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/reservation/repository"
	roomEntity "github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	roomRepository "github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/repository"
	"github.com/google/uuid"
	"time"
)

type ReservationService interface {
	CreateReservation(ctx context.Context, reservation *entity.Reservation) (*entity.Reservation, error)
	UpdateReservation(ctx context.Context, reservation *entity.Reservation) (*entity.Reservation, error)

	GetUserReservations(ctx context.Context, userID uuid.UUID) ([]*entity.Reservation, error)
	GetReservation(ctx context.Context, id uuid.UUID) (*entity.Reservation, error)
	GetRoomByNumber(ctx context.Context, roomNumber int) (*roomEntity.Room, error)

	ValidateReservation(ctx context.Context, reservation *entity.Reservation) error
}

type ReservationServiceImpl struct {
	reservationRepo repository.ReservationRepository
	roomRepo        roomRepository.RoomRepository
}

func NewReservationService(reservationRepo repository.ReservationRepository, roomRepo roomRepository.RoomRepository) *ReservationServiceImpl {
	return &ReservationServiceImpl{
		reservationRepo: reservationRepo,
		roomRepo:        roomRepo,
	}
}

func (r *ReservationServiceImpl) CreateReservation(ctx context.Context, reservation *entity.Reservation) (*entity.Reservation, error) {
	if err := r.ValidateReservation(ctx, reservation); err != nil {
		return nil, err
	}

	if err := r.reservationRepo.Create(ctx, reservation); err != nil {
		return nil, fmt.Errorf("failed to create reservation: %w", err)
	}
	return reservation, nil
}

func (r *ReservationServiceImpl) UpdateReservation(ctx context.Context, reservation *entity.Reservation) (*entity.Reservation, error) {
	updatedReservation, err := r.reservationRepo.Update(ctx, reservation)
	if err != nil {
		return nil, fmt.Errorf("failed to update reservation: %w", err)
	}
	return updatedReservation, nil
}

func (r *ReservationServiceImpl) GetUserReservations(ctx context.Context, userID uuid.UUID) ([]*entity.Reservation, error) {
	reservations, err := r.reservationRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user reservations: %w", err)
	}
	return reservations, nil
}

func (r *ReservationServiceImpl) GetReservation(ctx context.Context, id uuid.UUID) (*entity.Reservation, error) {
	reservation, err := r.reservationRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get reservation: %w", err)
	}
	return reservation, nil
}

func (r *ReservationServiceImpl) GetRoomByNumber(ctx context.Context, roomNumber int) (*roomEntity.Room, error) {
	room, err := r.roomRepo.GetByNumber(ctx, roomNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get room by number: %w", err)
	}
	return room, nil
}

func (r *ReservationServiceImpl) ValidateReservation(ctx context.Context, reservation *entity.Reservation) error {
	// Check for valid dates
	if reservation.CheckInDate.After(reservation.CheckOutDate) {
		return fmt.Errorf("check-in date must be before check-out date")
	}

	if reservation.CheckInDate.Before(time.Now()) {

		return fmt.Errorf("check-in date must be in the future")
	}

	conflictingReservations, err := r.reservationRepo.GetByDateRange(ctx, reservation.CheckInDate, reservation.CheckOutDate)
	if err != nil {
		return fmt.Errorf("failed to check room availability: %w", err)
	}

	for _, existing := range conflictingReservations {
		if existing.RoomID == reservation.RoomID && existing.ID != reservation.ID {
			return fmt.Errorf("room is not available for the selected dates, already booked/it might have been booked before you reserved it")
		}
	}

	// Check if room exists and has not been marked under maintenance
	room, err := r.roomRepo.GetByID(ctx, reservation.RoomID)
	if err != nil || room == nil {
		return fmt.Errorf("failed to get room with that ID: %v", err)
	}

	if room.UnderMaintenance {
		return fmt.Errorf("room has been marked under maintenance")
	}

	return nil
}