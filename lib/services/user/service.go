package user

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kennydo/artesia/lib/services"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// DBService implements the user service, backed by the DB
type DBService struct {
	log *zap.SugaredLogger
}

// GetByID gets a user by ID. Returns an error if no user is found.
func (s *DBService) GetByID(ctx context.Context, tx *sqlx.Tx, id string) (*services.User, error) {
	dbUser := DBUser{}
	err := tx.GetContext(ctx, &dbUser, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, services.ErrUserNotFound
	}

	user := services.User(dbUser)
	return &user, nil
}

// GetByEmail gets a user by their (case-insensitive) email. Returns an error if no user is found.
func (s *DBService) GetByEmail(ctx context.Context, tx *sqlx.Tx, email string) (*services.User, error) {
	dbUser := DBUser{}
	err := tx.GetContext(ctx, &dbUser, "SELECT * FROM users WHERE lower(email) = lower($1)", email)
	if err != nil {
		return nil, services.ErrUserNotFound
	}

	user := services.User(dbUser)
	return &user, nil
}

// CreateUser creates a user. Returns an error if unable to create user.
func (s *DBService) CreateUser(ctx context.Context, tx *sqlx.Tx, email string, plaintextPassword string) (*services.User, error) {
	// We don't want to create a user if there's already one with the same email
	existingUser, err := s.GetByEmail(ctx, tx, email)
	if existingUser != nil || err == nil {
		return nil, services.ErrUserEmailTaken
	}

	passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)
	if err != nil {
		s.log.Infow("Unable to hash password", zap.Error(err), zap.String("email", email))
		return nil, fmt.Errorf("Invalid password")
	}

	stmt, err := tx.PrepareNamedContext(
		ctx,
		`INSERT INTO users (
			id,
			email,
			password_hash,
			created_at
		) VALUES (
			:id,
			:email,
			:password_hash,
			:created_at
		)
		`)
	if err != nil {
		s.log.Infow("Unable to prepare insert statement", zap.Error(err))
		return nil, fmt.Errorf("Unable to prepare user for insert")
	}

	newUser := DBUser{
		ID:           uuid.NewV4().String(),
		Email:        email,
		PasswordHash: string(passwordHashBytes),
		CreatedAt:    time.Now().In(time.UTC),
	}

	result, err := stmt.ExecContext(ctx, &newUser)
	if err != nil {
		s.log.Infow("Unable to insert user into DB", zap.Error(err), zap.String("email", email))
		return nil, services.ErrUnableToCreateUser
	}

	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		s.log.Infow("Unable to get newly inserted row", zap.Error(err), zap.String("id", newUser.ID))
		return nil, services.ErrUnableToCreateUser
	}

	if numRowsAffected != 1 {
		s.log.Infow("Inserting new user did not affect one row", zap.Int64("numRowsAffected", numRowsAffected))
		return nil, services.ErrUnableToCreateUser
	}

	s.log.Infow("Created new user", zap.String("id", newUser.ID), zap.String("email", email))
	user, err := s.GetByID(ctx, tx, newUser.ID)
	if err != nil {
		return nil, fmt.Errorf("Unable to get user by ID")
	}

	return user, nil
}
