package order

import (
	dfo "github.com/H1dEx/go-rocket/order/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)
var _ dfo.OrderRepository = (*repository)(nil)
type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: pool,
	}
}