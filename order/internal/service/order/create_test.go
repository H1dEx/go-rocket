package order

import (
	"github.com/H1dEx/go-rocket/order/internal/model"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/mock"
)

func (s *ServiceSuite) TestCreateOrderPartsFailure() {
	var (
		partIds = []string{gofakeit.UUID(), gofakeit.UUID()}
		userId  = gofakeit.UUID()
	)

	s.inventoryCli.On("ListParts", s.ctx, model.FilterParts{
		Uuids: partIds,
	}).Return(nil, model.ErrPartsNotFound).Once()
	res, err := s.service.CreateOrder(s.ctx, model.OrderCreateParam{
		UserUUID:  userId,
		PartUuids: partIds,
	})
	s.Error(err)
	s.ErrorIs(err, model.ErrPartsNotFound)
	s.Empty(res)
}

func (s *ServiceSuite) TestCreateOrderPartsEmpty() {
	var (
		partIds = []string{gofakeit.UUID(), gofakeit.UUID()}
		userId  = gofakeit.UUID()
	)

	s.inventoryCli.On("ListParts", s.ctx, model.FilterParts{
		Uuids: partIds,
	}).Return([]model.Part{}, nil).Once()
	res, err := s.service.CreateOrder(s.ctx, model.OrderCreateParam{
		UserUUID:  userId,
		PartUuids: partIds,
	})
	s.Error(err)
	s.ErrorIs(err, model.ErrPartsNotFound)
	s.Empty(res)
}

func (s *ServiceSuite) TestCreateOrderPartsNotAllFound() {
	var (
		partIds = []string{gofakeit.UUID(), gofakeit.UUID(), gofakeit.UUID()}
		userId  = gofakeit.UUID()
	)

	s.inventoryCli.On("ListParts", s.ctx, model.FilterParts{
		Uuids: partIds,
	}).Return([]model.Part{
		{
			Price: gofakeit.Uint64(),
		},
		{
			Price: gofakeit.Uint64(),
		},
	}, nil).Once()
	res, err := s.service.CreateOrder(s.ctx, model.OrderCreateParam{
		UserUUID:  userId,
		PartUuids: partIds,
	})
	s.Error(err)
	s.ErrorIs(err, model.ErrPartsWrongAmoundFound)
	s.Empty(res)
}

func (s *ServiceSuite) TestCreateOrderPartsCreatingError() {
	var (
		partIds = []string{gofakeit.UUID(), gofakeit.UUID()}
		userId  = gofakeit.UUID()
	)

	s.inventoryCli.On("ListParts", s.ctx, model.FilterParts{
		Uuids: partIds,
	}).Return([]model.Part{
		{
			Price: gofakeit.Uint64(),
		},
		{
			Price: gofakeit.Uint64(),
		},
	}, nil).Once()
	s.repo.On("CreateOrder", s.ctx, mock.Anything).Return(model.Order{}, model.ErrOrderAlreadyExists)
	
	res, err := s.service.CreateOrder(s.ctx, model.OrderCreateParam{
		UserUUID:  userId,
		PartUuids: partIds,
	})
	s.Error(err)
	s.ErrorIs(err, model.ErrOrderAlreadyExists)
	s.Empty(res)
}

func (s *ServiceSuite) TestCreateOrderPartsSuccess() {
	var (
		partIds = []string{gofakeit.UUID(), gofakeit.UUID()}
		userId  = gofakeit.UUID()
		order = model.Order{
			UserUUID: userId,
			PartUuids: partIds,
			OrderUUID: mock.Anything,
			Status: model.OrderStatusPendingPayment,
			TotalPrice: gofakeit.Float32(),
		}
	)

	s.inventoryCli.On("ListParts", s.ctx, model.FilterParts{
		Uuids: partIds,
	}).Return([]model.Part{
		{
			Price: gofakeit.Uint64(),
		},
		{
			Price: gofakeit.Uint64(),
		},
	}, nil).Once()
	s.repo.On("CreateOrder", s.ctx, mock.Anything).Return(order, nil)
	
	res, err := s.service.CreateOrder(s.ctx, model.OrderCreateParam{
		UserUUID:  userId,
		PartUuids: partIds,
	})

	s.NoError(err)
	s.Equal(res, order)
}
