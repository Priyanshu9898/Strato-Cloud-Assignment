package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Priyanshu9898/Strato-Cloud-Assignment/internal/model"
	"github.com/Priyanshu9898/Strato-Cloud-Assignment/internal/store"
)

// UsersHandler serves GET /api/users
type UsersHandler struct {
	Store store.Store
}

// NewUsersHandler creates a UsersHandler
func NewUsersHandler(s store.Store) *UsersHandler {
	return &UsersHandler{Store: s}
}

func (h *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	users, err := h.Store.GetAllUsers()
	if err != nil {
		http.Error(w, "failed to load users", http.StatusInternalServerError)
		return
	}
	// Check for mfa query param
	if mfaParam := r.URL.Query().Get("mfa"); mfaParam != "" {
		want, err := strconv.ParseBool(mfaParam)
		if err != nil {
			http.Error(w, "invalid mfa filter", http.StatusBadRequest)
			return
		}
		filtered := make([]model.User, 0, len(users))
		for _, u := range users {
			if u.MfaEnabled == want {
				filtered = append(filtered, u)
			}
		}
		users = filtered
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
