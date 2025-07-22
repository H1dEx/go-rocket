package order

import (
	"github.com/H1dEx/go-rocket/order/internal/model"
	"github.com/brianvoe/gofakeit/v6"
)

func (s *ServiceSuite) TestGetOrderByIDFailure() {
	orderId := gofakeit.UUID()
	s.repo.On("GetOrderByID", s.ctx, orderId).Return(model.Order{}, model.ErrOrderNotFound).Once()
	res, err := s.service.GetOrderByID(s.ctx, orderId)
	s.Error(err)
	s.ErrorIs(err, model.ErrOrderNotFound)
	s.Empty(res)
}

func (s *ServiceSuite) TestGetOrderByIDSuccess() {
	orderId := gofakeit.UUID()
	s.repo.On("GetOrderByID", s.ctx, orderId).Return(model.Order{
		OrderUUID: orderId,
	}, nil).Once()
	res, err := s.service.GetOrderByID(s.ctx, orderId)
	s.NoError(err)
	s.Equal(res.OrderUUID, orderId)
}
