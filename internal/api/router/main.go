package router

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/database"
	"net/http"
)

// RegisterRouter : main entrance of all routes (route groups)
func RegisterRouter(dbService *database.Service, r *http.ServeMux) {
	//__ 1.  USER ROUTES (auth + profile) __//
	r.Handle("/api/v1/auth/", RegisterAuthRoutes(dbService, r))
	r.Handle("/api/v1/user/", RegisterUserProfileRoutes(dbService, r))

	//__ 2. HOTEL __//
	r.Handle("/api/v1/hotel/", RegisterHotelRoutes(dbService, r))

	//__ 3. ROOMS __//
	r.Handle("/api/v1/room/", RegisterRoomRoutes(dbService, r))

	////__ 4. RESERVATIONS __//
	//r.Handle("/reservation/", RegisterHotelRoutes(dbService, r))
	//
	////__ 5. PAYMENTS __//
	//r.Handle("/payment/", RegisterHotelRoutes(dbService, r))
	//
	////__ 6. NOTIFICATIONS __//
	//r.Handle("/notification/", RegisterHotelRoutes(dbService, r))
}