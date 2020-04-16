package server

import (
	"net/http"
	"time"

	"taxiRequests/taxi_requests"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

//
// Server is a basic application struct
//
type Server struct {
	DB     *taxi_requests.DB
	Router *mux.Router
	Config *Config
}

//
// New retu
//
func New(db *taxi_requests.DB) (*Server, error) {

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
