package api

import (
	"github.com/Unknowns24/backend-template/internal/api/handlers"
	"github.com/Unknowns24/backend-template/pkg/middlewares"
	"github.com/go-chi/chi"
)

// SetupRoutes configures the API routes
func SetupRoutes(router *chi.Mux) {
	// Route /foo
	router.With(middlewares.PrintMiddleware).Get("/foo", handlers.HandleFoo)

	// Route /bar
	router.Post("/bar", handlers.HandleBar)
}
