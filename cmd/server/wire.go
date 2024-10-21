//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-parrot/parrot-starter/internal/server"
	parrot "github.com/go-parrot/parrot/pkg/app"
	logger "github.com/go-parrot/parrot/pkg/log"
	"github.com/go-parrot/parrot/pkg/transport/grpc"
	httpSrv "github.com/go-parrot/parrot/pkg/transport/http"
	"github.com/google/wire"
)

func InitApp(cfg *parrot.Config) (*parrot.App, func(), error) {
	// wire.Build(server.ProviderSet, service.ProviderSet, repository.ProviderSet, cache.ProviderSet, newApp)
	wire.Build(server.ProviderSet, newApp)
	return &parrot.App{}, nil, nil
}

func newApp(cfg *parrot.Config, hs *httpSrv.Server, gs *grpc.Server) *parrot.App {
	logger.Init(logger.WithFilename("app"))

	return parrot.New(
		parrot.WithName(cfg.Name),
		parrot.WithVersion(cfg.Version),
		parrot.WithLogger(logger.GetLogger()),
		parrot.WithServer(
			// init HTTP server
			hs,
			// init GRPC server
			gs,
		),
	)
}
