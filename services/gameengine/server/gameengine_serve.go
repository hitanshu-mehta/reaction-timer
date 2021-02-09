package grpc

import (
	"context"
	"net"

	logic "github.com/hitanshu-mehta/reaction-timer/api/proto/gameengine/logic"
	pb "github.com/hitanshu-mehta/reaction-timer/api/proto/gameengine/v1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// GameenginServer creates new instance of gRPC server
type GameenginServer struct {
	pb.UnimplementedGameengineServer
	srv     *grpc.Server
	address string
}

// NewServer creates new instance of the server
func NewServer(address string) *GameenginServer {
	return &GameenginServer{
		address: address,
	}
}

// SetScore saves the last score
func (g *GameenginServer) SetScore(ctx context.Context, input *pb.SetScoreRequest) (*pb.SetScoreResponse, error) {
	log.Info().Msg("SetScore in gameengine called")
	set := logic.SetScore(input.Score)
	return &pb.SetScoreResponse{
		Set: set,
	}, nil
}

// GetSize returns the size according to performance of the user
func (g *GameenginServer) GetSize(ctx context.Context, input *pb.GetSizeRequest) (*pb.GetSizeResponse, error) {
	log.Info().Msg("GetSize in gameengine called")
	size := logic.GetSize()
	return &pb.GetSizeResponse{
		Size: size,
	}, nil
}

// ListenAndServe starts to listen tcp port and start the server
func (g *GameenginServer) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open tcp port")
	}

	serverOpts := []grpc.ServerOption{}

	g.srv = grpc.NewServer(serverOpts...)

	pb.RegisterGameengineServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("starting gRPC server for gameengine microservice")

	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "failed to start gRPC server for gameengine microservice")
	}

	return nil
}
