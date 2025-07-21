package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/H1dEx/go-rocket/order/internal/converter"
	"github.com/H1dEx/go-rocket/order/internal/model"
	order_v1 "github.com/H1dEx/go-rocket/shared/pkg/openapi/order/v1"
)

func (a *api) PayOrderById(ctx context.Context, req *order_v1.PayOrderRequest, params order_v1.PayOrderByIdParams) (order_v1.PayOrderByIdRes, error) {
	if params.OrderUUID == "" {
		return &order_v1.BadRequestError{
			Code:    400,
			Message: "order id is empty",
		}, nil
	}

	if req.PaymentMethod == order_v1.PaymentMethodUNKNOWN {
		return &order_v1.BadRequestError{
			Code:    400,
			Message: "unknown payment method",
		}, nil
	}

	order, err := a.service.GetOrderByID(ctx, params.OrderUUID)

	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return &order_v1.NotFoundError{
				Code: 404,
				Message: fmt.Sprintf("Not found: %w", err),
			}, nil
		}
		return &order_v1.InternalServerError{
			Code: 500,
			Message: err.Error(),
		}, nil
	}

	if order.Status == model.OrderStatusCancelled {
		return &order_v1.BadRequestError{
			Code: 400,
			Message: fmt.Sprint("order %s already cancelled", params.OrderUUID), 
		}, nil
	}
	if order.Status == model.OrderStatusPaid {
		return &order_v1.BadRequestError{
			Code: 400,
			Message: fmt.Sprint("order %s already paid", params.OrderUUID), 
		}, nil
	}

	transactionId, err := a.service.PayOrderById(ctx, params.OrderUUID, order.UserUUID, converter.OrderPaymentToModel(req.PaymentMethod))
	if err != nil {
		return &order_v1.InternalServerError{
			Code: 500,
			Message: err.Error(),
		}, nil
	}

	return &order_v1.PayOrderResponse{
		TransactionUUID: transactionId,
	}, nil
}
