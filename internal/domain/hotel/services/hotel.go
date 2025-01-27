package services

import (
	"context"
	"fmt"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/hotel/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/hotel/repository"
	"github.com/google/uuid"
	"time"
)

type HotelService interface {
	CreateHotel(ctx context.Context, hotel *entity.Hotel) (*entity.Hotel, error)
	UpdateHotel(ctx context.Context, id uuid.UUID, hotel *entity.Hotel) (*entity.Hotel, error)
	DeactivateHotel(ctx context.Context, id uuid.UUID) error
	GetHotels(ctx context.Context) ([]*entity.Hotel, error)
	GetHotelsByCity(ctx context.Context, city string) ([]*entity.Hotel, error)
	GetHotel(ctx context.Context, id uuid.UUID) (*entity.Hotel, error)
	DeleteHotel(ctx context.Context, id string) error
}
type HotelServiceImpl struct {
	hotelRepo repository.HotelRepository
}

func NewHotelService(hotelRepo repository.HotelRepository) *HotelServiceImpl {
	return &HotelServiceImpl{
		hotelRepo: hotelRepo,
	}
}

func (h *HotelServiceImpl) CreateHotel(ctx context.Context, hotel *entity.Hotel) (*entity.Hotel, error) {
	//get manager id from context
	ownerID := ctx.Value("userID").(string)
	hotel.OwnerID, _ = uuid.Parse(ownerID)

	if err := h.hotelRepo.Create(ctx, hotel); err != nil {
		return nil, fmt.Errorf("failed to create hotel: %w", err)
	}
	return hotel, nil
}

func (h *HotelServiceImpl) UpdateHotel(ctx context.Context, id uuid.UUID, hotel *entity.Hotel) (*entity.Hotel, error) {
	existingHotel, err := h.hotelRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get hotel: %w", err)
	}

	hotel.ID = existingHotel.ID
	hotel.UpdatedAt = time.Now()

	if err := h.hotelRepo.Update(ctx, hotel); err != nil {
		return nil, fmt.Errorf("failed to update hotel: %w", err)
	}
	return hotel, nil
}

func (h *HotelServiceImpl) DeactivateHotel(ctx context.Context, id uuid.UUID) error {
	hotel, err := h.hotelRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get hotel: %w", err)
	}

	hotel.IsActive = false
	hotel.UpdatedAt = time.Now()

	if err := h.hotelRepo.Update(ctx, hotel); err != nil {
		return fmt.Errorf("failed to deactivate hotel: %w", err)
	}
	return nil
}

func (h *HotelServiceImpl) GetHotels(ctx context.Context) ([]*entity.Hotel, error) {
	hotels, err := h.hotelRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get hotels: %w", err)
	}
	return hotels, nil
}

func (h *HotelServiceImpl) GetHotelsByCity(ctx context.Context, city string) ([]*entity.Hotel, error) {
	hotels, err := h.hotelRepo.GetByCity(ctx, city)
	if err != nil {
		return nil, fmt.Errorf("failed to get hotels by city: %w", err)
	}
	return hotels, nil
}

func (h *HotelServiceImpl) GetHotel(ctx context.Context, id uuid.UUID) (*entity.Hotel, error) {
	hotel, err := h.hotelRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get hotel: %w", err)
	}
	return hotel, nil
}

func (h *HotelServiceImpl) DeleteHotel(ctx context.Context, id string) error {
	hotelID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid hotel ID: %w", err)
	}

	if err := h.hotelRepo.Delete(ctx, hotelID); err != nil {
		return fmt.Errorf("failed to delete hotel: %w", err)
	}
	return nil
}