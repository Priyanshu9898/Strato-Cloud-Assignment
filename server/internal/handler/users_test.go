package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Priyanshu9898/Strato-Cloud-Assignment/internal/model"
	"github.com/Priyanshu9898/Strato-Cloud-Assignment/internal/store"
)

func TestUsersHandler_NoFilter(t *testing.T) {
	ms := store.NewMemoryStore()
	allUsers, _ := ms.GetAllUsers()

	req := httptest.NewRequest("GET", "/api/users", nil)
	w := httptest.NewRecorder()

	NewUsersHandler(ms).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", ct)
	}

	var got []model.User
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}
	if len(got) != len(allUsers) {
		t.Errorf("expected %d users, got %d", len(allUsers), len(got))
	}
}

func TestUsersHandler_MfaFilter(t *testing.T) {
	ms := store.NewMemoryStore()
	allUsers, _ := ms.GetAllUsers()

	// count how many in seed have MfaEnabled=true
	wantCount := 0
	for _, u := range allUsers {
		if u.MfaEnabled {
			wantCount++
		}
	}

	req := httptest.NewRequest("GET", "/api/users?mfa=true", nil)
	w := httptest.NewRecorder()

	NewUsersHandler(ms).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var got []model.User
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}
	if len(got) != wantCount {
		t.Errorf("expected %d MFA-enabled users, got %d", wantCount, len(got))
	}
	for _, u := range got {
		if !u.MfaEnabled {
			t.Errorf("user with MFA=false found in filtered result: %+v", u)
		}
	}
}
