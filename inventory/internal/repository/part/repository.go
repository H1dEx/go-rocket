package part

import (
	def "github.com/H1dEx/go-rocket/inventory/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ def.InventoryRepository = (*repository)(nil)

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *repository {
	collection := db.Collection("parts")

	r := &repository{
		collection,
	}
	// r.parts["111"] = model.Part{
	// 	UUID:  "111",
	// 	Name:  "one",
	// 	Price: 100,
	// }
	// r.parts["222"] = model.Part{
	// 	UUID:  "222",
	// 	Name:  "two",
	// 	Price: 200,
	// }
	return r
}
