package repository

import (
	"context"
	"errors"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/hotel/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HotelRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Hotel, error)
	GetByCity(ctx context.Context, city string) ([]*entity.Hotel, error)
	Create(ctx context.Context, hotel *entity.Hotel) error
	Update(ctx context.Context, hotel *entity.Hotel) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*entity.Hotel, error)
}

type HotelRepositoryImpl struct {
	db *database.Service
}

func NewHotelRepository(db *database.Service) *HotelRepositoryImpl {
	return &HotelRepositoryImpl{
		db: db,
	}
}

func (repo *HotelRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*entity.Hotel, error) {
	var hotel entity.Hotel
	err := repo.db.WithContext(ctx).First(&hotel, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &hotel, nil
}

func (repo *HotelRepositoryImpl) GetByCity(ctx context.Context, city string) ([]*entity.Hotel, error) {
	var hotels []*entity.Hotel
	err := repo.db.WithContext(ctx).Where("city = ?", city).Find(&hotels).Error
	if err != nil {
		return nil, err
	}
	return hotels, nil
}

func (repo *HotelRepositoryImpl) Create(ctx context.Context, h *entity.Hotel) error {
	return repo.db.Transaction(ctx, func(tx *gorm.DB) error {
		return tx.Create(h).Error
	})
}

func (repo *HotelRepositoryImpl) Update(ctx context.Context, h *entity.Hotel) error {
	return repo.db.Transaction(ctx, func(tx *gorm.DB) error {
		result := tx.Save(h)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("hotel not found")
		}
		return nil
	})
}

func (repo *HotelRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return repo.db.Transaction(ctx, func(tx *gorm.DB) error {
		result := tx.Delete(&entity.Hotel{}, "id = ?", id)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("hotel not found")
		}
		return nil
	})
}

func (repo *HotelRepositoryImpl) GetAll(ctx context.Context) ([]*entity.Hotel, error) {
	var hotels []*entity.Hotel
	err := repo.db.WithContext(ctx).Find(&hotels).Error
	if err != nil {
		return nil, err
	}
	return hotels, nil
}