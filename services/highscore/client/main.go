package main

import (
	"context"
	"flag"
	"os"
	"time"

	pb "github.com/hitanshu-mehta/reaction-timer/services/highscore/pb/v1"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

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
		log.Info().Msg("Client is nil")
	}

	r, err := c.GetHighScore(timeoutCtx, &pb.GetHighScoreRequest{})
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to get response")
	}

	if r != nil {
		log.Info().Interface("highscore", r.GetHighscore()).Msg("highscore from highscore microservice")
	} else {
		log.Error().Msg("couldn't get highscore")
	}

}
