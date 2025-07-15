package order

import (
	"context"

	"github.com/H1dEx/go-rocket/order/internal/model"
	"github.com/H1dEx/go-rocket/order/internal/repository/converter"
	repoModel "github.com/H1dEx/go-rocket/order/internal/repository/model"
)

func (r *repository) CreateOrder(ctx context.Context, params model.OrderCreateParam) (model.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.orders[params.OrderUUID]
	if ok {
		return model.Order{}, model.ErrOrderAlreadyExists
	}
	order := repoModel.Order{
		OrderUUID:  params.OrderUUID,
		UserUUID:   params.UserUUID,
		PartUuids:  params.PartUuids,
		TotalPrice: params.TotalPrice,
		Status:     repoModel.OrderStatusPendingPayment,
	}
	r.orders[params.OrderUUID] = order

	return converter.OrderToModel(order), nil
}
