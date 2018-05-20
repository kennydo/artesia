package services

import (
	"time"
)

// User represents a local user that can login and do things
type User struct {
	ID           int
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

// UserService exposes CRUD operations on Users
type UserService interface {
	GetByID(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	CreateUser(email string, plaintextPassword string) (*User, error)
}
