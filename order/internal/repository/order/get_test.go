package order

import (
	"github.com/H1dEx/go-rocket/order/internal/repository/converter"
	"github.com/brianvoe/gofakeit/v6"
)

func (s *ServiceSuite) TestGetOrderSuccess() {
	order := s.GenFakeRepoOrder()
	s.repo.orders[order.OrderUUID] = order

	res, err := s.repo.GetOrderByID(s.ctx, order.OrderUUID)
	s.NoError(err)
	s.Equal(res, converter.OrderToModel(order))
}
func (s *ServiceSuite) TestGetOrderFailure() {
	order := s.GenFakeRepoOrder()
	s.repo.orders[order.OrderUUID] = order

	res, err := s.repo.GetOrderByID(s.ctx, gofakeit.UUID())
	s.Error(err)
	s.Empty(res)
}
