package router

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/api/handlers"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/repository"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/services"
	"github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/database"
	"net/http"
)

// RegisterAuthRoutes registers auth related API endpoints
// @param db -> database service
// @param r -> http ServeMux (router)
// @return http.Handler
func RegisterAuthRoutes(db *database.Service, r *http.ServeMux) http.Handler {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	handler := handlers.NewUserHandler(userService)

	r.HandleFunc("POST /signup", handler.SignUp)
	r.HandleFunc("POST /signin", handler.Login) // TODO -> add constraints; 5 signin attempts
	//r.HandleFunc("POST /signout", tokenHandler.SignOut)
	//r.HandleFunc("POST /refresh", handler.RefreshToken)
	//
	//// Google Auth
	//r.HandleFunc("GET /google/login", handler.GoogleLogin)
	//r.HandleFunc("GET /google/callback", handler.GoogleCallback)

	return http.StripPrefix("/api/v1/auth", r)
}