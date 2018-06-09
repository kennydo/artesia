package user

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/kennydo/artesia/lib/services"

	"github.com/jmoiron/sqlx"
	"github.com/kennydo/artesia/cmd/artesia/app"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type DBServiceTestSuite struct {
	suite.Suite
	service *DBService
	ctx     context.Context
	tx      *sqlx.Tx
}

func (suite *DBServiceTestSuite) SetupTest() {
	config, err := app.LoadConfig()
	if err != nil {
		// Errs during setup aren't handled by suite's helper methods, so we manually call Fatal
		log.Fatalf("Failed to load config: %v", err)
	}

	dbConnectInfo := fmt.Sprintf(
		`host=%s dbname=%s user=%s password=%s sslmode=disable`,
		config.Database.Host,
		config.Database.DBName,
		config.Database.User,
		config.Database.Password,
	)

	db, err := sqlx.Connect("postgres", dbConnectInfo)
	if err != nil {
		log.Fatalf("Failed to connect to the DB: %v", err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugarLogger := logger.Sugar()

	suite.service = &DBService{
		log: sugarLogger,
	}
	suite.ctx = context.TODO()
	suite.tx = db.MustBegin()
}

func (suite *DBServiceTestSuite) TearDownTest() {
}

func (suite *DBServiceTestSuite) TestCreatingUser() {
	email := "artesia@example.com"
	plaintextPassword := "合言葉"

	createdUser, err := suite.service.CreateUser(suite.ctx, suite.tx, email, plaintextPassword)
	suite.Assert().Nil(err)

	suite.Assert().Equal(createdUser.Email, email)

	err = bcrypt.CompareHashAndPassword([]byte(createdUser.PasswordHash), []byte(plaintextPassword))
	suite.Assert().Nil(err)

	userGottenByEmail, err := suite.service.GetByEmail(suite.ctx, suite.tx, email)
	suite.Assert().Nil(err)

	suite.Assert().Equal(userGottenByEmail, createdUser)
}

func (suite *DBServiceTestSuite) TestDuplicateCreateUserErrors() {
	email := "artesia@example.com"
	plaintextPassword := "合言葉"

	_, err := suite.service.CreateUser(suite.ctx, suite.tx, email, plaintextPassword)
	suite.Assert().Nil(err)

	secondAttempt, err := suite.service.CreateUser(suite.ctx, suite.tx, email, plaintextPassword)
	suite.Assert().Nil(secondAttempt)
	suite.Assert().Equal(err, services.ErrUserEmailTaken)
}

func (suite *DBServiceTestSuite) TestCreatingUserWithSameCanonicalEmailErrors() {
	email := "artesia@example.com"
	plaintextPassword := "合言葉"

	_, err := suite.service.CreateUser(suite.ctx, suite.tx, email, plaintextPassword)
	suite.Assert().Nil(err)

	similarEmail := "ArTeSiA@example.com"
	_, err = suite.service.CreateUser(suite.ctx, suite.tx, similarEmail, plaintextPassword)
	suite.Assert().Equal(err, services.ErrUserEmailTaken)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(DBServiceTestSuite))
}
