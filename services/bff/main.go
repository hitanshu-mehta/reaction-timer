package main

import (
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	bff "github.com/hitanshu-mehta/reaction-timer/services/bff/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	grpcAddressHighScore := flag.String("address-highscore", "localhost:50051", "gRPC address of highscore microservice")
	grpcAddressGameEngine := flag.String("address-gameengine", "localhost:50061", "gRPC address of gameengine microservice")

	serverAddress := flag.String("address-http", ":8081", "Http server address")

	flag.Parse()

	gameClient, err := bff.NewGrpcGameServiceClient(*grpcAddressHighScore)
	if err != nil {
		log.Error().Err(err).Msg("failed to start highscore service client")
	}

	gameEngineClient, err := bff.NewGrpcGameEngineServiceClient(*grpcAddressGameEngine)
	if err != nil {
		log.Error().Err(err).Msg("failed to start gameengine service client")
	}

	gr := bff.NewGameResource(gameClient, gameEngineClient)

	router := gin.Default()
	router.GET("/highscore", gr.GetHighScore)
	router.GET("/highscore/:hs", gr.SetHighScore)
	router.GET("/score/:score", gr.SetScore)
	router.GET("/size", gr.GetSize)

	err = router.Run(*serverAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start bff")
	}

	log.Info().Msgf("started http server at: %v", serverAddress)
}
