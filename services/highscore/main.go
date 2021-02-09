package main

import (
	"flag"
	"log"

	grpcSetup "github.com/hitanshu-mehta/reaction-timer/services/highscore/server/"
)

func main() {
	var addressPtr = flag.String("address", ":50051", "address to connect with highscore microservice")
	flag.Parse()

	s := grpcSetup.NewServer(*addressPtr)

	err := s.ListenAndServe(*addressPtr)
	if err != nil {
		log.Fatal().Msg("Failed to start grpc server of highscore microservice")
	}
}
