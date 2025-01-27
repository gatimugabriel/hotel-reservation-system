package handlers

import (
	"encoding/json"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/reservation/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/reservation/services"
	"github.com/gatimugabriel/hotel-reservation-system/pkg/utils"
	"github.com/google/uuid"
	"io"
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

type CreateReservationRequest struct {
	RoomID       uuid.UUID `json:"room_id"`
	UserID       uuid.UUID `json:"user_id"`
	CheckInDate  time.Time `json:"check_in_date"`
	CheckoutDate time.Time `json:"check_out_date"`
}

func (h *ReservationHandler) CreateReservation(w http.ResponseWriter, r *http.Request) {
	var req CreateReservationRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Date validations
	if req.CheckInDate.After(req.CheckoutDate) {
		utils.RespondError(w, http.StatusBadRequest, "Start date must be before end date")
		return
	}

	if req.CheckInDate.Before(time.Now()) {
		utils.RespondError(w, http.StatusBadRequest, "Start date cannot be in the past")
		return
	}

	newReservation := &entity.Reservation{
		RoomID:       req.RoomID,
		UserID:       req.UserID,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckoutDate,
		Status:       "pending",
	}

	createdReservation, err := h.reservationService.CreateReservation(r.Context(), newReservation)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, createdReservation)
}

func (h *ReservationHandler) GetUserReservations(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("userId")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	reservations, err := h.reservationService.GetUserReservations(r.Context(), userID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, reservations)
}

func (h *ReservationHandler) UpdateReservationStatus(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid reservation ID")
		return
	}

	var updateReq struct {
		Status string `json:"status"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &updateReq); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Add status validation if needed
	if updateReq.Status == "" {
		utils.RespondError(w, http.StatusBadRequest, "Status is required")
		return
	}

	updatedReservation, err := h.reservationService.UpdateReservation(r.Context(), id)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, updatedReservation)
}

func (h *ReservationHandler) CancelReservation(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid reservation ID")
		return
	}

	if err := h.reservationService.CancelReservation(r.Context(), id); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Reservation cancelled successfully"})
}