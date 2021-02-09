package bff

import (
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	pbgameengine "github.com/hitanshu-mehta/reaction-timer/api/proto/gameengine/v1"
	pbhighscore "github.com/hitanshu-mehta/reaction-timer/api/proto/highscore/v1"
	"google.golang.org/grpc"
)

// GameResource contains client of highscore and gameengine microservice
type GameResource struct {
	gameClient       pbhighscore.GameClient
	gameEngineClient pbgameengine.GameengineClient
}

// NewGameResource returns new instance of gameResource
func NewGameResource(gameClient pbhighscore.GameClient, gameEngineClient pbgameengine.GameengineClient) *GameResource {
	return &GameResource{
		gameClient:       gameClient,
		gameEngineClient: gameEngineClient,
	}
}

// NewGrpcGameServiceClient dials up grpc server of highscore microservice and returns the client
func NewGrpcGameServiceClient(serverAddr string) (pbhighscore.GameClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msgf("failed to dial: %v", serverAddr)
		return nil, err
	}
	log.Info().Msgf("successfully connected to server at: %v", serverAddr)

	if conn == nil {
		log.Info().Msg("highscore microservice connection is nil in bff")
	}

	c := pbhighscore.NewGameClient(conn)

	return c, nil
}

// NewGrpcGameEngineServiceClient dials up grpc server of gameengine microservice and returns the client
func NewGrpcGameEngineServiceClient(serverAddr string) (pbgameengine.GameengineClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msgf("failed to dial: %v", serverAddr)
	}
	log.Info().Msgf("succesfully connected to server at: %v", serverAddr)

	if conn == nil {
		log.Info().Msg("gameengine microservice connection is nil in bff")
	}

	c := pbgameengine.NewGameengineClient(conn)

	return c, nil
}

// SetHighScore listens to http request and call SetHighScore of highScore microservice
func (gr *GameResource) SetHighScore(c *gin.Context) {
	highScoreString := c.Param("hs")
	highScoreFloat64, err := strconv.ParseFloat(highScoreString, 64)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert highscore to float64")
	}

	set, err := gr.gameClient.SetHighScore(c, &pbhighscore.SetHighScoreRequest{
		Highscore: highScoreFloat64,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to set highScore")
	}
	c.JSONP(200, gin.H{
		"set": set,
	})
}

// GetHighScore listens to http request and call GetHighScore of highScore microservice
func (gr *GameResource) GetHighScore(c *gin.Context) {
	highScoreResponse, err := gr.gameClient.GetHighScore(c, &pbhighscore.GetHighScoreRequest{})
	if err != nil {
		log.Error().Err(err).Msg("failed to get highscore")
	}

	hsstring := strconv.FormatFloat(highScoreResponse.Highscore, 'e', -1, 64)

	c.JSONP(200, gin.H{
		"hs": hsstring,
	})

}

// SetScore listens to http request and call SetScore of gameengine microservice
func (gr *GameResource) SetScore(c *gin.Context) {
	scoreString := c.Param("score")
	scoreFloat64, err := strconv.ParseFloat(scoreString, 64)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert score to float64")
	}

	set, err := gr.gameEngineClient.SetScore(c, &pbgameengine.SetScoreRequest{Score: scoreFloat64})
	if err != nil {
		// TODO: Return server failure response
		log.Error().Err(err).Msg("failed to set score")
	}

	c.JSONP(200, gin.H{
		"set": set,
	})
}

// GetSize listens to http request and call GetSize of gameengine microservice
func (gr *GameResource) GetSize(c *gin.Context) {
	size, err := gr.gameEngineClient.GetSize(c, &pbgameengine.GetSizeRequest{})
	if err != nil {
		log.Error().Err(err).Msg("failed to get size")
	}

	c.JSON(200, gin.H{
		"size": size,
	})
}
