package v1

import (
	"context"

	"github.com/H1dEx/go-rocket/inventory/internal/converter"
	inventory_v1 "github.com/H1dEx/go-rocket/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) ListParts(ctx context.Context, in *inventory_v1.ListPartsRequest) (*inventory_v1.ListPartsResponse, error) {
	if in.Filter == nil {
		return nil, status.Error(codes.InvalidArgument, "Filters is invalid")
	}

	filter := converter.ListPartsFilterToModel(in.Filter)
	parts, err := a.service.ListParts(ctx, filter)

	if err != nil {
		return nil, err
	}

	if len(parts) == 0 {
		return nil, status.Error(codes.NotFound, "Not found")
	}

	arr := make([]*inventory_v1.Part, 0, len(parts))
	for _, p := range parts {
		arr = append(arr, converter.PartToProtoModel(p))
	}

	return &inventory_v1.ListPartsResponse{
		Parts: arr,
	}, nil
}
