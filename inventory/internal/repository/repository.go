package repository

import (
	"context"

	"github.com/H1dEx/go-rocket/inventory/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryRepository interface {
	GetPart(context context.Context, uuid string) (model.Part, error)
	ListParts(context context.Context, filters model.FilterParts) ([]model.Part, error)
	AddPart(ctx context.Context, part model.Part) (primitive.ObjectID, error)
}