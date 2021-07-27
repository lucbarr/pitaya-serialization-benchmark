package services

import (
	"context"
	"pitaya-serialization-benchmark/protos"

	"github.com/topfreegames/pitaya/component"
)

type BenchmarkHandler struct {
	component.Base
}

type FetchJSONDataRequest struct {
	Size string `json:"size"`
}

type FetchJSONDataResponse struct {
	Data map[string]interface{} `json:"data"`
}

func (b *BenchmarkHandler) FetchJSONData(ctx context.Context, req *FetchJSONDataRequest) (*FetchJSONDataResponse, error) {
	return &FetchJSONDataResponse{
		Data: map[string]interface{}{
			"potato": "tomato",
		},
	}, nil
}

func (b *BenchmarkHandler) FetchProtoData(ctx context.Context, req *protos.FetchProtoDataRequest) (*protos.FetchProtoDataResponse, error) {
	return &protos.FetchProtoDataResponse{
		AStructMap: map[string]*protos.AStruct{
			"potato": &protos.AStruct{
				AString: "tomato",
			},
		},
	}, nil
}
