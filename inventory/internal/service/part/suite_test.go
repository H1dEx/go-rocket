package part

import (
	"context"
	"testing"

	"github.com/H1dEx/go-rocket/inventory/internal/model"
	"github.com/H1dEx/go-rocket/inventory/internal/repository/mocks"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	repo    *mocks.InventoryRepository
	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.repo = mocks.NewInventoryRepository(s.T())
	s.service = NewService(s.repo)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) GenPart() model.Part {
	return model.Part{
		UUID:          gofakeit.UUID(),
		Name:          gofakeit.Name(),
		Price:         gofakeit.Uint64(),
		StockQuantity: gofakeit.Int64(),
	}
}
