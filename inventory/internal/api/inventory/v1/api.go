package v1

import (
	"github.com/H1dEx/go-rocket/inventory/internal/service"
	inventory_v1 "github.com/H1dEx/go-rocket/shared/pkg/proto/inventory/v1"
)

type api struct {
	inventory_v1.UnimplementedInventoryServiceServer
	service service.PartService
}

func NewApi(service service.PartService) *api {
	return &api{
		service: service,
	}
}