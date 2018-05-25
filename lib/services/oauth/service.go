package oauth

import (
	//	"fmt"
	//	"time"

	"github.com/jmoiron/sqlx"
	//	"github.com/kennydo/artesia/lib/services"
	"go.uber.org/zap"
	//	"golang.org/x/crypto/bcrypt"
)

// OauthService implements the Oauth service, returning Oauth data from the DB
type OauthService struct {
	log *zap.SugaredLogger
	db  *sqlx.DB
}

//func (s *OauthService) GetClientByID(id int) ()
