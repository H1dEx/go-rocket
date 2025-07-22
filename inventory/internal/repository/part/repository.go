package part

import (
	"sync"

	def "github.com/H1dEx/go-rocket/inventory/internal/repository"
	"github.com/H1dEx/go-rocket/inventory/internal/repository/model"
)

var _ def.InventoryRepository = (*repository)(nil)

type repository struct {
	mu    sync.RWMutex
	parts map[string]model.Part
}

func NewRepository() *repository {
	r := &repository{
		parts: make(map[string]model.Part),
	}
	r.parts["111"] = model.Part{
		UUID:  "111",
		Name:  "one",
		Price: 100,
	}
	r.parts["222"] = model.Part{
		UUID:  "222",
		Name:  "two",
		Price: 200,
	}
	return r
}
