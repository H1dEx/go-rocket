package v1

import (
	"context"

	payment_v1 "github.com/H1dEx/go-rocket/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) PayOrder(ctx context.Context, req *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error) {
	if req.GetOrderUuid() == "" {
		return nil, status.Error(codes.InvalidArgument, "field order_uuid is empty")
	}
	if req.GetPaymentMethod() == payment_v1.PaymentMethod_PAYMENT_METHOD_UNKNOWN {
		return nil, status.Error(codes.InvalidArgument, "field payment_metod is unknown")
	}
	if req.GetUserUuid() == "" {
		return nil, status.Error(codes.InvalidArgument, "field user_uuid is empty")
	}

	transaction_uuid, err := a.service.PayOrder(ctx, req.OrderUuid, req.UserUuid, string(req.PaymentMethod))
	if err != nil {
		return nil, err
	}

	return &payment_v1.PayOrderResponse{
		TransactionUuid: transaction_uuid,
	}, nil
}
