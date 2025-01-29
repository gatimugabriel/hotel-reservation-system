package repository

import (
	"context"
	"errors"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/database"
	"log"
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

// RoomRepository : data persistence database interaction interface
type RoomRepository interface {
	GetAvailableRooms(ctx context.Context, hotelID uuid.UUID, checkIn, checkOut time.Time) ([]*entity.Room, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Room, error)
	GetByNumber(ctx context.Context, roomNumber int) (*entity.Room, error)
	GetByHotelID(ctx context.Context, hotelID uuid.UUID) ([]*entity.Room, error)

	GetReservationsForDate(ctx context.Context, date time.Time) ([]*entity.Room, error)
	GetReservationsForDateRange(ctx context.Context, in time.Time, out time.Time) ([]*entity.Room, error)

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

func (repo *RoomRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*entity.Room, error) {
	var room entity.Room
	err := repo.db.WithContext(ctx).
		Preload("Hotel").
		Preload("RoomType").
		First(&room, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	// Map hotel fields
	room.HotelInfo = entity.HotelLimited{
		Name:          room.Hotel.Name,
		ContactNumber: room.Hotel.ContactNumber,
	}

	//Map room type fields
	room.RoomTypeInfo = entity.RoomTypeLimited{
		Name:         room.RoomType.Name,
		Description:  room.RoomType.Description,
		BasePrice:    room.RoomType.BasePrice,
		MaxOccupancy: room.RoomType.MaxOccupancy,
		NumBeds:      room.RoomType.NumBeds,
		BedType:      room.RoomType.BedType,
	}

	return &room, nil
}

func (repo *RoomRepositoryImpl) GetByNumber(ctx context.Context, roomNumber int) (*entity.Room, error) {
	var room entity.Room
	err := repo.db.WithContext(ctx).
		Preload("Hotel").
		Preload("RoomType").
		First(&room, "room_number = ?", roomNumber).Error
	if err != nil {
		log.Println("db error", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	// Map hotel fields
	room.HotelInfo = entity.HotelLimited{
		Name:          room.Hotel.Name,
		ContactNumber: room.Hotel.ContactNumber,
	}

	//Map room type fields
	room.RoomTypeInfo = entity.RoomTypeLimited{
		Name:         room.RoomType.Name,
		Description:  room.RoomType.Description,
		BasePrice:    room.RoomType.BasePrice,
		MaxOccupancy: room.RoomType.MaxOccupancy,
		NumBeds:      room.RoomType.NumBeds,
		BedType:      room.RoomType.BedType,
	}

	return &room, nil
}

func (repo *RoomRepositoryImpl) GetByHotelID(ctx context.Context, hotelID uuid.UUID) ([]*entity.Room, error) {
	var responses []*entity.Room

	err := repo.db.WithContext(ctx).
		Preload("Hotel").
		Preload("RoomType").
		Where("hotel_id = ?", hotelID).
		Find(&responses).Error

	if err != nil {
		return nil, err
	}

	// Map the full objects to limited ones
	for _, room := range responses {
		// Map hotel fields
		room.HotelInfo = entity.HotelLimited{
			Name:          room.Hotel.Name,
			ContactNumber: room.Hotel.ContactNumber,
		}

		// Map room type fields
		room.RoomTypeInfo = entity.RoomTypeLimited{
			Name:         room.RoomType.Name,
			Description:  room.RoomType.Description,
			BasePrice:    room.RoomType.BasePrice,
			MaxOccupancy: room.RoomType.MaxOccupancy,
			NumBeds:      room.RoomType.NumBeds,
			BedType:      room.RoomType.BedType,
			SquareMeters: room.RoomType.SquareMeters,
		}
	}

	return responses, nil
}

func (repo *RoomRepositoryImpl) Create(ctx context.Context, r *entity.Room) error {
	return repo.db.Transaction(ctx, func(tx *gorm.DB) error {
		return tx.Create(r).Error
	})
}

func (repo *RoomRepositoryImpl) GetAvailableRooms(ctx context.Context, hotelID uuid.UUID, checkIn, checkOut time.Time) ([]*entity.Room, error) {
	var rooms []*entity.Room

	err := repo.db.WithContext(ctx).
		Preload("Hotel").
		Preload("RoomType").
		Where("hotel_id = ? AND is_available = ? AND is_maintenance = ?", hotelID, true, false).
		Find(&rooms).Error
	if err != nil {
		return nil, err
	}

	// Map the full objects to limited ones
	for _, room := range rooms {
		// Map hotel fields
		room.HotelInfo = entity.HotelLimited{
			Name:          room.Hotel.Name,
			ContactNumber: room.Hotel.ContactNumber,
		}

		// Map room type fields
		room.RoomTypeInfo = entity.RoomTypeLimited{
			Name:         room.RoomType.Name,
			Description:  room.RoomType.Description,
			BasePrice:    room.RoomType.BasePrice,
			MaxOccupancy: room.RoomType.MaxOccupancy,
			NumBeds:      room.RoomType.NumBeds,
			BedType:      room.RoomType.BedType,
			SquareMeters: room.RoomType.SquareMeters,
		}
	}

	return rooms, nil
}

func (repo *RoomRepositoryImpl) UpdateAvailability(ctx context.Context, id uuid.UUID, isAvailable bool) error {
	return repo.db.Transaction(ctx, func(tx *gorm.DB) error {
		result := tx.Model(&entity.Room{}).Where("id = ?", id).Update("is_available", isAvailable)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("room not found")
		}
		return nil
	})
}

func (repo *RoomRepositoryImpl) GetReservationsForDate(ctx context.Context, in time.Time) ([]*entity.Room, error) {
	var rooms []*entity.Room

	err := repo.db.WithContext(ctx).
		Where("check_in", in).
		Find(&rooms).Error

	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (repo *RoomRepositoryImpl) GetReservationsForDateRange(ctx context.Context, in time.Time, out time.Time) ([]*entity.Room, error) {
	var rooms []*entity.Room

	err := repo.db.WithContext(ctx).
		Where("check_in <= ? AND check_out >= ?", out, in).
		Find(&rooms).Error

	if err != nil {
		return nil, err
	}

	return rooms, nil
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