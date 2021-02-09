package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	grpcSetup "github.com/hitanshu-mehta/reaction-timer/services/gameengine/server"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	var addressPtr = flag.String("address", ":50061", "address to connect with highscore microservice")
	flag.Parse()

	s := grpcSetup.NewServer(*addressPtr)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal().Msg("failed to start gRPC server for highscore microservice")
	}

}
