package order

import (
	"context"

	"github.com/H1dEx/go-rocket/order/internal/model"
)

func (s *service) PayOrderById(ctx context.Context, orderId, userId string, paymentMethod model.PaymentMethod) (transactionId string, err error) {
	transactionId, err = s.paymentClient.PayOrder(ctx, orderId, userId, paymentMethod)
	if err != nil {
		return "", err
	}
	payData := model.OrderUpdateParam{
		OrderId: orderId,
		PaymentMethod: paymentMethod,
		TransactionUUID: transactionId,
		Status: model.OrderStatusPaid,
	}
	err = s.repo.UpdateOrder(ctx, payData)
	if err != nil {
		return "", err
	}
	return transactionId, nil
}