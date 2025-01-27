package router

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/api/handlers"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/hotel/repository"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/hotel/services"
	"github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/database"
	"net/http"
)

// RegisterHotelRoutes registers hotel API endpoints
// @param db -> database service
// @param r -> http ServeMux (router)
// @return http.Handler
func RegisterHotelRoutes(db *database.Service, r *http.ServeMux) http.Handler {
	hotelRepo := repository.NewHotelRepository(db)
	hotelService := services.NewHotelService(hotelRepo)
	handler := handlers.NewHotelHandler(hotelService)

	r.HandleFunc("POST /hotel/create", handler.CreateHotel)
	r.HandleFunc("GET /hotel/:id", handler.GetHotel)

	return http.StripPrefix("/api/v1/hotel", r)
}