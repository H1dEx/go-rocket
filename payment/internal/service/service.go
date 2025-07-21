package service

import "context"

type PaymentService interface {
	PayOrder(ctx context.Context, orderId string, userId string, paymentMethod string) (transactionId string, err error)
}