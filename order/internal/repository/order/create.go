package order

import (
	"context"
	"log"

	"github.com/H1dEx/go-rocket/order/internal/model"
	"github.com/H1dEx/go-rocket/order/internal/repository/converter"
	repoModel "github.com/H1dEx/go-rocket/order/internal/repository/model"
	sq "github.com/Masterminds/squirrel"
)

func (r *repository) CreateOrder(ctx context.Context, params repoModel.OrderCreateParam) (model.Order, error) {
	builderInsert := sq.Insert("orders").PlaceholderFormat(sq.Dollar).
		Columns("orderUuid", "userUuid", "partUuids", "totalPrice", "status").
		Values(params.OrderUUID, params.UserUUID, params.PartUuids, params.TotalPrice, repoModel.OrderStatusPendingPayment).
		Suffix("RETURNING orderUuid, userUuid, partUuids, totalPrice, status")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v\n", err)
		return model.Order{}, err
	}

	var (
		orderUUID  string
		userUUID   string
		partUuids  []string // или pgtype.UUIDArray, если используется pgx
		totalPrice float32
		status     string
	)
	
	err = r.pool.QueryRow(ctx, query, args...).Scan(&orderUUID, &userUUID, &partUuids, &totalPrice, &status)

	if err != nil {
		log.Printf("failed to insert note: %v\n", err)
		return model.Order{}, err
	}

	return model.Order{
		OrderUUID:  orderUUID,
		UserUUID:   userUUID,
		PartUuids:  partUuids,
		TotalPrice: totalPrice,
		Status:     model.OrderStatus(converter.OrderStatusFromDb(status)),
	}, nil
}
