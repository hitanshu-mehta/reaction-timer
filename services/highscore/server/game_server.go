package grpc

import (
	"context"
	"net"

	pb "github.com/hitanshu-mehta/reaction-timer/api/proto/v1/"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// GameServer server
type GameServer struct {
	address string
	srv     *grpc.Server
}

var highScore = 999999.0

func NewServer(address string) *GameServer {
	return &GameServer{address: address}
}

func (g *GameServer) SetHighScore(ctx context.Context, input *pb.SetHighScoreRequest) (*pb.SetHighScoreResponse, error) {
	log.Info().Msg("SetHighScore in highscore is called")
	highScore = input.HighScore
	return &pb.SetHighScoreResponse{
		Set: true,
	}, nil
}

func (g *GameServer) GetHighScore(ctx context.Context, input *pb.GetHighScoreRequest) (*pb.GetHighScoreResponse, error) {
	log.Info().Msg("GetHighScore in highscore is called")
	return &pb.GetHighScoreResponse{
		HighScore: highScore,
	}, nil
}

func (g *GameServer) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open tcp port")
	}

	serverOpts := []grpc.ServerOption{}

	g.srv = grpc.NewServer(serverOpts...)

	pb.RegisterGameServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("starting gRPC server for m-highscore microservice")

	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "failed to start gRPC server for m-highscore microservice")
	}
	return nil
}
