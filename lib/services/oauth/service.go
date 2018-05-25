package oauth

import (
	//	"fmt"
	//	"time"

	"github.com/jmoiron/sqlx"
	//	"github.com/kennydo/artesia/lib/services"
	"github.com/RangelReale/osin"
	"go.uber.org/zap"
	//	"golang.org/x/crypto/bcrypt"
)

// Service implements the Oauth service, returning Oauth data from the DB
type Service struct {
	log *zap.SugaredLogger
	db  *sqlx.DB
}

// GetClientByID returns an osin.Client by its ID
func (s *Service) GetClientByID(id int) (osin.Client, error) {
	dbClient := Client{}
	err := s.db.Get(&dbClient, "SELECT * FROM clients WHERE id = $1", id)
	if err != nil {
		return &osin.DefaultClient{}, err
	}
	return &osin.DefaultClient{
		Id:          dbClient.ExternalID,
		Secret:      dbClient.Secret,
		RedirectUri: dbClient.RedirectURI,
		UserData:    dbClient.UserData,
	}, nil
}

// GetAuthorizationClientById returns osin.AuthorizationData for a client
// TODO: make canned erors
func (s *Service) GetAuthorizationClientByID(id int) (osin.AuthorizeData, error) {
	dbAuthorization := Authorization{}
	err := s.db.Get(&dbAuthorization, "SELECT * FROM authorizations WHERE client_id = $1", id)
	if err != nil {
		return osin.AuthorizeData{}, err
	}

	client, err := s.GetClientByID(dbAuthorization.ClientID)
	if err != nil {
		return osin.AuthorizeData{}, err
	}

	return osin.AuthorizeData{
		Client:      client,
		Code:        dbAuthorization.Code,
		ExpiresIn:   dbAuthorization.Expiration,
		Scope:       dbAuthorization.Scope,
		RedirectUri: dbAuthorization.RedirectURI,
		UserData:    dbAuthorization.UserData,
	}, nil
}
