//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"net/http"

	"github.com/Pranc1ngPegasus/connect-go-practice/adapter/handler"
	"github.com/Pranc1ngPegasus/connect-go-practice/adapter/server"
	"github.com/Pranc1ngPegasus/connect-go-practice/infra/configuration"
	"github.com/Pranc1ngPegasus/connect-go-practice/infra/logger"
	"github.com/google/wire"
)

type app struct {
	server *http.Server
}

func initialize() (*app, error) {
	wire.Build(
		context.Background,
		logger.NewLoggerSet,
		configuration.NewConfigurationSet,
		handler.NewAPIV1HandlerSet,
		server.NewServer,

		wire.Struct(new(app), "*"),
	)

	return nil, nil
}
