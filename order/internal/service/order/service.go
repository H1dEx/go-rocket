package order

import (
	"github.com/H1dEx/go-rocket/order/internal/client/grpc"
	"github.com/H1dEx/go-rocket/order/internal/repository"
	def "github.com/H1dEx/go-rocket/order/internal/service"
)

var _ def.OrderService = (*service)(nil)

type service struct {
	repo repository.OrderRepository

	inventoryClient grpc.InventoryClient
	paymentClient grpc.PaymentClient
}

func NewService(repo repository.OrderRepository, inventoryClient grpc.InventoryClient, paymentClient grpc.PaymentClient) *service {
	return &service{
		repo: repo,
		inventoryClient: inventoryClient,
		paymentClient: paymentClient,
	}
}