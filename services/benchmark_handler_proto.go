// +build proto

package services

import (
	"context"
	"pitaya-serialization-benchmark/protos"

	"github.com/golang/protobuf/proto"
)

func (b *BenchmarkHandler) FetchProtoData(ctx context.Context, req *protos.FetchProtoDataRequest) (*protos.FetchProtoDataResponse, error) {
	data, err := b.dataExamplesFetcher.Fetch(Size(req.Size), Protobuf)
	if err != nil {
		return nil, err
	}

	resp := &protos.FetchProtoDataResponse{}

	err = proto.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
