package server

import (
	"context"
	"net/http"
	"time"

	"github.com/Pranc1ngPegasus/connect-go-practice/adapter/server/middleware"
	"github.com/Pranc1ngPegasus/connect-go-practice/domain/configuration"
	"github.com/Pranc1ngPegasus/connect-go-practice/domain/logger"
	connectv1 "github.com/Pranc1ngPegasus/connect-go-practice/proto/api/v1/v1connect"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewServer(
	ctx context.Context,
	logger logger.Logger,
	cfg configuration.Configuration,
	h connectv1.APIServiceHandler,
) *http.Server {
	config := cfg.Server()
	mux := http.NewServeMux()

	reflector := grpcreflect.NewStaticReflector(
		"proto.api.v1.APIService",
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	path, connhandler := connectv1.NewAPIServiceHandler(h)
	mux.Handle(path, connhandler)

	handler := middleware.Chain(mux,
		middleware.Logger(logger),
	)

	logger.Debug(ctx, "listen on", logger.Field("port", config.Port))

	return &http.Server{
		Addr:              ":" + config.Port,
		Handler:           h2c.NewHandler(handler, &http2.Server{}),
		ReadHeaderTimeout: 10 * time.Second,
	}
}
