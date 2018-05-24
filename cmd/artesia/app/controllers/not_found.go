package controllers

import (
	"net/http"

	"go.uber.org/zap"
)

// NotFound handles requests to paths that don't have a handle registered
func NotFound(log *zap.SugaredLogger, w http.ResponseWriter, r *http.Request) {
	log.Infow(
		"No HTTP handler found",
		zap.String("url", r.URL.String()),
		zap.String("method", r.Method),
	)
	w.WriteHeader(http.StatusNotFound)
}
