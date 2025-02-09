package httpServer

import (
	"github.com/gatimugabriel/hotel-reservation-system/internal/config"
	middleware2 "github.com/gatimugabriel/hotel-reservation-system/internal/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func StartServer(configurations *config.Config, router *http.ServeMux) {
	address := ":" + configurations.Server.Port

	//____ apply global middlewares ____//
	// 1. rate limiter
	rateLimiter := middleware2.NewIPRateLimiter(5, 10)
	limitedRouter := middleware2.Limit(rateLimiter)(router)

	//	2.Request Logger
	wrappedRouter := middleware2.RequestLogger(limitedRouter)

	//____ CORS Setup ____//
	corsHandler := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowedMethods:   []string{"POST", "GET", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD", "CONNECT", "TRACE"},
		AllowedOrigins:   configurations.Server.AllowedOrigins,
	}).Handler(wrappedRouter)

	//___ start server ____//
	log.Println("Server listening on port", address)
	if err := http.ListenAndServe(address, corsHandler); err != nil {
		log.Fatalf("Error starting server %v", err)
	}
}