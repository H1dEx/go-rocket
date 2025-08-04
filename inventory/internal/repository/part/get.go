package part

import (
	"context"
	"errors"
	"log"

	"github.com/H1dEx/go-rocket/inventory/internal/model"
	"github.com/H1dEx/go-rocket/inventory/internal/repository/converter"
	repoModel "github.com/H1dEx/go-rocket/inventory/internal/repository/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) GetPart(ctx context.Context, uuid string) (model.Part, error) {
	var part *repoModel.Part
	err := r.collection.FindOne(ctx, bson.M{"_id": uuid}).Decode(&part)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Part{}, model.ErrPartNotFound
		}
		log.Print("error getting part: %v", err)
		return model.Part{}, err
	}

	return converter.PartToModel(*part), nil
}
