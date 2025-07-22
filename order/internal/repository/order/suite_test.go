package order

import (
	"context"
	"testing"

	"github.com/H1dEx/go-rocket/order/internal/model"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"
	repoModel "github.com/H1dEx/go-rocket/order/internal/repository/model"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	repo *repository
}


func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()
	s.repo = NewRepository()
}

func (s *ServiceSuite) TearDownTest() {
}

func (s *ServiceSuite) GenFakeOrder() model.Order {
	return model.Order{
		OrderUUID: gofakeit.UUID(),
		UserUUID: gofakeit.UUID(),
		PartUuids: []string{gofakeit.UUID(), gofakeit.UUID(), gofakeit.UUID()},
		TotalPrice: gofakeit.Float32(),
		TransactionUUID: gofakeit.UUID(),
		PaymentMethod: model.PaymentMethodUnknown,
		Status: model.OrderStatusPendingPayment,
	}
}

func (s *ServiceSuite) GenFakeRepoOrder() repoModel.Order {
	return repoModel.Order{
		OrderUUID: gofakeit.UUID(),
		UserUUID: gofakeit.UUID(),
		PartUuids: []string{gofakeit.UUID(), gofakeit.UUID(), gofakeit.UUID()},
		TotalPrice: gofakeit.Float32(),
		TransactionUUID: gofakeit.UUID(),
		PaymentMethod: repoModel.PaymentMethodUnknown,
		Status: repoModel.OrderStatusPendingPayment,
	}
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
