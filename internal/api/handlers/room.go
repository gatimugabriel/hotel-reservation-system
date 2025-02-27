package handlers

import (
	"encoding/json"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/services"
	"github.com/gatimugabriel/hotel-reservation-system/pkg/utils"
	"github.com/gatimugabriel/hotel-reservation-system/pkg/utils/input"
	"github.com/google/uuid"
	"net/http"
)

type RoomHandler struct {
	roomService     services.RoomService
	roomTypeService services.RoomTypeService
}

func NewRoomHandler(roomService services.RoomService, roomTypeService services.RoomTypeService) *RoomHandler {
	return &RoomHandler{
		roomService:     roomService,
		roomTypeService: roomTypeService,
	}
}

func (h *RoomHandler) GetRooms(w http.ResponseWriter, r *http.Request) {
	filters := map[string]interface{}{} //none

	rooms, err := h.roomService.GetRooms(r.Context(), filters)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to get rooms")
		return
	}

	roomsCount := len(rooms)
	utils.RespondPaginatedJSON(w, http.StatusOK, rooms, roomsCount, 0000, 0000)
}

func (h *RoomHandler) GetAvailableRooms(w http.ResponseWriter, r *http.Request) {
	// get dates from query params
	checkinDate, err := utils.GetDateFromURL(r, "check_in")
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid checkin date: "+err.Error())
		return
	}
	checkoutDate, err := utils.GetDateFromURL(r, "check_out")
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid checkout date: "+err.Error())
		return
	}

	// Validate check-in/check-out relationship
	if err := utils.ValidateCheckInCheckOutDates(checkinDate, checkoutDate); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	categorizedRooms, err := h.roomService.CheckAvailability(r.Context(), checkinDate, checkoutDate)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to get available rooms:"+err.Error())
		return
	}
	if categorizedRooms == nil {
		utils.RespondError(w, http.StatusInternalServerError, "No available rooms")
		return
	}

	utils.RespondJSON(w, http.StatusOK, categorizedRooms)
}

func (h *RoomHandler) GetRoom(w http.ResponseWriter, r *http.Request) {
	idStr := utils.GetResourceIDFromURL(r)
	room, err := h.roomService.GetRoom(r.Context(), idStr)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, "Room not found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, room)
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var roomData entity.Room
	if err := json.NewDecoder(r.Body).Decode(&roomData); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	createdRoom, err := h.roomService.CreateRoom(r.Context(), &roomData)
	if err != nil {
		if status, message, ok := utils.HandleUniqueConstraintError(err); ok {
			utils.RespondError(w, status, message)
			return
		}

		// default
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, createdRoom)
}

// ___ Room Types____//

// CreateRoomType creates a new room type
func (h *RoomHandler) CreateRoomType(w http.ResponseWriter, r *http.Request) {
	var roomData entity.RoomType
	if err := json.NewDecoder(r.Body).Decode(&roomData); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	//Sanitize & ValidateStruct
	if validationErrors := input.ValidateStruct(roomData); validationErrors != nil {
		utils.RespondJSON(w, http.StatusBadRequest, validationErrors)
		return
	}

	createdRoomType, err := h.roomTypeService.CreateRoomType(r.Context(), &roomData)
	if err != nil {
		if status, message, ok := utils.HandleUniqueConstraintError(err); ok {
			utils.RespondError(w, status, message)
			return
		}
		// default
		utils.RespondError(w, http.StatusInternalServerError, "Failed to create room type:"+err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, createdRoomType)
}

// GetTypeDetails retrieves details of a specific room type
func (h *RoomHandler) GetTypeDetails(w http.ResponseWriter, r *http.Request) {
	idStr := utils.GetResourceIDFromURL(r)
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid room type ID")
		return
	}

	roomType, err := h.roomTypeService.GetRoomType(r.Context(), id)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, "Room type not found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, roomType)
}

// ListRoomTypes retrieves a list of all room types
func (h *RoomHandler) ListRoomTypes(w http.ResponseWriter, r *http.Request) {
	roomTypes, err := h.roomTypeService.ListRoomTypes(r.Context())
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to retrieve room types")
		return
	}

	utils.RespondJSON(w, http.StatusOK, roomTypes)
}

// CreateBedType creates a new bed type
func (h *RoomHandler) CreateBedType(w http.ResponseWriter, r *http.Request) {
	var bedType entity.BedType
	if err := json.NewDecoder(r.Body).Decode(&bedType); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	//Sanitize & ValidateStruct
	if validationErrors := input.ValidateStruct(bedType); validationErrors != nil {
		utils.RespondJSON(w, http.StatusBadRequest, validationErrors)
		return
	}

	createdBedType, err := h.roomTypeService.CreateBedType(r.Context(), &bedType)
	if err != nil {
		if status, message, ok := utils.HandleUniqueConstraintError(err); ok {
			utils.RespondError(w, status, message)
			return
		}
		// default
		utils.RespondError(w, http.StatusInternalServerError, "Failed to create bed type:"+err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, createdBedType)
}