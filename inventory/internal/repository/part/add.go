package part

import (
	"context"
	"log"

	"github.com/H1dEx/go-rocket/inventory/internal/model"
	"github.com/H1dEx/go-rocket/inventory/internal/repository/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repository) AddPart(ctx context.Context, part model.Part) (primitive.ObjectID, error) {
	convPart := converter.PartToRepoModel(part)
	res, err := r.collection.InsertOne(ctx, convPart)
	if err != nil {
		return primitive.NilObjectID, err
	}

	log.Printf("gotten: %+v", res.InsertedID)

	return primitive.NilObjectID, nil
}
