package order

import (
	"errors"

	"github.com/H1dEx/go-rocket/order/internal/model"
	"github.com/brianvoe/gofakeit/v6"
)

func (s *ServiceSuite) TestPayOrderSuccess() {
	var (
		orderId       = gofakeit.UUID()
		userId        = gofakeit.UUID()
		transactionId = gofakeit.UUID()
		paymentMethod = model.PaymentMethodCard
	)
	s.paymentCli.On("PayOrder", s.ctx, orderId, userId, paymentMethod).Return(transactionId, nil).Once()
	payData := model.OrderUpdateParam{
		OrderId:         orderId,
		PaymentMethod:   paymentMethod,
		TransactionUUID: transactionId,
		Status:          model.OrderStatusPaid,
	}
	s.repo.On("UpdateOrder", s.ctx, payData).Return(nil)
	res, err := s.service.PayOrderById(s.ctx, orderId, userId, paymentMethod)
	s.NoError(err)
	s.Equal(res, transactionId)
}

func (s *ServiceSuite) TestPayOrderUpdateFailure() {
	var (
		orderId       = gofakeit.UUID()
		userId        = gofakeit.UUID()
		transactionId = gofakeit.UUID()
		paymentMethod = model.PaymentMethodCard
	)
	s.paymentCli.On("PayOrder", s.ctx, orderId, userId, paymentMethod).Return(transactionId, nil).Once()
	payData := model.OrderUpdateParam{
		OrderId:         orderId,
		PaymentMethod:   paymentMethod,
		TransactionUUID: transactionId,
		Status:          model.OrderStatusPaid,
	}
	s.repo.On("UpdateOrder", s.ctx, payData).Return(model.ErrOrderNotFound)
	res, err := s.service.PayOrderById(s.ctx, orderId, userId, paymentMethod)
	s.Error(err)
	s.ErrorIs(err, model.ErrOrderNotFound)
	s.Empty(res)
}
func (s *ServiceSuite) TestPayOrderPayFailure() {
	var (
		orderId       = gofakeit.UUID()
		userId        = gofakeit.UUID()
		paymentMethod = model.PaymentMethodCard
		mockError     = errors.New("some error")
	)
	s.paymentCli.On("PayOrder", s.ctx, orderId, userId, paymentMethod).Return("", mockError).Once()
	res, err := s.service.PayOrderById(s.ctx, orderId, userId, paymentMethod)
	s.Error(err)
	s.ErrorIs(err, mockError)
	s.Empty(res)
}
