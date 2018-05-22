package controllers

import (
	"encoding/json"
	"net/http"
)

// URLs has data about related URLs for this instance
type URLs struct {
	StreamingAPI string `json:"streaming_api"`
}

// Instance holds data about this current instance
type Instance struct {
	URI         string   `json:"uri"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Email       string   `json:"email"`
	Version     string   `json:"version"`
	URLs        URLs     `json:"urls"`
	Languages   []string `json:"languages"`
}

// GetCurrentInstance returns information about this instance
func GetCurrentInstance(w http.ResponseWriter, r *http.Request) {
	currentInstance := Instance{
		URI:         r.Host,
		Title:       "artesia",
		Description: "Lorem ipsum imouto",
		Email:       "artesia@example.com",
		Version:     "2.4.0 (compatible; Artesia dev)",
		URLs: URLs{
			StreamingAPI: "",
		},
		Languages: []string{"en", "jp"},
	}
	json.NewEncoder(w).Encode(currentInstance)
}
