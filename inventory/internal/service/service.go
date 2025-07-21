package service

import (
	"context"

	"github.com/H1dEx/go-rocket/inventory/internal/model"
)

type PartService interface {
	GetPart(ctx context.Context, id string) (model.Part, error)
	ListParts(ctx context.Context, filters model.FilterParts) ([]model.Part, error)
}
