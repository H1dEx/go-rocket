package order

import (
	"github.com/H1dEx/go-rocket/order/internal/model"
	repoModel "github.com/H1dEx/go-rocket/order/internal/repository/model"
)

func (s *ServiceSuite) TestCreateOrderSuccess() {
	fakeOrder := s.GenFakeRepoOrder()
	params := repoModel.OrderCreateParam{
		OrderUUID:  fakeOrder.OrderUUID,
		UserUUID:   fakeOrder.UserUUID,
		PartUuids:  fakeOrder.PartUuids,
		TotalPrice: fakeOrder.TotalPrice,
	}

	res, err := s.repo.CreateOrder(s.ctx, params)

	s.NoError(err)
	s.Equal(res.OrderUUID, params.OrderUUID)
	s.Equal(res.UserUUID, params.UserUUID)
	s.Equal(res.PartUuids, params.PartUuids)
	s.Equal(res.TotalPrice, params.TotalPrice)
	s.Equal(res.Status, model.OrderStatusPendingPayment)
}

func (s *ServiceSuite) TestCreateOrderFailed() {
	fakeOrder := s.GenFakeRepoOrder()
	s.repo.orders[fakeOrder.OrderUUID] = fakeOrder
	params := repoModel.OrderCreateParam{
		OrderUUID:  fakeOrder.OrderUUID,
		UserUUID:   fakeOrder.UserUUID,
		PartUuids:  fakeOrder.PartUuids,
		TotalPrice: fakeOrder.TotalPrice,
	}

	res, err := s.repo.CreateOrder(s.ctx, params)

	s.Error(err)
	s.ErrorIs(err, model.ErrOrderAlreadyExists)
	s.Empty(res)
}
