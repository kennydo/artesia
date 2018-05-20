package user

import (
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
	db  *sqlx.DB
}

// GetByID gets a user by ID. Returns an error if no user is found.
func (s *DBService) GetByID(id int) (*services.User, error) {
	dbUser := DBUser{}
	err := s.db.Get(&dbUser, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("No user found")
	}

	user := services.User(dbUser)
	return &user, nil
}

// GetByEmail gets a user by their (case-insensitive) email. Returns an error if no user is found.
func (s *DBService) GetByEmail(email string) (*services.User, error) {
	dbUser := DBUser{}
	err := s.db.Get(&dbUser, "SELECT * FROM users WHERE lower(email) = lower(?)", email)
	if err != nil {
		return nil, fmt.Errorf("No user found")
	}

	user := services.User(dbUser)
	return &user, nil
}

// CreateUser creates a user. Returns an error if unable to create user.
func (s *DBService) CreateUser(email string, plaintextPassword string) (*services.User, error) {
	// We don't want to create a user if there's already one with the same email
	existingUser, err := s.GetByEmail(email)
	if existingUser != nil || err == nil {
		return nil, fmt.Errorf("Email already taken")
	}

	passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)
	if err != nil {
		s.log.Info("Unable to hash password", zap.Error(err), zap.String("email", email))
		return nil, fmt.Errorf("Invalid password")
	}

	result, err := s.db.NamedExec(
		`INSERT INTO users (
			email,
			password_hash,
			created_at
		) VALUES (
			:email,
			:password_hash,
			:created_at
		)`,
		&DBUser{
			Email:        email,
			PasswordHash: string(passwordHashBytes),
			CreatedAt:    time.Now().In(time.UTC),
		})
	if err != nil {
		s.log.Info("Unable to insert into users table", zap.Error(err))
		return nil, fmt.Errorf("Unable to insert user")
	}
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		s.log.Info("Unable to get inserted user ID", zap.Error(err), zap.String("email", email))
		return nil, fmt.Errorf("Unable to get inserted user ID")
	}

	s.log.Info("Created new user", zap.Int64("id", lastInsertedID), zap.String("email", email))
	user, err := s.GetByID(int(lastInsertedID))
	if err != nil {
		return nil, fmt.Errorf("Unable to get user by ID")
	}

	return user, nil
}
