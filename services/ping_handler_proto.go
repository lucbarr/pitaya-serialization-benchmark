//+build proto

package services

import (
	"context"
	"pitaya-serialization-benchmark/protos"
)

func (t *PingHandler) PingProto(ctx context.Context) (*protos.PingResponse, error) {
	return &protos.PingResponse{
		Pong: 1,
	}, nil
}
