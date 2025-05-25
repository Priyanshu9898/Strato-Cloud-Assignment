package internal

import (
	"net/http"
	"time"

	"github.com/Priyanshu9898/Strato-Cloud-Assignment/internal/handler"
	"github.com/Priyanshu9898/Strato-Cloud-Assignment/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// NewRouter sets up routes and middleware.
func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// local dev, Vercel Preview, and Production frontend
		AllowedOrigins: []string{
			"http://localhost:5173",                                // Vite dev server
			"https://strato-cloud-assignment-8hthq7cn7.vercel.app", // Vercel preview
			"https://strato-cloud-assignment.vercel.app",           // Vercel production
		},
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	// Standard middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(5 * time.Second))

	// Wire up our in-memory store + handler
	ms := store.NewMemoryStore()
	r.Get("/api/users", handler.NewUsersHandler(ms).ServeHTTP)

	return r
}
