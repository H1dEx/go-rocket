package order

import (
	"context"

	"github.com/H1dEx/go-rocket/order/internal/model"
)

func (s *service) GetOrderByID(ctx context.Context, id string) (model.Order, error) {
	order, err := s.repo.GetOrderByID(ctx, id)

	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}
