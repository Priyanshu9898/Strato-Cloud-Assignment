package store

import "github.com/Priyanshu9898/Strato-Cloud-Assignment/internal/model"

// Store declares methods to fetch users.
type Store interface {
	GetAllUsers() ([]model.User, error)
}
