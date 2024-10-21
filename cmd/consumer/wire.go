//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-parrot/parrot-starter/internal/server"

	//"github.com/go-parrot/parrot-starter/internal/service"
	"github.com/go-parrot/parrot-starter/internal/tasks"
	parrot "github.com/go-parrot/parrot/pkg/app"
	logger "github.com/go-parrot/parrot/pkg/log"
	redisMQ "github.com/go-parrot/parrot/pkg/transport/consumer/redis"
	httpSrv "github.com/go-parrot/parrot/pkg/transport/http"
	"github.com/google/wire"
)

func InitApp(cfg *parrot.Config, config *parrot.ServerConfig, tc *tasks.Config) (*parrot.App, func(), error) {
	wire.Build(server.ProviderSetForConsumer, newApp)
	return &parrot.App{}, nil, nil
}

// 第三个参数需要根据使用的server 进行调整
// 默认使用 redis, 如果使用 rabbitmq 可以改为: rs *rabbitmq.Server
// 然后执行 wire
func newApp(cfg *parrot.Config, hs *httpSrv.Server, rs *redisMQ.Server) *parrot.App {
	logger.Init(logger.WithFilename("consumer"))

	return parrot.New(
		parrot.WithName(cfg.Name),
		parrot.WithVersion(cfg.Version),
		parrot.WithLogger(logger.GetLogger()),
		parrot.WithServer(
			// init HTTP server
			hs,
			// init consumer server
			rs,
		),
	)
}
