package main

import (
	"context"
	"flag"
	"time"

	"github.com/rs/zerolog/log"

	pb "github.com/hitanshu-mehta/reaction-timer/api/proto/gameengine/v1"
	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", "localhost:50061", "address to connect")

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to dail connection to gameengine microservice")
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	c := pb.NewGameenginClient(conn)

	if c == nil {
		log.Info().Msg("Client is nil")
	}

	r, err := c.GetSize(timeoutCtx, &pb.GetScoreRequest{})
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to get response")
	}

	if r != nil {
		log.Info().Interface("size", r.GetSize()).Msg("size from gameengine microservice")
	} else {
		log.Error().Msg("couldn't get size")
	}
}
