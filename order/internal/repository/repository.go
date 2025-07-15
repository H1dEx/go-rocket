package repository

import (
	"context"
	"github.com/H1dEx/go-rocket/order/internal/model"
)

type OrderRepository interface {
	CreateOrder (ctx context.Context, params model.OrderCreateParam) (model.Order, error)
	GetOrderByID (ctx context.Context, id string) (model.Order, error)
	UpdateOrder (ctx context.Context, params model.OrderUpdateParam) error
}