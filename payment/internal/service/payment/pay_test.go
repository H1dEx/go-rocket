package payment

import (
	"regexp"

	"github.com/H1dEx/go-rocket/payment/internal/model"
	"github.com/brianvoe/gofakeit/v6"
)

func (s *ServiceSuite) TestPaySuccess() {
	var (
		userId        = gofakeit.UUID()
		paymentMethod = "CARD"
		orderId       = gofakeit.UUID()
	)
	res, err := s.service.PayOrder(s.ctx, orderId, userId, paymentMethod)

	s.NoError(err)
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	s.Regexp(uuidRegex, res)
}
func (s *ServiceSuite) TestPayUserIdFailed() {
	var (
		userId        = ""
		paymentMethod = "CARD"
		orderId       = gofakeit.UUID()
	)
	res, err := s.service.PayOrder(s.ctx, orderId, userId, paymentMethod)

	s.Error(err)
	s.ErrorIs(err, model.ErrUserIdInvalid)
	s.Empty(res)
}
func (s *ServiceSuite) TestPayPaymentMethodFailed() {
	var (
		userId        = gofakeit.UUID()
		paymentMethod = ""
		orderId       = gofakeit.UUID()
	)
	res, err := s.service.PayOrder(s.ctx, orderId, userId, paymentMethod)

	s.Error(err)
	s.ErrorIs(err, model.ErrPaymentMethodInvalid)
	s.Empty(res)
}

func (s *ServiceSuite) TestPayOrderIdFailed() {
	var (
		userId        = gofakeit.UUID()
		paymentMethod = "CARD"
		orderId       = ""
	)
	res, err := s.service.PayOrder(s.ctx, orderId, userId, paymentMethod)

	s.Error(err)
	s.ErrorIs(err, model.ErrOrderIdInvalid)
	s.Empty(res)
}
