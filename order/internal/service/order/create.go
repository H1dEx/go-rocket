package order

import (
	"context"

	"github.com/google/uuid"

	"github.com/H1dEx/go-rocket/order/internal/model"
	repoModel "github.com/H1dEx/go-rocket/order/internal/repository/model"
)

func (s *service) CreateOrder(ctx context.Context, body model.OrderCreateParam) (model.Order, error) {

	parts, err := s.inventoryClient.ListParts(ctx, model.FilterParts{
		Uuids: body.PartUuids,
	})

	if err != nil {
		return model.Order{}, err
	}

	if len(parts) == 0 {
		return model.Order{}, model.ErrPartsNotFound
	}

	if len(parts) != len(body.PartUuids) {
		return model.Order{}, model.ErrPartsWrongAmoundFound
	}

	var totalPrice int
	for _, part := range parts {
		totalPrice += int(part.Price)
	}
	orderUUID := uuid.NewString()

	createParams := repoModel.OrderCreateParam{
		UserUUID:   body.UserUUID,
		PartUuids:  body.PartUuids,
		OrderUUID:  orderUUID,
		Status:     repoModel.OrderStatusPendingPayment,
		TotalPrice: float32(totalPrice),
	}

	order, err := s.repo.CreateOrder(ctx, createParams)

	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}
