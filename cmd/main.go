package main

import (
	"github.com/rs/zerolog/log"

	"github.com/nelsonalves117/go-orders-api/internal/channels/rest"
	"github.com/nelsonalves117/go-orders-api/internal/config"
)

func main() {
	err := config.Parse()
	if err != nil {
		return
	}

	server := rest.New()

	err = server.Start()
	if err != nil {
		log.Panic().Err(err).Msg("an error occurred while trying to start the server")
	}
}
