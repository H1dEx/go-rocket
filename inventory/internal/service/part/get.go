package part

import (
	"context"

	"github.com/H1dEx/go-rocket/inventory/internal/model"
)

func (s *service) GetPart(ctx context.Context, id string) (model.Part, error) {
	part, err := s.repo.GetPart(ctx, id)

	if err != nil {
		return model.Part{}, err
	}

	return part, nil
}
