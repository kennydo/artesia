package controllers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

// OAuthCredentials holds the data required for OAuth apps to authenticate to us
type OAuthCredentials struct {
	ID           string `json:"id"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// RegisterOAuthApplication creates new OAuth credentials for users
func RegisterOAuthApplication(log *zap.SugaredLogger, w http.ResponseWriter, r *http.Request) {
	credentials := OAuthCredentials{
		ID:           "imouto",
		ClientID:     "imoutoID",
		ClientSecret: "imoutoHimitsu",
	}
	json.NewEncoder(w).Encode(credentials)
}
