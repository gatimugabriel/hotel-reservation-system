package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// RoomRepository : data persistence database interaction interface
type RoomRepository interface {
	GetRooms(ctx context.Context, filters map[string]interface{}) ([]*entity.Room, error)
	GetAvailableRooms(ctx context.Context, checkIn time.Time) ([]*entity.Room, error)

	Create(ctx context.Context, room *entity.Room) error
	Update(ctx context.Context, room *entity.Room) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// RoomRepositoryImpl implements the RoomRepository interface
type RoomRepositoryImpl struct {
	db *database.Service
}

func NewRoomRepository(db *database.Service) *RoomRepositoryImpl {
	return &RoomRepositoryImpl{
		db: db,
	}
}

// GetRooms can get single or multiple rooms that mat the given filters
func (repo *RoomRepositoryImpl) GetRooms(ctx context.Context, filters map[string]interface{}) ([]*entity.Room, error) {
	var rooms []*entity.Room

	query := repo.db.WithContext(ctx).
		Preload("RoomType")

	for key, value := range filters {
		query = query.Where(key, value)
	}

	err := query.Find(&rooms).Error
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

// GetAvailableRooms fetches available rooms that are available as the from given date
// looks similar to GetRooms above but due query needs , I decided to keep it this way
func (repo *RoomRepositoryImpl) GetAvailableRooms(ctx context.Context, checkIn time.Time) ([]*entity.Room, error) {
	var rooms []*entity.Room

	err := repo.db.WithContext(ctx).
		Preload("RoomType").
		Where("under_maintenance = ? AND (is_available = ? OR available_from <= ?)", false, true, checkIn).
		Find(&rooms).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get available rooms: %w", err)
	}

	return rooms, nil
}

func (repo *RoomRepositoryImpl) Create(ctx context.Context, r *entity.Room) error {
	return repo.db.Transaction(ctx, func(tx *gorm.DB) error {
		return tx.Create(r).Error
	})
}

func (repo *RoomRepositoryImpl) Update(ctx context.Context, r *entity.Room) error {
	return repo.db.Transaction(ctx, func(tx *gorm.DB) error {
		result := tx.Save(r)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("room not found")
		}
		return nil
	})
}

func (repo *RoomRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return repo.db.Transaction(ctx, func(tx *gorm.DB) error {
		result := tx.Delete(&entity.Room{}, "id = ?", id)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("room not found")
		}
		return nil
	})
}