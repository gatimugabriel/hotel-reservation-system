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

	MarkRoomMaintenance(ctx context.Context, id uuid.UUID, maintenance bool) error

	CheckAvailability(ctx context.Context, hotelID uuid.UUID, checkIn, checkOut time.Time) ([]*entity.Room, error)
	GetRooms(ctx context.Context, hotelID uuid.UUID) ([]*entity.Room, error)
	GetRoom(ctx context.Context, id string) (*entity.Room, error)
}

type RoomServiceImpl struct {
	roomRepo     repository.RoomRepository
	roomTypeRepo repository.RoomTypeRepository
}

func NewRoomService(roomRepo repository.RoomRepository, roomTypeRepo repository.RoomTypeRepository) *RoomServiceImpl {
	return &RoomServiceImpl{
		roomRepo:     roomRepo,
		roomTypeRepo: roomTypeRepo,
	}
}

func (r *RoomServiceImpl) CreateRoom(ctx context.Context, room *entity.Room) (*entity.Room, error) {
	// If RoomTypeID is not provided, look it up by RoomTypeName
	if room.RoomTypeID == uuid.Nil && room.RoomTypeInfo.Name != "" {
		roomType, err := r.roomTypeRepo.GetByName(ctx, room.RoomTypeInfo.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to get room type by name: %w", err)
		}
		if roomType == nil {
			return nil, fmt.Errorf("room type with name '%s' not found", room.RoomTypeInfo.Name)
		}
		room.RoomTypeID = roomType.ID
	}

	// Validate that RoomTypeID is set
	if room.RoomTypeID == uuid.Nil {
		return nil, fmt.Errorf("room type ID or name is required")
	}

	// Create the room
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
	//// Get all rooms for the hotel
	//rooms, err := r.roomRepo.GetByHotelID(ctx, hotelID)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to get rooms: %w", err)
	//}

	//// Get reservations for the date range
	//reservations, err := r.roomRepo.GetReservationsForDateRange(ctx, checkIn, checkOut)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to get reservations: %w", err)
	//}
	//
	//// Create a map of reserved room IDs
	//reservedRooms := make(map[uuid.UUID]bool)
	//for _, reservation := range reservations {
	//	reservedRooms[reservation.ID] = true
	//}
	//
	//// Filter available rooms
	//var availableRooms []*entity.Room
	//for _, room := range rooms {
	//	if !reservedRooms[room.ID] && !room.IsMaintenance {
	//		availableRooms = append(availableRooms, room)
	//	}
	//}
	//
	//return availableRooms, nil

	rooms, err := r.roomRepo.GetAvailableRooms(ctx, hotelID, checkIn, checkOut)
	if err != nil {
		return nil, fmt.Errorf("failed to get available rooms: %w", err)
	}

	return rooms, nil
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

func (r *RoomServiceImpl) GetRooms(ctx context.Context, hotelID uuid.UUID) ([]*entity.Room, error) {
	rooms, err := r.roomRepo.GetByHotelID(ctx, hotelID)
	if err != nil {
		return nil, fmt.Errorf("failed to get rooms: %w", err)
	}

	return rooms, nil
}