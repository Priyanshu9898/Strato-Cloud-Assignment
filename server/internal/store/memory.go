package store

import "github.com/Priyanshu9898/Strato-Cloud-Assignment/internal/model"

// MemoryStore holds a hard-coded slice of Users.
type MemoryStore struct {
	data []model.User
}

// NewMemoryStore seeds the five example users.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: []model.User{
		{
			Name:                "Foo Bar1",
			CreateDate:          "2020-10-01",
			PasswordChangedDate: "2021-10-01",
			LastAccessDate:      "2025-01-04",
			MfaEnabled:          true,
		},
		{
			Name:                "Foo1 Bar1",
			CreateDate:          "2019-09-20",
			PasswordChangedDate: "2019-09-22",
			LastAccessDate:      "2025-02-08",
			MfaEnabled:          false,
		},
		{
			Name:                "Foo2 Bar2",
			CreateDate:          "2022-02-03",
			PasswordChangedDate: "2022-02-03",
			LastAccessDate:      "2025-02-12",
			MfaEnabled:          false,
		},
		{
			Name:                "Foo3 Bar3",
			CreateDate:          "2023-03-07",
			PasswordChangedDate: "2023-03-10",
			LastAccessDate:      "2022-01-03",
			MfaEnabled:          true,
		},
		{
			Name:                "Foo Bar4",
			CreateDate:          "2018-04-08",
			PasswordChangedDate: "2020-04-12",
			LastAccessDate:      "2022-10-04",
			MfaEnabled:          false,
		},
		{
			Name:                "John Doe",
			CreateDate:          "2021-06-15",
			PasswordChangedDate: "2024-05-10",
			LastAccessDate:      "2025-03-20",
			MfaEnabled:          true,
		},
		{
			Name:                "Jane Roe",
			CreateDate:          "2020-12-01",
			PasswordChangedDate: "2021-12-01",
			LastAccessDate:      "2022-11-30",
			MfaEnabled:          false,
		},
		{
			Name:                "Test User",
			CreateDate:          "2019-07-22",
			PasswordChangedDate: "2023-01-10",
			LastAccessDate:      "2025-05-01",
			MfaEnabled:          true,
		},
		{
			Name:                "Alpha Beta",
			CreateDate:          "2018-02-14",
			PasswordChangedDate: "2019-02-14",
			LastAccessDate:      "2021-02-14",
			MfaEnabled:          false,
		},
		{
			Name:                "Gamma Delta",
			CreateDate:          "2022-11-11",
			PasswordChangedDate: "2023-11-11",
			LastAccessDate:      "2025-04-30",
			MfaEnabled:          true,
		},
	}}
}

// GetAllUsers returns all seeded users.
func (m *MemoryStore) GetAllUsers() ([]model.User, error) {
	return m.data, nil
}
