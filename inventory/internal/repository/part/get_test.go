package part

import "github.com/brianvoe/gofakeit/v6"

func (s *ServiceSuite) TestGetPartNoOrderIdFailed() {
	uuid := ""
	res, err := s.repo.GetPart(s.ctx, uuid)

	s.Error(err)
	s.Empty(res)
}

func (s *ServiceSuite) TestGetPartNotFoundFailed() {
	uuid := gofakeit.UUID()

	res, err := s.repo.GetPart(s.ctx, uuid)

	s.Error(err)
	s.Empty(res)
}
