package handlers

import (
	"encoding/json"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/services"
	"github.com/gatimugabriel/hotel-reservation-system/pkg/utils"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type RoomHandler struct {
	roomService services.RoomService
}

func NewRoomHandler(roomService services.RoomService) *RoomHandler {
	return &RoomHandler{
		roomService: roomService,
	}
}

func (h *RoomHandler) GetRooms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	hotelID := r.URL.Query().Get("hotel_id")
	rooms, err := h.roomService.GetRooms(r.Context(), hotelID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to get rooms")
		return
	}

	utils.RespondJSON(w, http.StatusOK, rooms)
}

func (h *RoomHandler) GetRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/api/v1/rooms/")
	room, err := h.roomService.GetRoom(r.Context(), id)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, "Room not found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, room)
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var roomData entity.Room
	if err := json.NewDecoder(r.Body).Decode(&roomData); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	createdRoom, err := h.roomService.CreateRoom(r.Context(), &roomData)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to create room")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, createdRoom)
}

func (h *RoomHandler) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/rooms/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid room ID")
		return
	}

	var roomData entity.Room
	if err := json.NewDecoder(r.Body).Decode(&roomData); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	updatedRoom, err := h.roomService.UpdateRoom(r.Context(), id, &roomData)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to update room")
		return
	}

	utils.RespondJSON(w, http.StatusOK, updatedRoom)
}

func (h *RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/rooms/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid room ID")
		return
	}

	err = h.roomService.DeleteRoom(r.Context(), id)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to delete room")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Room deleted successfully"})
}