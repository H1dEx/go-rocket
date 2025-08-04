package part

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	repo *repository
}


func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()
	// s.repo = NewRepository()
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
