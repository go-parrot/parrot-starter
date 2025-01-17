// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-parrot/parrot-starter/internal/server"
	"github.com/go-parrot/parrot/pkg/app"
	"github.com/go-parrot/parrot/pkg/log"
	"github.com/go-parrot/parrot/pkg/transport/grpc"
	"github.com/go-parrot/parrot/pkg/transport/http"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func InitApp(cfg *app.Config) (*app.App, func(), error) {
	httpServer := server.NewHTTPServer(cfg)
	grpcServer := server.NewGRPCServer(cfg)
	appApp := newApp(cfg, httpServer, grpcServer)
	return appApp, func() {
	}, nil
}

// wire.go:

func newApp(cfg *app.Config, hs *http.Server, gs *grpc.Server) *app.App {
	log.Init(log.WithFilename("app"))

	return app.New(app.WithName(cfg.Name), app.WithVersion(cfg.Version), app.WithLogger(log.GetLogger()), app.WithServer(

		hs,

		gs,
	),
	)
}
