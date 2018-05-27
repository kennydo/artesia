package controllers

import (
	"net/http"

	"go.uber.org/zap"
)

// MethodNotAllowed handles requests to paths that don't support a certain method
func MethodNotAllowed(log *zap.SugaredLogger, w http.ResponseWriter, r *http.Request) {
	log.Infow(
		"HTTP method not allowed handler",
		zap.String("url", r.URL.String()),
		zap.String("method", r.Method),
	)
	w.WriteHeader(http.StatusMethodNotAllowed)
}
