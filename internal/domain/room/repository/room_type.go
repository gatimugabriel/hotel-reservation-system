package repository

import (
	"context"
	"errors"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/database"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

// RoomTypeRepository defines the interface for room type persistence operations
type RoomTypeRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entity.RoomType, error)
	GetByName(ctx context.Context, name string) (*entity.RoomType, error)
	GetAll(ctx context.Context) ([]*entity.RoomType, error)
	Create(ctx context.Context, roomType *entity.RoomType) error
	Update(ctx context.Context, roomType *entity.RoomType) error
	Delete(ctx context.Context, id uuid.UUID) error

	CreateBed(ctx context.Context, bedType *entity.BedType) error
}

// RoomTypeRepositoryImpl implements the RoomTypeRepository interface
type RoomTypeRepositoryImpl struct {
	db *database.Service
}

func NewRoomTypeRepository(db *database.Service) *RoomTypeRepositoryImpl {
	return &RoomTypeRepositoryImpl{
		db: db,
	}
}

func (repo *RoomTypeRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*entity.RoomType, error) {
	var roomType entity.RoomType
	err := repo.db.WithContext(ctx).First(&roomType, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &roomType, nil
}

func (repo *RoomTypeRepositoryImpl) GetByName(ctx context.Context, name string) (*entity.RoomType, error) {
	var roomType entity.RoomType
	err := repo.db.WithContext(ctx).First(&roomType, "name = ?", name).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &roomType, nil
}

func (repo *RoomTypeRepositoryImpl) Create(ctx context.Context, rt *entity.RoomType) error {
	return repo.db.Transaction(ctx, func(tx *gorm.DB) error {
		return tx.Create(rt).Error
	})
}

func (repo *RoomTypeRepositoryImpl) Update(ctx context.Context, rt *entity.RoomType) error {
	return repo.db.Transaction(ctx, func(tx *gorm.DB) error {
		result := tx.Save(rt)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("room type not found")
		}
		return nil
	})
}

func (repo *RoomTypeRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return repo.db.Transaction(ctx, func(tx *gorm.DB) error {
		result := tx.Delete(&entity.RoomType{}, "id = ?", id)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("room type not found")
		}
		return nil
	})
}

// GetAll returns all room types
func (repo *RoomTypeRepositoryImpl) GetAll(ctx context.Context) ([]*entity.RoomType, error) {
	var roomTypes []*entity.RoomType
	err := repo.db.WithContext(ctx).Find(&roomTypes).Error
	if err != nil {
		return nil, err
	}
	return roomTypes, nil
}

func (repo *RoomTypeRepositoryImpl) CreateBed(ctx context.Context, bt *entity.BedType) error {
	return repo.db.Transaction(ctx, func(tx *gorm.DB) error {
		return tx.Create(bt).Error
	})
}