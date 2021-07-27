package services

import (
	"context"
	"pitaya-serialization-benchmark/protos"

	"github.com/topfreegames/pitaya/component"
)

type PingHandler struct {
	component.Base
}

type PingResponse struct {
	Message string
}

func (t *PingHandler) PingJSON(ctx context.Context) (*PingResponse, error) {
	return &PingResponse{Message: "pong"}, nil
}

func (t *PingHandler) PingProto(ctx context.Context) (*protos.PingResponse, error) {
	return &protos.PingResponse{
		Pong: 1,
	}, nil
}
