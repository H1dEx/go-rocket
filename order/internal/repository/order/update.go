package order

import (
	"context"

	"github.com/H1dEx/go-rocket/order/internal/model"
	"github.com/H1dEx/go-rocket/order/internal/repository/converter"
)

func (r *repository) UpdateOrder(ctx context.Context, params model.OrderUpdateParam) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	order, ok := r.orders[params.OrderId]

	if !ok {
		return model.ErrOrderNotFound
	}
	order.TransactionUUID = params.TransactionUUID
	order.Status = converter.OrderStatusToRepoModel(params.Status)
	order.PaymentMethod = converter.OrderPaymentMethodToRepoModel(params.PaymentMethod)
	r.orders[params.OrderId] = order

	return nil
}
