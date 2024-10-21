// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-parrot/parrot-starter/internal/server"
	"github.com/go-parrot/parrot-starter/internal/tasks"
	"github.com/go-parrot/parrot/pkg/app"
	"github.com/go-parrot/parrot/pkg/log"
	"github.com/go-parrot/parrot/pkg/transport/consumer/redis"
	"github.com/go-parrot/parrot/pkg/transport/http"
)

// Injectors from wire.go:

func InitApp(cfg *app.Config, config *app.ServerConfig, tc *tasks.Config) (*app.App, func(), error) {
	httpServer := server.NewHTTPServer(config)
	redisServer := server.NewRedisConsumerServer(tc)
	appApp := newApp(cfg, httpServer, redisServer)
	return appApp, func() {
	}, nil
}

// wire.go:

// 第三个参数需要根据使用的server 进行调整
// 默认使用 redis, 如果使用 rabbitmq 可以改为: rs *rabbitmq.Server
// 然后执行 wire
func newApp(cfg *app.Config, hs *http.Server, rs *redis.Server) *app.App {
	log.Init(log.WithFilename("consumer"))

	return app.New(app.WithName(cfg.Name), app.WithVersion(cfg.Version), app.WithLogger(log.GetLogger()), app.WithServer(

		hs,

		rs,
	),
	)
}
