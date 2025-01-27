package services

import (
	"context"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	"github.com/google/uuid"
)

type RoomTypeService interface {
	CreateRoomType(ctx context.Context, roomType *entity.RoomType) error
	UpdateRoomType(ctx context.Context, roomType *entity.RoomType) error
	DeleteRoomType(ctx context.Context, id uuid.UUID) error
}