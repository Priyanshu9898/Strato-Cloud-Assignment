package store

import (
	"testing"
)

func TestMemoryStore_GetAllUsers(t *testing.T) {
	ms := NewMemoryStore()
	users, err := ms.GetAllUsers()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(users) == 0 {
		t.Fatalf("expected at least one user, got %d", len(users))
	}

	// Spot‐check that seeded data contains exactly 10 users
	const expected = 10
	if len(users) != expected {
		t.Errorf("expected %d users, got %d", expected, len(users))
	}
}
