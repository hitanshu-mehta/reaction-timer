package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	grpcSetup "github.com/hitanshu-mehta/reaction-timer/services/highscore/server"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	var addressPtr = flag.String("address", ":50051", "address to connect with highscore microservice")
	flag.Parse()

	s := grpcSetup.NewServer(*addressPtr)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal().Msg("failed to start grpc server of highscore microservice")
	}
}
