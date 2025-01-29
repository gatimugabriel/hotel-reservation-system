package router

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/api/handlers"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/repository"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/services"
	"github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/database"
	"net/http"
)

// RegisterUserProfileRoutes registers User API endpoints
// @param db -> database service
// @param r -> http ServeMux (router)
// @return http.Handler
func RegisterUserProfileRoutes(db *database.Service, r *http.ServeMux) http.Handler {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	handler := handlers.NewUserHandler(userService)

	r.HandleFunc("GET /:id", handler.GetUserProfile)
	//r.HandleFunc("PUT /:id", handler.UpdateUserProfile)
	//r.HandleFunc("DELETE /:id", handler.DeleteUser)
	//r.HandleFunc("GET /all", handler.GetUsers)

	return http.StripPrefix("/api/v1/user", r)
}