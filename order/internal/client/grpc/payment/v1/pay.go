package v1

import (
	"context"

	"github.com/H1dEx/go-rocket/order/internal/client/converter"
	"github.com/H1dEx/go-rocket/order/internal/model"
	payment_v1 "github.com/H1dEx/go-rocket/shared/pkg/proto/payment/v1"
)

func (c *client) PayOrder(ctx context.Context, orderId, userId string, paymentMethod model.PaymentMethod) (string, error) {
	res, err := c.genCli.PayOrder(ctx, &payment_v1.PayOrderRequest{
		OrderUuid: orderId,
		UserUuid: userId,
		PaymentMethod: converter.OrderPaymentToPaymentMethod(paymentMethod),
	})

	if err != nil {
		return "", err
	}

	return res.TransactionUuid, nil
}
