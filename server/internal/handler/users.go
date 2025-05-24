package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Priyanshu9898/Strato-Cloud-Assignment/internal/store"
)

// UsersHandler serves GET /api/users.
type UsersHandler struct {
	Store store.Store
}

// NewUsersHandler constructs a UsersHandler.
func NewUsersHandler(s store.Store) *UsersHandler {
	return &UsersHandler{Store: s}
}

// ServeHTTP writes the JSON array of users.
func (h *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	users, err := h.Store.GetAllUsers()
	if err != nil {
		http.Error(w, "failed to load users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
