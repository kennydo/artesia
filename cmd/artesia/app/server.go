package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kennydo/artesia/cmd/artesia/app/controllers"
	"go.uber.org/zap"
)

// Server contains information needed to start the HTTP Server
type Server struct {
	log        *zap.SugaredLogger
	httpServer *http.Server
	config     *Config
}

type artesiaHandlerFunc func(*zap.SugaredLogger, http.ResponseWriter, *http.Request)

// NewServer creates an instance of Server
func NewServer(config *Config) (*Server, error) {
	bindAddress := config.ListenAddress

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugarLogger := logger.Sugar()

	mux := mux.NewRouter()

	httpServer := &http.Server{
		Handler:      mux,
		Addr:         bindAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server := &Server{
		log:        sugarLogger,
		httpServer: httpServer,
		config:     config,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]string{"Hello", "World"})
	}).Methods("GET")

	// Register routes for our app
	mux.HandleFunc("/api/v1/instance", server.addMiddleware(controllers.GetCurrentInstance)).Methods("GET")
	mux.HandleFunc("/api/v1/apps", server.addMiddleware(controllers.RegisterOAuthApplication)).Methods("POST")

	// Register error handlers
	mux.NotFoundHandler = server.addMiddleware(controllers.NotFound)
	mux.MethodNotAllowedHandler = server.addMiddleware(controllers.MethodNotAllowed)

	return server, nil
}

func (s *Server) addMiddleware(h artesiaHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(s.log, w, r)
	}
}

// Run runs the HTTP Server
func (s *Server) Run() error {
	s.log.Infof("Server listening on: https://%s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
