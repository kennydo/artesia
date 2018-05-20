package user

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type DBServiceTestSuite struct {
	suite.Suite
	service *DBService
}

func (suite *DBServiceTestSuite) SetupTest() {
	db, err := sqlx.Connect("postgres", "host=127.0.0.1 user=artesia dbname=artesia password=saylamass sslmode=disable")
	suite.Assert().Nil(err)

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugarLogger := logger.Sugar()

	suite.service = &DBService{
		db:  db,
		log: sugarLogger,
	}
}

func (suite *DBServiceTestSuite) TearDownTest() {
	suite.service.db.Exec("TRUNCATE TABLE users")
}

func (suite *DBServiceTestSuite) TestCreatingUser() {
	email := "artesia@example.com"
	plaintextPassword := "合言葉"

	createdUser, err := suite.service.CreateUser(email, plaintextPassword)
	suite.Assert().Nil(err)

	suite.Assert().Equal(createdUser.Email, email)

	err = bcrypt.CompareHashAndPassword([]byte(createdUser.PasswordHash), []byte(plaintextPassword))
	suite.Assert().Nil(err)

	userGottenByEmail, err := suite.service.GetByEmail(email)
	suite.Assert().Nil(err)

	suite.Assert().Equal(userGottenByEmail, createdUser)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(DBServiceTestSuite))
}
