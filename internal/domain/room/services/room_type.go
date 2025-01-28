package services

import (
	"context"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/repository"
	"github.com/google/uuid"
)

type RoomTypeService interface {
	CreateRoomType(ctx context.Context, roomType *entity.RoomType) (*entity.RoomType, error)
	GetRoomType(ctx context.Context, id uuid.UUID) (*entity.RoomType, error)
	ListRoomTypes(ctx context.Context) ([]*entity.RoomType, error)
	UpdateRoomType(ctx context.Context, roomType *entity.RoomType) (*entity.RoomType, error)
}

type RoomTypeServiceImpl struct {
	repo repository.RoomTypeRepository
}

func (r RoomTypeServiceImpl) CreateRoomType(ctx context.Context, roomType *entity.RoomType) (*entity.RoomType, error) {
	if err := r.repo.Create(ctx, roomType); err != nil {
		return nil, err
	}
	return roomType, nil
}

func (r RoomTypeServiceImpl) GetRoomType(ctx context.Context, id uuid.UUID) (*entity.RoomType, error) {
	return r.repo.GetByID(ctx, id)
}

func (r RoomTypeServiceImpl) ListRoomTypes(ctx context.Context) ([]*entity.RoomType, error) {
	return r.repo.GetAll(ctx)
}

func (r RoomTypeServiceImpl) UpdateRoomType(ctx context.Context, roomType *entity.RoomType) (*entity.RoomType, error) {
	if err := r.repo.Update(ctx, roomType); err != nil {
		return nil, err
	}
	return roomType, nil
}

func NewRoomTypeService(roomTypeRepo repository.RoomTypeRepository) *RoomTypeServiceImpl {
	return &RoomTypeServiceImpl{
		repo: roomTypeRepo,
	}
}