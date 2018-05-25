package services

import "github.com/RangelReale/osin"

// OauthService exposes CRUD operations on OAuth Clients
// authorizations, and access tokens
type OauthService interface {
	GetClientByID(id int) (osin.Client, error)
	GetAuthorizationByClientID(clientID int) (osin.AuthorizeData, error)
	GetAccessByToken(token string) (osin.AccessData, error)

	StoreClient(client osin.Client) error
	StoreAuthorization(data osin.AuthorizeData) error
	StoreAccess(data osin.AccessData) error
}
