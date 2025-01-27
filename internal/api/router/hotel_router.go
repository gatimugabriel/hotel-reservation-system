package router

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/api/handlers"
	"github.com/gatimugabriel/hotel-reservation-system/internal/api/middleware"
	"github.com/gatimugabriel/hotel-reservation-system/internal/constants"
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

	r.Handle("POST /create-hotel",
		middleware.Authenticate(
			middleware.RoleCheck([]constants.Role{constants.PROPERTYOWNER},
				http.HandlerFunc(handler.CreateHotel)),
		),
	)

	r.HandleFunc("GET /{hotelID}", handler.GetHotel)
	r.HandleFunc("GET /all", handler.GetHotels)

	return http.StripPrefix("/api/v1/hotel", r)
}