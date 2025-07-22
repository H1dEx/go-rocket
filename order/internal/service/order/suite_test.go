package order

import (
	"context"
	"testing"

	repoMock "github.com/H1dEx/go-rocket/order/internal/repository/mocks"
	cliMock "github.com/H1dEx/go-rocket/order/internal/client/grpc/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context
	inventoryCli *cliMock.InventoryClient
	paymentCli *cliMock.PaymentClient
	repo *repoMock.OrderRepository
	service *service
}


func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()
	s.repo = repoMock.NewOrderRepository(s.T())

	s.inventoryCli = cliMock.NewInventoryClient(s.T())
	s.paymentCli = cliMock.NewPaymentClient(s.T())
	s.service = NewService(s.repo, s.inventoryCli, s.paymentCli)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
