package order

import (
	"context"

	"github.com/H1dEx/go-rocket/order/internal/model"
)

func (s *service) OrderCancelById(ctx context.Context, id string) error {
	updData := model.OrderUpdateParam{
		OrderId: id,
		Status: model.OrderStatusCancelled,
	}
	return s.repo.UpdateOrder(ctx, updData)
}