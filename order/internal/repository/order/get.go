package order

import (
	"context"

	"github.com/H1dEx/go-rocket/order/internal/model"
	"github.com/H1dEx/go-rocket/order/internal/repository/converter"
)

func (r *repository) GetOrderByID (ctx context.Context, id string) (model.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	order, ok := r.orders[id]

	if !ok {
		return model.Order{}, model.ErrOrderNotFound
	}

	return converter.OrderToModel(order), nil
}