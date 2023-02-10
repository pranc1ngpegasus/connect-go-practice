package handler

import (
	"context"
	"fmt"

	"github.com/Pranc1ngPegasus/connect-go-practice/domain/logger"
	apiv1 "github.com/Pranc1ngPegasus/connect-go-practice/proto/api/v1"
	connectv1 "github.com/Pranc1ngPegasus/connect-go-practice/proto/api/v1/v1connect"
	connect "github.com/bufbuild/connect-go"
	"github.com/google/wire"
)

var _ connectv1.APIServiceHandler = (*APIV1Handler)(nil)

var NewAPIV1HandlerSet = wire.NewSet(
	wire.Bind(new(connectv1.APIServiceHandler), new(*APIV1Handler)),
	NewAPIV1Handler,
)

type APIV1Handler struct {
	logger logger.Logger
}

func NewAPIV1Handler(
	logger logger.Logger,
) *APIV1Handler {
	return &APIV1Handler{
		logger: logger,
	}
}

func (h *APIV1Handler) Greet(ctx context.Context, req *connect.Request[apiv1.GreetRequest]) (*connect.Response[apiv1.GreetResponse], error) {
	res := connect.NewResponse(&apiv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})

	return res, nil
}
