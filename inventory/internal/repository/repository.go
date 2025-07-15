package repository

import (
	"context"
	"github.com/H1dEx/go-rocket/inventory/internal/model"
)

type InventoryRepository interface {
	GetPart(context context.Context, uuid string) (model.Part, error)
	ListParts(context context.Context, filters model.FilterParts) ([]model.Part, error)
}