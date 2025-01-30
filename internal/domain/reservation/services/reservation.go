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
	"log"
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

	if err := r.paymentRepo.Create(ctx, &reservation.Payment); err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	if err := r.reservationRepo.Create(ctx, reservation); err != nil {
		return nil, fmt.Errorf("failed to create reservation: %w", err)
	}

	// Mark room as unavailable
	go func() {
		ctx := context.Background()
		rooms, err := r.roomRepo.GetRooms(ctx, map[string]interface{}{"id": reservation.RoomID})
		if err != nil || len(rooms) == 0 {
			log.Printf("Failed to get room: %v", err)
			return
		}

		room := rooms[0]
		room.IsAvailable = false
		room.AvailableFrom = reservation.CheckOutDate

		if err := r.roomRepo.Update(ctx, room); err != nil {
			log.Printf("Failed to update room availability: %v", err)
		}
	}()

	createdReservation, err := r.reservationRepo.GetByID(ctx, reservation.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get created reservation: %w", err)
	}

	//send email
	go func() {
		reservationData := utils.ReservationEmailData{
			ID:            createdReservation.ID.String(),
			CheckInDate:   createdReservation.CheckInDate,
			CheckOutDate:  createdReservation.CheckOutDate,
			RoomNumber:    createdReservation.Room.RoomNumber,
			RoomType:      createdReservation.Room.RoomType.Name,
			HotelName:     createdReservation.Room.Hotel.Name,
			GuestName:     fmt.Sprintf("%s %s", createdReservation.User.FirstName, createdReservation.User.LastName),
			TotalPrice:    createdReservation.TotalPrice,
			PaymentStatus: string(createdReservation.Payment.PaymentStatus),
		}

		err := utils.SendEmailNotification(createdReservation.User.Email, reservationData)
		if err != nil {
			// I could add this to something like a queue to retry later
			// I will log it and continue for now
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
	if room.UnderMaintenance {
		return fmt.Errorf("room has been marked under maintenance")
	}

	// Check if the room is available for the given dates
	if !room.IsAvailable && room.AvailableFrom.After(reservation.CheckInDate) {
		return fmt.Errorf("room is not available until %s", room.AvailableFrom.Format(time.RFC3339))
	}

	return nil
}