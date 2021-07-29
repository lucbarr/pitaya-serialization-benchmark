//+build json

package services

import "context"

type PingResponse struct {
	Pong int
}

func (t *PingHandler) PingJSON(ctx context.Context) (*PingResponse, error) {
	return &PingResponse{Pong: 1}, nil
}
