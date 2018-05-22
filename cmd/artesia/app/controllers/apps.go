package controllers

import (
	"encoding/json"
	"net/http"
)

// OAuthCredentials holds the data required for OAuth apps to authenticate to us
type OAuthCredentials struct {
	ID           string `json:"id"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// RegisterOAuthApplication creates new OAuth credentials for users
func RegisterOAuthApplication(w http.ResponseWriter, r *http.Request) {
	credentials := OAuthCredentials{
		ID:           "imouto",
		ClientID:     "imoutoID",
		ClientSecret: "imoutoHimitsu",
	}
	json.NewEncoder(w).Encode(credentials)
}
