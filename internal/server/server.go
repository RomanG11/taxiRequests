package server

import (
	"net/http"
	"taxiRequests/taxiRequests"
	"time"

	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type Config struct {
	HttpPort string `env:"HTTP_PORT,required"`
}

type Server struct {
	DB     *taxiRequests.DB
	Router *mux.Router
	Config *Config
}

//
// ConfigFromEnv func - reads env by struct's fields 'env' annotation
//
func ConfigFromEnv() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}

func New(db *taxiRequests.DB) (*Server, error) {

	config, err := ConfigFromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid server config")
	}

	return &Server{
		Config: config,
		Router: mux.NewRouter(),
		DB:     db,
	}, nil

}

func (s *Server) ConfigureRouter() {
	//routes
	s.Router.HandleFunc("/request", s.requestGet).Methods(http.MethodGet)
	s.Router.HandleFunc("/admin/requests", s.adminRequestGet).Methods(http.MethodGet)

	//middlewares
	s.Router.Use(s.loggingMiddleware)
}

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (s *Server) ListenAndServe() error {
	s.ConfigureRouter()

	log.Info().Msgf("Server: http://[[::]:%s]", s.Config.HttpPort)

	srv := &http.Server{
		Handler:      s.Router,
		Addr:         "0.0.0.0:" + s.Config.HttpPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
