package api

import (
	"backend/internal/db"
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type HandlerRouter struct {
	router  *chi.Mux
	queries *db.Queries
}

func NewHandlerRouter(database *sql.DB) *HandlerRouter {
	r := chi.NewRouter()

	// Create new queries instance
	queries := db.New(database)

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Initialize router with database queries
	router := &HandlerRouter{
		router:  r,
		queries: queries,
	}

	RegisterRoutes(router, "/users")

	return router
}

func (r *HandlerRouter) GetChi() *chi.Mux {
	return r.router
}
