package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/services"
	"github.com/gatimugabriel/hotel-reservation-system/pkg/utils"
	"github.com/gatimugabriel/hotel-reservation-system/pkg/utils/input"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"net/http"
	"regexp"
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
	hotelID := r.URL.Query().Get("hotel_id")
	rooms, err := h.roomService.GetRooms(r.Context(), hotelID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to get rooms")
		return
	}

	roomsCount := len(rooms)
	utils.RespondPaginatedJSON(w, http.StatusOK, rooms, roomsCount, 0000, 0000)
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

		errorMessage := ""
		errorStatus := http.StatusInternalServerError

		// Check for unique constraint violation
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				//  Extract column name and value for the detail message
				re := regexp.MustCompile(`Key \(([^)]+)\)=\([^)]*\) already exists.`)
				matches := re.FindStringSubmatch(pgErr.Detail)
				if len(matches) == 2 {
					columnName := matches[1]
					errorMessage = fmt.Sprintf("%s already exists", columnName)
					errorStatus = http.StatusConflict
				} else {
					errorMessage = fmt.Sprintf("%s already exists", pgErr.ConstraintName)
					errorStatus = http.StatusConflict
				}

				utils.RespondError(w, errorStatus, errorMessage)
				return
			}
		}

		// default
		utils.RespondError(w, http.StatusInternalServerError, "Failed to create room")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, createdRoom)
}

//func (h *RoomHandler) UpdateRoom(w http.ResponseWriter, r *http.Request) {
//	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/rooms/")
//	id, err := uuid.Parse(idStr)
//	if err != nil {
//		utils.RespondError(w, http.StatusBadRequest, "Invalid room ID")
//		return
//	}
//
//	var roomData entity.Room
//	if err := json.NewDecoder(r.Body).Decode(&roomData); err != nil {
//		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
//		return
//	}
//
//	updatedRoom, err := h.roomService.UpdateRoom(r.Context(), id, &roomData)
//	if err != nil {
//		utils.RespondError(w, http.StatusInternalServerError, "Failed to update room")
//		return
//	}
//
//	utils.RespondJSON(w, http.StatusOK, updatedRoom)
//}
//
//func (h *RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
//	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/rooms/")
//	id, err := uuid.Parse(idStr)
//	if err != nil {
//		utils.RespondError(w, http.StatusBadRequest, "Invalid room ID")
//		return
//	}
//
//	err = h.roomService.DeleteRoom(r.Context(), id)
//	if err != nil {
//		utils.RespondError(w, http.StatusInternalServerError, "Failed to delete room")
//		return
//	}
//
//	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Room deleted successfully"})
//}

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
	}

	createdRoomType, err := h.roomTypeService.CreateRoomType(r.Context(), &roomData)
	if err != nil {
		errorMessage := ""
		errorStatus := http.StatusInternalServerError

		// Check for unique constraint violation
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				//  Extract column name and value for the detail message
				re := regexp.MustCompile(`Key \(([^)]+)\)=\([^)]*\) already exists.`)
				matches := re.FindStringSubmatch(pgErr.Detail)
				if len(matches) == 2 {
					columnName := matches[1]
					errorMessage = fmt.Sprintf("%s already exists", columnName)
					errorStatus = http.StatusConflict
				} else {
					errorMessage = fmt.Sprintf("%s already exists", pgErr.ConstraintName)
					errorStatus = http.StatusConflict
				}

				utils.RespondError(w, errorStatus, errorMessage)
				return
			}
		}

		// default
		utils.RespondError(w, http.StatusInternalServerError, "Failed to create room type")
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