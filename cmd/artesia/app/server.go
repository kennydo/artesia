package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Server struct {
	Log        *zap.SugaredLogger
	httpServer *http.Server
}

func NewServer() (*Server, error) {
	bindAddress := "0.0.0.0:8080"

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugarLogger := logger.Sugar()

	mux := mux.NewRouter()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]string{"Hello", "World"})
	})

	httpServer := &http.Server{
		Handler:      mux,
		Addr:         bindAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server := &Server{
		Log:        sugarLogger,
		httpServer: httpServer,
	}

	return server, nil
}

func (s *Server) Run() error {
	s.Log.Infof("Server listening on: %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
