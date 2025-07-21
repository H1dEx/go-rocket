package v1

import (
	inventory_v1 "github.com/H1dEx/go-rocket/shared/pkg/proto/inventory/v1"
	def "github.com/H1dEx/go-rocket/order/internal/client/grpc"
)

var _ def.InventoryClient = (*client)(nil)

type client struct {
	genClient inventory_v1.InventoryServiceClient
}

func NewClient(cli inventory_v1.InventoryServiceClient) *client {
	return &client{
		genClient: cli,
	}
}