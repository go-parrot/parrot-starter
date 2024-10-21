package server

import (
	"github.com/go-parrot/parrot/pkg/app"
	"github.com/go-parrot/parrot/pkg/transport/grpc"
)

// NewGRPCServer creates a gRPC server
func NewGRPCServer(cfg *app.Config) *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.Network("tcp"),
		grpc.Address(cfg.GRPC.Addr),
		grpc.Timeout(cfg.GRPC.ReadTimeout),
	)

	// register biz service
	// v1.RegisterUserServiceServer(grpcServer, service.Svc.Users())

	return grpcServer
}
