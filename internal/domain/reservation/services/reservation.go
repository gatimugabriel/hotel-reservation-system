package services

import (
	"context"
	"fmt"
	paymentRepository "github.com/gatimugabriel/hotel-reservation-system/internal/domain/payment/repository"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/reservation/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/reservation/repository"
	roomEntity "github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	roomRepository "github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/repository"
	"github.com/gatimugabriel/hotel-reservation-system/pkg/utils"
	"github.com/google/uuid"
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
	paymentRepo     paymentRepository.PaymentRepository
}

func NewReservationService(reservationRepo repository.ReservationRepository, roomRepo roomRepository.RoomRepository, paymentRepo paymentRepository.PaymentRepository) *ReservationServiceImpl {
	return &ReservationServiceImpl{
		reservationRepo: reservationRepo,
		roomRepo:        roomRepo,
		paymentRepo:     paymentRepo,
	}
}

func (r *ReservationServiceImpl) CreateReservation(ctx context.Context, reservation *entity.Reservation) (*entity.Reservation, error) {
	if err := r.ValidateReservation(ctx, reservation); err != nil {
		return nil, err
	}

	// Create the payment record
	if err := r.paymentRepo.Create(ctx, &reservation.Payment); err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	// Create the reservation
	if err := r.reservationRepo.Create(ctx, reservation); err != nil {
		return nil, fmt.Errorf("failed to create reservation: %w", err)
	}

	// Fetch the created reservation with preloaded associations
	createdReservation, err := r.reservationRepo.GetByID(ctx, reservation.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get created reservation: %w", err)
	}

	// Send email notification
	go func() {
		reservationData := utils.ReservationEmailData{
			ID:            createdReservation.ID.String(),
			CheckInDate:   createdReservation.CheckInDate,
			CheckOutDate:  createdReservation.CheckOutDate,
			RoomNumber:    createdReservation.Room.RoomNumber,
			RoomType:      createdReservation.Room.RoomType.Name,
			GuestName:     fmt.Sprintf("%s %s", createdReservation.User.FirstName, createdReservation.User.LastName),
			TotalPrice:    createdReservation.TotalPrice,
			PaymentStatus: string(createdReservation.Payment.PaymentStatus),
		}

		err := utils.SendEmailNotification(createdReservation.User.Email, reservationData)
		if err != nil {
			fmt.Println("Failed to send email notification:", err)
		}
	}()

	return createdReservation, nil
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

// GetRoomByNumber returns a room with the given number
func (r *ReservationServiceImpl) GetRoomByNumber(ctx context.Context, roomNumber int) (*roomEntity.Room, error) {
	rooms, err := r.roomRepo.GetRooms(ctx, map[string]interface{}{"room_number": roomNumber})
	if err != nil {
		return nil, fmt.Errorf("failed to get room by number: %w", err)
	}
	if rooms == nil || len(rooms) == 0 {
		return nil, fmt.Errorf("room with that number not found")
	}
	return rooms[0], nil
}

func (r *ReservationServiceImpl) ValidateReservation(ctx context.Context, reservation *entity.Reservation) error {
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
	rooms, err := r.roomRepo.GetRooms(ctx, map[string]interface{}{"id": reservation.RoomID})
	if err != nil || len(rooms) == 0 {
		return fmt.Errorf("failed to get room with that ID: %v", err)
	}

	room := rooms[0]
	if room.Status == "under_maintenance" {
		return fmt.Errorf("room is under maintenance")
	}

	return nil
}