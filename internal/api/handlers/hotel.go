package handlers

import (
	"encoding/json"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/hotel/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/hotel/services"
	"github.com/gatimugabriel/hotel-reservation-system/pkg/utils"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type HotelHandler struct {
	hotelService services.HotelService
}

func NewHotelHandler(hotelService services.HotelService) *HotelHandler {
	return &HotelHandler{
		hotelService: hotelService,
	}
}

func (h *HotelHandler) GetHotels(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	hotels, err := h.hotelService.GetHotels(r.Context())
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to get hotels")
		return
	}

	utils.RespondJSON(w, http.StatusOK, hotels)
}

func (h *HotelHandler) GetHotel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/api/v1/hotels/")
	hotel, err := h.hotelService.GetHotel(r.Context(), id)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, "Hotel not found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, hotel)
}

func (h *HotelHandler) CreateHotel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var hotelData entity.Hotel
	if err := json.NewDecoder(r.Body).Decode(&hotelData); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	createdHotel, err := h.hotelService.CreateHotel(r.Context(), &hotelData)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to create hotel")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, createdHotel)
}

func (h *HotelHandler) UpdateHotel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/hotels/")
	// Convert string to UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid hotel ID format")
		return
	}

	var hotelData entity.Hotel
	if err := json.NewDecoder(r.Body).Decode(&hotelData); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	updatedHotel, err := h.hotelService.UpdateHotel(r.Context(), id, &hotelData)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to update hotel")
		return
	}

	utils.RespondJSON(w, http.StatusOK, updatedHotel)
}

func (h *HotelHandler) DeleteHotel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/api/v1/hotels/")

	err := h.hotelService.DeleteHotel(r.Context(), id)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to delete hotel")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Hotel deleted successfully"})
}