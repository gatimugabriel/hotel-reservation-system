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

	//__ 2. ROOMS __//
	r.Handle("/api/v1/room/", RegisterRoomRoutes(dbService, r))

	//__ 3. RESERVATIONS __//
	r.Handle("/api/v1/reservation/", RegisterReservationRoutes(dbService, r))

	////__ 4. PAYMENTS __//
	//r.Handle("/payment/", RegisterPaymentRoutes(dbService, r))
	//
	////__ 5. NOTIFICATIONS __//
	//r.Handle("/notification/", RegisterNotificationRoutes(dbService, r))
}