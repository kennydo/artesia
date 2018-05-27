package controllers

import (
	"encoding/json"
	"net/http"
)

type OauthAuthorize struct{}

func (o *OauthAuthorize) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]string{"Hello world!"})
}
