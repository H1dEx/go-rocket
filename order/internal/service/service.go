package service

import (
	"context"

	"github.com/H1dEx/go-rocket/order/internal/model"
)

type OrderService interface {
	CreateOrder(ctx context.Context, body model.OrderCreateParam) (model.Order, error)
	GetOrderByID(ctx context.Context, id string) (model.Order, error)
	OrderCancelById(ctx context.Context, id string) error
	PayOrderById(ctx context.Context, orderId, userId string, paymentMethod model.PaymentMethod) (transactionId string, err error)
}