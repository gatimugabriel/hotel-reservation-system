package main

import (
	routes "github.com/gatimugabriel/hotel-reservation-system/internal/api/router"
	"github.com/gatimugabriel/hotel-reservation-system/internal/config"
	"github.com/gatimugabriel/hotel-reservation-system/internal/infrastructure/database"
	"github.com/gatimugabriel/hotel-reservation-system/internal/server/httpServer"
	"net/http"
)

var configurations *config.Config

func init() {
	loadConfig, err := config.LoadConfig()
	if err != nil {
		panic("failed to load configuration")
	}
	configurations = loadConfig
}

func main() {
	// connect to DB
	dbService, err := database.NewDatabaseService(configurations)
	if err != nil {
		panic("failed to connect to database!")
	}

	//router setup
	router := http.NewServeMux()
	routes.RegisterRouter(dbService, router)

	// start server
	httpServer.StartServer(configurations, router)
}