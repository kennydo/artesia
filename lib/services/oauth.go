package services

import (
	"github.com/RangelReale/osin"
	"github.com/kennydo/artesia/lib/services/oauth"
)

// OauthService exposes CRUD operations on OAuth Clients
// authorizations, and access tokens
type OauthService interface {
	GetClientByID(id int) (osin.Client, error)
	GetAuthorizationByClientID(clientID int) (osin.AuthorizeData, error)
	GetAccessByToken(token string) (osin.AccessData, error)

	CreateClient(client osin.Client) error
	CreateAuthorization(authID oauth.AuthorizeID, data osin.AuthorizeData) error
	CreateAccess(data osin.AccessData) error
}
