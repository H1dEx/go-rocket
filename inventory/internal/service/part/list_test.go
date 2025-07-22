package part

import (
	"errors"

	"github.com/H1dEx/go-rocket/inventory/internal/model"
)

func (s *ServiceSuite) TestListPartsSuccess() {
	parts := []model.Part{s.GenPart(), s.GenPart()}
	s.repo.On("ListParts", s.ctx, model.FilterParts{}).Return(parts, nil).Once()
	
	res, err := s.service.ListParts(s.ctx, model.FilterParts{})

	s.NoError(err)
	s.Equal(res, parts)
}
func (s *ServiceSuite) TestListPartsFailure() {
	s.repo.On("ListParts", s.ctx, model.FilterParts{}).Return([]model.Part{}, errors.New("some error")).Once()
	
	res, err := s.service.ListParts(s.ctx, model.FilterParts{})

	s.Error(err)
	s.Empty(res)
}