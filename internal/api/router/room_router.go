package router

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/api/handlers"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/repository"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/services"
	"github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/database"
	"net/http"
)

// RegisterRoomRoutes registers room API endpoints
// @param db -> database service
// @param r -> http ServeMux (router)
// @return http.Handler
func RegisterRoomRoutes(db *database.Service, r *http.ServeMux) http.Handler {
	roomRepo := repository.NewRoomRepository(db)
	roomService := services.NewRoomService(roomRepo)
	handler := handlers.NewRoomHandler(roomService)

	r.HandleFunc("POST /room/create", handler.CreateRoom)
	r.HandleFunc("GET /room/:id", handler.GetRoom)

	return http.StripPrefix("/api/v1/room", r)
}