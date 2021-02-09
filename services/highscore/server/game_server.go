package grpc

import (
	"context"
	"net"

	pb "github.com/hitanshu-mehta/reaction-timer/api/proto/v1/highscore"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// GameServer server
type GameServer struct {
	pb.UnimplementedGameServer
	address string
	srv     *grpc.Server
}

var highScore = 999999.0

// NewServer returns server
func NewServer(address string) *GameServer {
	return &GameServer{address: address}
}

// SetHighScore sets high score
func (g *GameServer) SetHighScore(ctx context.Context, input *pb.SetHighScoreRequest) (*pb.SetHighScoreResponse, error) {
	log.Info().Msg("SetHighScore in highscore is called")
	highScore = input.Highscore
	return &pb.SetHighScoreResponse{
		Set: true,
	}, nil
}

// GetHighScore returns current highscore
func (g *GameServer) GetHighScore(ctx context.Context, input *pb.GetHighScoreRequest) (*pb.GetHighScoreResponse, error) {
	log.Info().Msg("GetHighScore in highscore is called")
	return &pb.GetHighScoreResponse{
		Highscore: highScore,
	}, nil
}

// ListenAndServe start highscore microservice server
func (g *GameServer) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrapf(err, "failed to open tcp port")
	}

	serverOpts := []grpc.ServerOption{}

	g.srv = grpc.NewServer(serverOpts...)

	pb.RegisterGameServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("starting gRPC server for highscore microservice")

	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrapf(err, "failed to start gRPC server for highscore microservice")
	}

	return nil
}
