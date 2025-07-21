package payment

import (
	"context"
	"log"

	"github.com/H1dEx/go-rocket/payment/internal/model"
	"github.com/google/uuid"
)

func (s *service) PayOrder(ctx context.Context, orderId string, userId string, paymentMethod string) (transactionId string, err error) {
	if userId == "" {
		return "", model.ErrUserIdInvalid
	}
	if paymentMethod == "" {
		return "", model.ErrPaymentMethodInvalid
	}
	if orderId == "" {
		return "", model.ErrOrderIdInvalid
	}

	transactionUUID := uuid.NewString()
	log.Printf("Оплата прошла успешно, transaction_uuid: %s", transactionUUID)

	return transactionUUID, nil
}
