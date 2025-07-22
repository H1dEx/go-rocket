package part

import (
	"github.com/H1dEx/go-rocket/inventory/internal/model"
	"github.com/brianvoe/gofakeit/v6"
)

func (s *ServiceSuite) TestGetPartSuccess() {
	part := s.GenPart()
	uuid := part.UUID
	s.repo.On("GetPart", s.ctx, uuid).Return(part, nil).Once()

	resPart, err := s.service.GetPart(s.ctx, uuid)

	s.NoError(err)
	s.Equal(resPart, part)
}
func (s *ServiceSuite) TestGetPartFailure() {
	uuid := gofakeit.UUID()

	s.repo.On("GetPart", s.ctx, uuid).Return(model.Part{}, model.ErrPartNotFound).Once()

	res, err := s.service.GetPart(s.ctx, uuid)

	s.Error(err)
	s.ErrorIs(err, model.ErrPartNotFound)
	s.Empty(res)
}
