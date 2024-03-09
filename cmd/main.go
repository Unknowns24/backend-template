package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Unknowns24/backend-template/config"
	"github.com/Unknowns24/backend-template/internal/api"
	"github.com/go-chi/chi"
)

func main() {
	// Load app config
	config.Load()

	// Create new router with Go-Chi
	router := chi.NewRouter()

	// Configure defined routes
	api.SetupRoutes(router)

	// Start http server
	fmt.Printf("[INFO] Server started in port %s\n", strings.Replace(config.ENV.APP_PORT, ":", "", -1))

	if err := http.ListenAndServe(config.ENV.APP_PORT, router); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
