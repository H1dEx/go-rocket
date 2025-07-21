package v1

import (
	"context"

	"github.com/H1dEx/go-rocket/order/internal/client/converter"
	"github.com/H1dEx/go-rocket/order/internal/model"
	inventory_v1 "github.com/H1dEx/go-rocket/shared/pkg/proto/inventory/v1"
)

func (c *client) ListParts(ctx context.Context, filter model.FilterParts) ([]model.Part, error) {
	parts, err := c.genClient.ListParts(ctx, &inventory_v1.ListPartsRequest{Filter: converter.PartsFilterToProto(filter)})

	if err != nil {
		return nil, err
	}

	return converter.PartListToModel(parts.Parts), nil
}
