package grpc

import "google.golang.org/grpc"

// BffGrpc is grpc server of bff microservice
type BffGrpc struct {
	address string
	srv     *grpc.Server
}
