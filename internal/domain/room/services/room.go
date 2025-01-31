package services

import (
	"context"
	"fmt"
	reservationRepository "github.com/gatimugabriel/hotel-reservation-system/internal/domain/reservation/repository"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/repository"
	"github.com/google/uuid"
	"time"
)

type RoomService interface {
	CreateRoom(ctx context.Context, room *entity.Room) (*entity.Room, error)
	UpdateRoom(ctx context.Context, id uuid.UUID, room *entity.Room) (*entity.Room, error)
	DeleteRoom(ctx context.Context, id uuid.UUID) error

	CheckAvailability(ctx context.Context, checkIn, checkOut time.Time) (map[string][]*entity.Room, error)
	GetRooms(ctx context.Context, filters map[string]interface{}) ([]*entity.Room, error)
	GetRoom(ctx context.Context, id string) (*entity.Room, error)
}

type RoomServiceImpl struct {
	roomRepo        repository.RoomRepository
	roomTypeRepo    repository.RoomTypeRepository
	reservationRepo reservationRepository.ReservationRepository
}

func NewRoomService(roomRepo repository.RoomRepository, roomTypeRepo repository.RoomTypeRepository, reservationRepo reservationRepository.ReservationRepository) *RoomServiceImpl {
	return &RoomServiceImpl{
		roomRepo:        roomRepo,
		roomTypeRepo:    roomTypeRepo,
		reservationRepo: reservationRepo,
	}
}

func (r *RoomServiceImpl) CheckAvailability(ctx context.Context, checkIn, checkOut time.Time) (map[string][]*entity.Room, error) {
	rooms, err := r.roomRepo.GetAvailableRooms(ctx, checkIn)
	if err != nil {
		return nil, fmt.Errorf("failed to get available rooms: %w", err)
	}

	// Categorize rooms by their type
	categorizedRooms := make(map[string][]*entity.Room)
	for _, room := range rooms {
		roomType := room.RoomType.Name
		categorizedRooms[roomType] = append(categorizedRooms[roomType], room)
	}

	return categorizedRooms, nil
}

func (r *RoomServiceImpl) CreateRoom(ctx context.Context, room *entity.Room) (*entity.Room, error) {
	// If RoomTypeID is not provided, look it up by RoomTypeName
	if room.RoomTypeID == uuid.Nil && room.RoomType.Name != "" {
		roomType, err := r.roomTypeRepo.GetByName(ctx, room.RoomType.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to get room type by name: %w", err)
		}
		if roomType == nil {
			return nil, fmt.Errorf("room type with name '%s' not found", room.RoomType.Name)
		}
		room.RoomTypeID = roomType.ID
	}

	if room.RoomTypeID == uuid.Nil {
		return nil, fmt.Errorf("room type ID or name is required")
	}

	// Create the room
	if err := r.roomRepo.Create(ctx, room); err != nil {
		return nil, fmt.Errorf("failed to create room: %w", err)
	}
	return room, nil
}

func (r *RoomServiceImpl) GetRoom(ctx context.Context, id string) (*entity.Room, error) {
	roomID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid room ID: %w", err)
	}

	rooms, err := r.roomRepo.GetRooms(ctx, map[string]interface{}{"id": roomID})
	if err != nil {
		return nil, fmt.Errorf("failed to get room: %w", err)
	}
	if len(rooms) == 0 {
		return nil, nil
	}
	return rooms[0], nil
}

func (r *RoomServiceImpl) GetRooms(ctx context.Context, filters map[string]interface{}) ([]*entity.Room, error) {
	rooms, err := r.roomRepo.GetRooms(ctx, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to get rooms: %w", err)
	}

	return rooms, nil
}

func (r *RoomServiceImpl) UpdateRoom(ctx context.Context, id uuid.UUID, room *entity.Room) (*entity.Room, error) {
	existingRoom, err := r.roomRepo.GetRooms(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, fmt.Errorf("failed to get room: %w", err)
	}
	if len(existingRoom) == 0 {
		return nil, fmt.Errorf("room not found")
	}

	room.ID = existingRoom[0].ID
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