package order

import (
	"github.com/H1dEx/go-rocket/order/internal/model"
	"github.com/brianvoe/gofakeit/v6"
)

func (s *ServiceSuite) TestOrderCancelSuccess() {
	orderId := gofakeit.UUID()
	param := model.OrderUpdateParam{
		OrderId: orderId,
		Status: model.OrderStatusCancelled,
	}
	s.repo.On("UpdateOrder", s.ctx, param).Return(nil).Once()
	err := s.service.OrderCancelById(s.ctx, orderId)

	s.NoError(err)
}

func (s *ServiceSuite) TestOrderCancelFailure() {
	orderId := gofakeit.UUID()
	param := model.OrderUpdateParam{
		OrderId: orderId,
		Status: model.OrderStatusCancelled,
	}
	s.repo.On("UpdateOrder", s.ctx, param).Return(model.ErrOrderNotFound).Once()
	err := s.service.OrderCancelById(s.ctx, orderId)

	s.Error(err)
	s.ErrorIs(err, model.ErrOrderNotFound)
}
