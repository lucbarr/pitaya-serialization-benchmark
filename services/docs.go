package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/protos"
)

type DocsHandler struct {
	component.Base
}

func (c *DocsHandler) Docs(ctx context.Context) (*protos.Doc, error) {
	docsMap, err := pitaya.Documentation(true)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(docsMap)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(data))

	return &protos.Doc{
		Doc: string(data),
	}, nil
}
