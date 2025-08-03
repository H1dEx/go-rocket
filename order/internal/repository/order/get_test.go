package order

// import (
// 	"fmt"

// 	"github.com/brianvoe/gofakeit/v6"
// )

// func (s *ServiceSuite) TestGetOrderSuccess() {
// 	order := s.GenFakeOrder()
// 	_, err := s.InsertOrder(order)

// 	s.NoError(err)
// 	// s.repo.orders[order.OrderUUID] = order

// 	res, err := s.repo.GetOrderByID(s.ctx, order.OrderUUID)
// 	s.NoError(err)
// 	fmt.Print(res)
// 	s.Equal(res, order)
// }
// func (s *ServiceSuite) TestGetOrderFailure() {
// 	order := s.GenFakeOrder()
// 	_, err := s.InsertOrder(order)
// 	s.NoError(err)

// 	res, err := s.repo.GetOrderByID(s.ctx, gofakeit.UUID())
// 	s.Error(err)
// 	s.Empty(res)
// }
