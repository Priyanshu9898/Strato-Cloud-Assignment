package internal

import (
	"net/http"
	"time"

	"github.com/Priyanshu9898/Strato-Cloud-Assignment/internal/handler"
	"github.com/Priyanshu9898/Strato-Cloud-Assignment/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter sets up routes and middleware.
func NewRouter() http.Handler {
	r := chi.NewRouter()

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
