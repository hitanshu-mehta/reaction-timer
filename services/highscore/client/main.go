package main

import (
	"context"
	"flag"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {

	var addressPtr = flag.String("address", "localhost:50051", "address to connect")
	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to dail highscore grpc connection")
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	c := pb.NewGameClient(conn)

	if c == nil {
		log.Info("Client is nil")
	}

}
