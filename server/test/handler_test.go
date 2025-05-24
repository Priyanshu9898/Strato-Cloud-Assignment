package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Priyanshu9898/Strato-Cloud-Assignment/internal/handler"
	"github.com/Priyanshu9898/Strato-Cloud-Assignment/internal/model"
	"github.com/Priyanshu9898/Strato-Cloud-Assignment/internal/store"
)

// helper to count how many users have MFA enabled in the seed data
func countMfaEnabled(t *testing.T, users []model.User) int {
	t.Helper()
	c := 0
	for _, u := range users {
		if u.MfaEnabled {
			c++
		}
	}
	return c
}

func TestUsersHandler_NoFilter(t *testing.T) {
	mem := store.NewMemoryStore()
	all, _ := mem.GetAllUsers()

	req := httptest.NewRequest("GET", "/api/users", nil)
	w := httptest.NewRecorder()
	handler.NewUsersHandler(mem).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", w.Code)
	}
	var resp []model.User
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if len(resp) != len(all) {
		t.Errorf("expected %d users, got %d", len(all), len(resp))
	}
}

func TestUsersHandler_MfaFilterTrue(t *testing.T) {
	mem := store.NewMemoryStore()
	all, _ := mem.GetAllUsers()
	wantCount := countMfaEnabled(t, all)

	req := httptest.NewRequest("GET", "/api/users?mfa=true", nil)
	w := httptest.NewRecorder()
	handler.NewUsersHandler(mem).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", w.Code)
	}
	var resp []model.User
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if len(resp) != wantCount {
		t.Errorf("expected %d MFA-enabled users, got %d", wantCount, len(resp))
	}
	for _, u := range resp {
		if !u.MfaEnabled {
			t.Errorf("got user with MFA=false in response: %+v", u)
		}
	}
}
