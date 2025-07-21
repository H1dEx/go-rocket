package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/H1dEx/go-rocket/order/internal/model"
	order_v1 "github.com/H1dEx/go-rocket/shared/pkg/openapi/order/v1"
)

func (a *api) OrderCancelById(ctx context.Context, params order_v1.OrderCancelByIdParams) (order_v1.OrderCancelByIdRes, error) {
	if params.OrderUUID == "" {
		return &order_v1.NotFoundError{
			Code:    403,
			Message: "order id is empty",
		}, nil
	}

	err := a.service.OrderCancelById(ctx, params.OrderUUID)

	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return &order_v1.NotFoundError{
				Code:    404,
				Message: fmt.Sprintf("sighting with UUID %s not found", params.OrderUUID),
			}, nil
		}
		if errors.Is(err, model.ErrOrderAlreadyExists) {
			return &order_v1.NotFoundError{
				Code:    409,
				Message: "Order already has been payed",
			}, nil
		}
		return nil, err
	}

	return &order_v1.OrderCancelByIdNoContent{}, nil
}
