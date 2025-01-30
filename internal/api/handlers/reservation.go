package handlers

import (
	"encoding/json"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/reservation/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/reservation/services"
	"github.com/gatimugabriel/hotel-reservation-system/pkg/utils"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type ReservationHandler struct {
	reservationService services.ReservationService
}

func NewReservationHandler(reservationService services.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		reservationService: reservationService,
	}
}

func (h *ReservationHandler) CreateReservation(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Context().Value("userID").(string)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var roomID uuid.UUID
	var totalPrice float64

	var req entity.CreateReservationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if req.RoomID == nil && req.RoomNumber == nil {
		utils.RespondError(w, http.StatusBadRequest, "Either room_id or room_number must be provided")
		return
	}

	checkInDate, err := time.Parse("2006-01-02", req.CheckInDate)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid check-in date format")
		return
	}

	checkOutDate, err := time.Parse("2006-01-02", req.CheckoutDate)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid check-out date format")
		return
	}

	//if room id is not provided, get room id from the room_number given in body
	if req.RoomID == nil {
		room, err := h.reservationService.GetRoomByNumber(r.Context(), *req.RoomNumber)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		roomID = room.ID
		totalPrice = float64(req.NumGuests) * room.RoomTypeInfo.BasePrice
	} else {
		id, err := uuid.Parse(*req.RoomID)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid room ID")
			return
		}
		roomID = id
	}

	newReservation := &entity.Reservation{
		RoomID:         roomID,
		UserID:         userID,
		CheckInDate:    checkInDate,
		CheckOutDate:   checkOutDate,
		NumGuests:      req.NumGuests,
		SpecialRequest: *req.SpecialRequest,
		TotalPrice:     totalPrice,
	}

	createdReservation, err := h.reservationService.CreateReservation(r.Context(), newReservation)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, createdReservation)
}

func (h *ReservationHandler) CancelReservation(w http.ResponseWriter, r *http.Request) {
	idStr := utils.GetResourceIDFromURL(r)
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid reservation ID")
		return
	}

	reservation, err := h.reservationService.GetReservation(r.Context(), id)
	if err != nil || reservation == nil {
		utils.RespondError(w, http.StatusInternalServerError, "failed to get reservation: ")
		return
	}
	reservation.Status = entity.StatusCancelled

	updated, err := h.reservationService.UpdateReservation(r.Context(), reservation)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "failed to cancel reservation: "+err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]any{"message": "Reservation cancelled successfully", "updated": updated})
}

func (h *ReservationHandler) GetUserReservations(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Context().Value("userID").(string)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Invalid user ID")
		return
	}

	reservations, err := h.reservationService.GetUserReservations(r.Context(), userID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, reservations)
}

func (h *ReservationHandler) GetReservation(w http.ResponseWriter, r *http.Request) {
	reservationID := utils.GetResourceIDFromURL(r)
	reservationIDParsed, err := uuid.Parse(reservationID)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid reservation ID")
		return
	}

	reservation, err := h.reservationService.GetReservation(r.Context(), reservationIDParsed)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, reservation)
}