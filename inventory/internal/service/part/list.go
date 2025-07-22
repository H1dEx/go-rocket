package part

import (
	"context"

	"github.com/H1dEx/go-rocket/inventory/internal/model"
)

func (s *service) ListParts(ctx context.Context, filters model.FilterParts) ([]model.Part, error) {
	parts, err := s.repo.ListParts(ctx, filters)

	if err != nil {
		return []model.Part{}, err
	}

	return parts, nil
}
