package user

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kennydo/artesia/lib/services"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// DBService implements the user service, backed by the DB
type DBService struct {
	log *zap.SugaredLogger
}

// GetByID gets a user by ID. Returns an error if no user is found.
func (s *DBService) GetByID(ctx context.Context, tx *sqlx.Tx, id int) (*services.User, error) {
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
			email,
			password_hash,
			created_at
		) VALUES (
			:email,
			:password_hash,
			:created_at
		)
		RETURNING id
		`)
	if err != nil {
		s.log.Infow("Unable to prepare insert statement", zap.Error(err))
		return nil, fmt.Errorf("Unable to prepare user for insert")
	}

	var insertedID int
	err = stmt.GetContext(ctx, &insertedID, &DBUser{
		Email:        email,
		PasswordHash: string(passwordHashBytes),
		CreatedAt:    time.Now().In(time.UTC),
	})
	if err != nil {
		s.log.Infow("Unable to insert user into DB", zap.Error(err), zap.String("email", email))
		return nil, fmt.Errorf("Unable to insert user into DB")
	}

	s.log.Infow("Created new user", zap.Int("id", insertedID), zap.String("email", email))
	user, err := s.GetByID(ctx, tx, insertedID)
	if err != nil {
		return nil, fmt.Errorf("Unable to get user by ID")
	}

	return user, nil
}
