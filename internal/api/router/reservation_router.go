package router

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/api/handlers"
	"github.com/gatimugabriel/hotel-reservation-system/internal/api/middleware"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/reservation/repository"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/reservation/services"
	roomRepository "github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/repository"
	"github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/database"
	"net/http"
)

// RegisterReservationRoutes registers bookings API endpoints
// @param db -> database service
// @param r -> http ServeMux (router)
// @return http.Handler
func RegisterReservationRoutes(db *database.Service, r *http.ServeMux) http.Handler {
	reservationRepo := repository.NewReservationRepository(db)
	roomRepo := roomRepository.NewRoomRepository(db)

	reservationService := services.NewReservationService(reservationRepo, roomRepo)
	//roomService := roomService.NewRoomService(reservationRepo, roomRepo)
	handler := handlers.NewReservationHandler(reservationService)

	r.HandleFunc("POST /create-reservation", handler.CreateReservation)
	r.HandleFunc("PATCH /cancel/{reservationID}", handler.CancelReservation)
	r.HandleFunc("GET /me", handler.GetUserReservations)
	//r.HandleFunc("GET /{reservationID}", handler.GetReservation)

	return middleware.Authenticate(http.StripPrefix("/api/v1/reservation", r))
}