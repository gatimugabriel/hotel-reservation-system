package services

import (
	"context"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/repository"
	"github.com/google/uuid"
)

type RoomTypeService interface {
	CreateRoomType(ctx context.Context, roomType *entity.RoomType) error
	GetRoomType(ctx context.Context, id uuid.UUID) (*entity.RoomType, error)
	ListRoomTypes(ctx context.Context, id uuid.UUID) ([]*entity.RoomType, error)
	UpdateRoomType(ctx context.Context, roomType *entity.RoomType) error
}

type RoomTypeServiceImpl struct {
	repo repository.RoomTypeRepository
}

func (r RoomTypeServiceImpl) CreateRoomType(ctx context.Context, roomType *entity.RoomType) error {
	//TODO implement me
	panic("implement me")
}

func (r RoomTypeServiceImpl) GetRoomType(ctx context.Context, id uuid.UUID) (*entity.RoomType, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoomTypeServiceImpl) ListRoomTypes(ctx context.Context, id uuid.UUID) ([]*entity.RoomType, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoomTypeServiceImpl) UpdateRoomType(ctx context.Context, roomType *entity.RoomType) error {
	//TODO implement me
	panic("implement me")
}

func NewRoomTypeService(roomTypeRepo repository.RoomTypeRepository) *RoomTypeServiceImpl {
	return &RoomTypeServiceImpl{
		repo: roomTypeRepo,
	}
}