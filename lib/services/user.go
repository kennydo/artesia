package services

import (
	"errors"
	"time"
)

// ErrUserNotFound is returned when we don't have a user satisfying whatever requirements
var ErrUserNotFound = errors.New("User not found")

// ErrUserEmailTaken is returned when a user's email is already taken
var ErrUserEmailTaken = errors.New("User email is taken")

// ErrUnableToCreateUser is returned when we're unable to create a user
var ErrUnableToCreateUser = errors.New("Unable to create user")

// User represents a local user that can login and do things
type User struct {
	ID           string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

// UserService exposes CRUD operations on Users
type UserService interface {
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	CreateUser(email string, plaintextPassword string) (*User, error)
}
