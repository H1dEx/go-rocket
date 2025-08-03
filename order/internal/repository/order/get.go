package order

import (
	"context"
	"errors"
	"log"

	"github.com/H1dEx/go-rocket/order/internal/model"
	"github.com/H1dEx/go-rocket/order/internal/repository/converter"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *repository) GetOrderByID(ctx context.Context, orderId string) (model.Order, error) {
	selectBuild := sq.Select("orderUuid", "userUuid", "partUuids", "totalPrice", "transactionUuid", "paymentMethod", "status").From("orders").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"orderUuid": orderId}).
		Limit(1)

	query, args, err := selectBuild.ToSql()

	if err != nil {
		log.Printf("failed to build query: %v\n", err)
		return model.Order{}, errors.New("failed to build query:" + err.Error())
	}

	var orderUuid, userUuid string
	var transactionUuid, paymentMethod, status *string
	var partUuids pgtype.Array[string]
	var totalPrice float32

	err = r.pool.QueryRow(ctx, query, args...).Scan(&orderUuid, &userUuid, &partUuids, &totalPrice, &transactionUuid, &paymentMethod, &status)
	if err != nil {
		log.Printf("failed to scan order: %v", err)
		return model.Order{}, err
	}

	if !partUuids.Valid {
		return model.Order{}, errors.New("error while parsing partUuids")
	}

	partUidsSlice := make([]string, 0, len(partUuids.Elements))

	partUidsSlice = append(partUidsSlice, partUuids.Elements...)

	var transactionUuidStr, paymentMethodStr, statusStr string

	if transactionUuid != nil {
		transactionUuidStr = *transactionUuid
	}

	if paymentMethod != nil {
		paymentMethodStr = *paymentMethod
	}

	if status != nil {
		statusStr = *status
	}
	order := model.Order{
		OrderUUID:       orderUuid,
		UserUUID:        userUuid,
		PartUuids:       partUidsSlice,
		TotalPrice:      totalPrice,
		TransactionUUID: transactionUuidStr,
		PaymentMethod:   model.PaymentMethod(converter.OrderPaymentMethodFromDb(paymentMethodStr)),
		Status:          model.OrderStatus(converter.OrderStatusFromDb(statusStr)),
	}

	return order, nil
}
