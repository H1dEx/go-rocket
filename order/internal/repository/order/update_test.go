package order

// import (
// 	"github.com/H1dEx/go-rocket/order/internal/model"
// 	"github.com/brianvoe/gofakeit/v6"
// )

// func (s *ServiceSuite)TestUpdateOrderSuccess() {
// 	transactionId := gofakeit.UUID()
// 	order := s.GenFakeRepoOrder()
// 	s.repo.orders[order.OrderUUID] = order

// 	params := model.OrderUpdateParam{
// 		OrderId: order.OrderUUID,
// 		TransactionUUID: transactionId,
// 		Status:  model.OrderStatusPaid,
// 		PaymentMethod: model.PaymentMethodCard,
// 	}

// 	err := s.repo.UpdateOrder(s.ctx, params)

// 	s.NoError(err)
// }

// func (s *ServiceSuite)TestUpdateOrderNotFoundFailure() {
// 	transactionId := gofakeit.UUID()
// 	order := s.GenFakeRepoOrder()

// 	params := model.OrderUpdateParam{
// 		OrderId: order.OrderUUID,
// 		TransactionUUID: transactionId,
// 		Status:  model.OrderStatusPaid,
// 		PaymentMethod: model.PaymentMethodCard,
// 	}

// 	err := s.repo.UpdateOrder(s.ctx, params)

// 	s.Error(err)
// 	s.ErrorIs(err, model.ErrOrderNotFound)
// }