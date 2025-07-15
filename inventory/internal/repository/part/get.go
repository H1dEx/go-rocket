package part

import (
	"context"

	"github.com/H1dEx/go-rocket/inventory/internal/model"
	"github.com/H1dEx/go-rocket/inventory/internal/repository/converter"
)

func (r *repository) GetPart(_ context.Context, uuid string) (model.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	part, ok := r.parts[uuid]

	if !ok {
		return model.Part{}, model.ErrPartNotFound
	}

	return converter.PartToModel(part), nil
}
