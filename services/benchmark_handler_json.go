//+build json

package services

import (
	"context"
	"encoding/json"
)

type FetchJSONDataRequest struct {
	Size int `json:"size"`
}

type FetchJSONDataResponse struct {
	Data map[string]interface{} `json:"data"`
}

func (b *BenchmarkHandler) FetchJSONData(ctx context.Context, req *FetchJSONDataRequest) (*FetchJSONDataResponse, error) {
	data, err := b.dataExamplesFetcher.Fetch(Size(req.Size), JSON)
	if err != nil {
		return nil, err
	}

	resp := &FetchJSONDataResponse{
		Data: map[string]interface{}{},
	}

	err = json.Unmarshal(data, &resp.Data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
