package services

import (
	"context"
	"fmt"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/repository"
	"github.com/google/uuid"
	"time"
)

type RoomService interface {
	CreateRoom(ctx context.Context, room *entity.Room) (*entity.Room, error)
	UpdateRoom(ctx context.Context, id uuid.UUID, room *entity.Room) (*entity.Room, error)
	DeleteRoom(ctx context.Context, id uuid.UUID) error
	CheckAvailability(ctx context.Context, hotelID uuid.UUID, checkIn, checkOut time.Time) ([]*entity.Room, error)
	MarkRoomMaintenance(ctx context.Context, id uuid.UUID, maintenance bool) error
	GetRoom(ctx context.Context, id string) (*entity.Room, error)
	GetRooms(ctx context.Context, id string) ([]*entity.Room, error)
}

type RoomServiceImpl struct {
	roomRepo repository.RoomRepository
}

func NewRoomService(roomRepo repository.RoomRepository) *RoomServiceImpl {
	return &RoomServiceImpl{
		roomRepo: roomRepo,
	}
}

func (r *RoomServiceImpl) CreateRoom(ctx context.Context, room *entity.Room) (*entity.Room, error) {
	if err := r.roomRepo.Create(ctx, room); err != nil {
		return nil, fmt.Errorf("failed to create room: %w", err)
	}
	return room, nil
}

func (r *RoomServiceImpl) UpdateRoom(ctx context.Context, id uuid.UUID, room *entity.Room) (*entity.Room, error) {
	existingRoom, err := r.roomRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get room: %w", err)
	}

	room.ID = existingRoom.ID
	room.UpdatedAt = time.Now()

	if err := r.roomRepo.Update(ctx, room); err != nil {
		return nil, fmt.Errorf("failed to update room: %w", err)
	}
	return room, nil
}

func (r *RoomServiceImpl) DeleteRoom(ctx context.Context, id uuid.UUID) error {
	if err := r.roomRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete room: %w", err)
	}
	return nil
}

func (r *RoomServiceImpl) CheckAvailability(ctx context.Context, hotelID uuid.UUID, checkIn, checkOut time.Time) ([]*entity.Room, error) {
	// Get all rooms for the hotel
	rooms, err := r.roomRepo.GetByHotelID(ctx, hotelID)
	if err != nil {
		return nil, fmt.Errorf("failed to get rooms: %w", err)
	}

	// Get reservations for the date range
	reservations, err := r.roomRepo.GetReservationsForDateRange(ctx, checkIn, checkOut)
	if err != nil {
		return nil, fmt.Errorf("failed to get reservations: %w", err)
	}

	// Create a map of reserved room IDs
	reservedRooms := make(map[uuid.UUID]bool)
	for _, reservation := range reservations {
		reservedRooms[reservation.ID] = true
	}

	// Filter available rooms
	var availableRooms []*entity.Room
	for _, room := range rooms {
		if !reservedRooms[room.ID] && !room.IsMaintenance {
			availableRooms = append(availableRooms, room)
		}
	}

	return availableRooms, nil
}

func (r *RoomServiceImpl) MarkRoomMaintenance(ctx context.Context, id uuid.UUID, maintenance bool) error {
	room, err := r.roomRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get room: %w", err)
	}

	room.IsMaintenance = maintenance
	room.UpdatedAt = time.Now()

	if err := r.roomRepo.Update(ctx, room); err != nil {
		return fmt.Errorf("failed to update room maintenance status: %w", err)
	}
	return nil
}

func (r *RoomServiceImpl) GetRoom(ctx context.Context, id string) (*entity.Room, error) {
	roomID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid room ID: %w", err)
	}

	room, err := r.roomRepo.GetByID(ctx, roomID)
	if err != nil {
		return nil, fmt.Errorf("failed to get room: %w", err)
	}
	return room, nil
}

func (r *RoomServiceImpl) GetRooms(ctx context.Context, hotelID string) ([]*entity.Room, error) {
	id, err := uuid.Parse(hotelID)
	if err != nil {
		return nil, fmt.Errorf("invalid hotel ID: %w", err)
	}

	rooms, err := r.roomRepo.GetByHotelID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get rooms: %w", err)
	}
	return rooms, nil
}