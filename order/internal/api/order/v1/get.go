package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/H1dEx/go-rocket/order/internal/converter"
	"github.com/H1dEx/go-rocket/order/internal/model"
	order_v1 "github.com/H1dEx/go-rocket/shared/pkg/openapi/order/v1"
)

func (a *api) GetOrderByID(ctx context.Context, params order_v1.GetOrderByIDParams) (order_v1.GetOrderByIDRes, error) {
	if params.OrderUUID == "" {
		return &order_v1.NotFoundError{
			Code:    403,
			Message: "order UUID is empty",
		}, nil
	}

	order, err := a.service.GetOrderByID(ctx, params.OrderUUID)

	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return &order_v1.NotFoundError{
				Code:    404,
				Message: fmt.Sprintf("order with UUID %s not found", params.OrderUUID),
			}, nil
		}
		return nil, err
	}

	return &order_v1.GetOrderResponse{
		Order: order_v1.OptOrderDto{Value: converter.OrderToDtoModel(order), Set: true},
	}, nil
}
