package order

import (
	"context"
	"log"

	"github.com/H1dEx/go-rocket/order/internal/model"
	sq "github.com/Masterminds/squirrel"
)

func (r *repository) UpdateOrder(ctx context.Context, params model.OrderUpdateParam) error {
	builderUpdate := sq.Update("orders").
		PlaceholderFormat(sq.Dollar).
		Set("transactionUuid", params.TransactionUUID).
		Set("status", params.Status).
		Set("paymentMethod", params.PaymentMethod).
		Where(sq.Eq{"orderUuid": params.OrderId})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v\n", err)
		return err
	}

	_, err = r.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to update note: %v\n", err)
		return err
	}

	return nil
}
