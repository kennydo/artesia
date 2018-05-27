package oauth

import (
	"github.com/RangelReale/osin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Service implements the Oauth service, returning Oauth data from the DB
type Service struct {
	log *zap.SugaredLogger
	db  *sqlx.DB
}

// AuthorizeID is the ID of an authorization object
type AuthorizeID string

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

// CreateClient creates a client DB object
func (s *Service) CreateClient(client osin.Client) error {
	_, err := s.db.Exec(
		`INSERT INTO clients (
                        external_id,
                        secret,
                        redirect_uri,
                        user_data) VALUES ($1, $2, $3, $4)`,
		client.GetId(),
		client.GetSecret(),
		client.GetRedirectUri(),
		client.GetUserData(),
	)
	return err
}

// CreateAuthorizationData creates a DB authorization object
func (s *Service) CreateAuthorizationData(data osin.AuthorizeData) error {
	_, err := s.db.Exec(
		`INSERT INTO authorizations (
                     client_id,
                     code,
                     expiration,
                     scope,
                     redirect_uri,
                     external_id,
                     state_data,
                     secret,
                     created_at,
                     user_data
                 ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		data.Client.GetId(),
		data.Code,
		data.ExpiresIn,
		data.Scope,
		data.RedirectUri,
		data.State,
		data.CreatedAt,
		data.UserData,
	)
	return err
}

// CreateAccess creates an access token DB object
func (s *Service) CreateAccess(authID AuthorizeID, data osin.AccessData) error {
	_, err := s.db.Exec(
		`INSERT INTO access_tokens (
                     client_id,
                     authorize_id,
                     token,
                     refresh_token,
                     expiration,
                     scope,
                     redirect_uri,
                     created_at
                 ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		data.Client.GetId(),
		authID,
		data.AccessToken,
		data.RefreshToken,
		data.ExpiresIn,
		data.Scope,
		data.RedirectUri,
		data.CreatedAt,
	)
	return err
}
