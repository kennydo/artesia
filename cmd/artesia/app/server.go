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
	Log        *zap.SugaredLogger
	httpServer *http.Server
	config     *Config
}

// NewServer creates an instance of Server
func NewServer(config *Config) (*Server, error) {
	bindAddress := config.ListenAddress

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugarLogger := logger.Sugar()

	mux := mux.NewRouter()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]string{"Hello", "World"})
	}).Methods("GET")
	// Register routes for our app
	mux.HandleFunc("/api/v1/instance", controllers.GetCurrentInstance).Methods("GET")
	mux.HandleFunc("/api/v1/apps", controllers.RegisterOAuthApplication).Methods("POST")

	// Register error handlers
	mux.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sugarLogger.Infow(
			"No HTTP handler found",
			zap.String("url", r.URL.String()),
			zap.String("method", r.Method),
		)
		w.WriteHeader(http.StatusNotFound)
	})

	mux.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sugarLogger.Infow(
			"HTTP method not allowed handler",
			zap.String("url", r.URL.String()),
			zap.String("method", r.Method),
		)
		w.WriteHeader(http.StatusMethodNotAllowed)
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
		config:     config,
	}

	return server, nil
}

// Run runs the HTTP Server
func (s *Server) Run() error {
	s.Log.Infof("Server listening on: http://%s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
