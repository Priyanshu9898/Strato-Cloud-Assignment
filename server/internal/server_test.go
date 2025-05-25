package internal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Priyanshu9898/Strato-Cloud-Assignment/internal/model"
)

func TestNewRouter_ServesUsersEndpoint(t *testing.T) {
	// Build the router
	router := NewRouter()

	// Make a GET /api/users request
	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 1) Status code
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", w.Code)
	}

	// 2) Content-Type header
	if ct := w.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", ct)
	}

	// 3) Body is valid JSON and has the expected number of users
	var users []model.User
	if err := json.Unmarshal(w.Body.Bytes(), &users); err != nil {
		t.Fatalf("failed to parse JSON response: %v", err)
	}

	const expectedCount = 10 // we seeded 10 users in MemoryStore
	if len(users) != expectedCount {
		t.Errorf("expected %d users, got %d", expectedCount, len(users))
	}

	// 4) Check one known user from the seed data
	found := false
	for _, u := range users {
		if u.Name == "John Doe" {
			found = true
			break
		}
	}
	if !found {
		t.Error("John Doe not found in response payload")
	}
}
