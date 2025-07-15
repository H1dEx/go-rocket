package order

import (
	"sync"

	"github.com/H1dEx/go-rocket/order/internal/repository/model"
	dfo "github.com/H1dEx/go-rocket/order/internal/repository"
)
var _ dfo.OrderRepository = (*repository)(nil)
type repository struct {
	mu sync.RWMutex
	orders map[string]model.Order
}

func NewRepository() *repository {
	return &repository{
		orders: make(map[string]model.Order),
	}
}