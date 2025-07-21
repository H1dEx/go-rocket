package grpc

import (
	"context"

	"github.com/H1dEx/go-rocket/order/internal/model"
)

type InventoryClient interface {
	ListParts(ctx context.Context, filter model.FilterParts) ([]model.Part, error)
}

type PaymentClient interface {
	PayOrder(ctx context.Context, orderId, userId string, paymentMethod model.PaymentMethod) (string, error)
}