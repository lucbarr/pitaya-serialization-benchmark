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
	d, err := pitaya.Documentation(true)
	if err != nil {
		return nil, fmt.Errorf("failed to generate documentation for pitaya routes: %w", err)
	}
	doc, err := json.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("failed to encode documentation JSON: %w", err)
	}
	return &protos.Doc{Doc: string(doc)}, nil
}

func (c *DocsHandler) Descriptors(ctx context.Context, names *protos.ProtoNames) (*protos.ProtoDescriptors, error) {
	descriptors := make([][]byte, len(names.Name))
	for i, protoName := range names.Name {
		desc, err := pitaya.Descriptor(protoName)
		if err != nil {
			return nil, fmt.Errorf("failed to get descriptor for '%s': %w", protoName, err)
		}
		descriptors[i] = desc
	}

	return &protos.ProtoDescriptors{Desc: descriptors}, nil
}
