package router

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/api/handlers"
	"github.com/gatimugabriel/hotel-reservation-system/internal/constants"
	reservationRepository "github.com/gatimugabriel/hotel-reservation-system/internal/domain/reservation/repository"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/repository"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/room/services"
	"github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/database"
	middleware2 "github.com/gatimugabriel/hotel-reservation-system/internal/middleware"
	"net/http"
)

// RegisterRoomRoutes registers room API endpoints
// @param db -> database service
// @param r -> http ServeMux (router)
// @return http.Handler
func RegisterRoomRoutes(db *database.Service, r *http.ServeMux) http.Handler {
	roomRepo := repository.NewRoomRepository(db)
	roomTypeRepo := repository.NewRoomTypeRepository(db)
	reservationRepo := reservationRepository.NewReservationRepository(db)

	roomService := services.NewRoomService(roomRepo, roomTypeRepo, reservationRepo)
	roomTypeService := services.NewRoomTypeService(roomTypeRepo)
	handler := handlers.NewRoomHandler(roomService, roomTypeService)

	allowedCreationRoles := []constants.Role{constants.MANAGER, constants.PROPERTYOWNER, constants.ADMIN}
	roleCheckMiddleware := middleware2.AuthWithRoleCheck(allowedCreationRoles)
	// apply to many handlers
	adminHandlers := middleware2.ApplyMiddlewareToMany(
		roleCheckMiddleware,
		handler.CreateRoom,
		handler.CreateRoomType,
		handler.CreateBedType,
	)

	//___Privileged routes ___//
	r.Handle("POST /create-room", adminHandlers[0])
	r.Handle("POST /create-type", adminHandlers[1])
	r.Handle("POST /create-bedtype", adminHandlers[2]) //bed types

	//r.Handle("POST /create-room",
	//	middleware.Authenticate(
	//		middleware.RoleCheck([]constants.Role{constants.MANAGER, constants.PROPERTYOWNER, constants.ADMIN},
	//			http.HandlerFunc(handler.CreateRoom)),
	//	),
	//)
	//r.Handle("POST /create-type",
	//	middleware.Authenticate(
	//		middleware.RoleCheck([]constants.Role{constants.MANAGER, constants.PROPERTYOWNER, constants.ADMIN},
	//			http.HandlerFunc(handler.CreateRoomType)),
	//	),
	//)

	//___Public routes ___//
	//1. rooms
	r.HandleFunc("GET /available", handler.GetAvailableRooms)
	r.HandleFunc("GET /all-rooms", handler.GetRooms)
	r.HandleFunc("GET /room-details/{roomID}", handler.GetRoom)

	//2. room types
	r.HandleFunc("GET /type-details/{roomTypeID}", handler.GetTypeDetails)
	r.HandleFunc("GET /type/all", handler.ListRoomTypes)

	return http.StripPrefix("/api/v1/room", r)
}