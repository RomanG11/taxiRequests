package main

import (
	"github.com/rs/zerolog/log"
	"taxiRequests/internal/server"
	"taxiRequests/taxiRequests"
	"time"
)

func main() {

	db := taxiRequests.InitDB()

	go func(db *taxiRequests.DB) {
		for {
			time.Sleep(200 * time.Millisecond)
			db.Roll()
		}
	}(db)

	srv, err := server.New(db)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize rest server")
	}

	log.Info().Msg("Starting server")
	log.Fatal().Err(srv.ListenAndServe()).Msg("Server stopped with error")
}
