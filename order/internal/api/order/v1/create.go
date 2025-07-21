package v1

import (
	"context"

	"github.com/H1dEx/go-rocket/order/internal/model"
	order_v1 "github.com/H1dEx/go-rocket/shared/pkg/openapi/order/v1"
)

func (a *api)CreateOrder(ctx context.Context, req *order_v1.CreateOrderRequest) (order_v1.CreateOrderRes, error) {
	if req.UserUUID == "" {
		return &order_v1.BadRequestError{
			Code:    400,
			Message: "user UUID is empty",
		}, nil
	}
	if len(req.PartUuids) == 0 {
		return &order_v1.BadRequestError{
			Code:    400,
			Message: "part uuids is empty",
		}, nil
	}

	order, err := a.service.CreateOrder(ctx, model.OrderCreateParam{
		PartUuids: req.PartUuids,
		UserUUID: req.UserUUID,
	})

	if err != nil {
		return nil, err
	}
	return &order_v1.CreateOrderResponse{
		OrderUUID: order.OrderUUID,
		TotalPrice: order.TotalPrice,
	}, nil	
}