package util

import (
	"github.com/go-pg/pg"
	"fmt"
	"errors"
)

// MockStore ...
type MockStore struct {
	CreateUpdateCalledTimes     int
	CreateUpdateCalledWithName  string
	CreateUpdateCalledWithEmail string
	CreateUpdateError string
}

// CreateUpdate ...
func (m *MockStore) CreateUpdate(u User, db pg.DB) (string, error) {
	m.CreateUpdateCalledTimes++
	m.CreateUpdateCalledWithName = u.Name
	m.CreateUpdateCalledWithEmail = u.Email

	if m.CreateUpdateError != "" {
		return "internal server error", errors.New("internal server error")
	}

	msg := fmt.Sprintf("...successfully inserted or updated %v %v", u.Name, u.Email)

	return msg, nil
}

// Get ...
func (m *MockStore) Get(u *User, db pg.DB) (string, error) {
	return "", nil
}

// Delete ...
func (m *MockStore) Delete(u *User, db pg.DB) (string, error) {
	return "", nil
}
