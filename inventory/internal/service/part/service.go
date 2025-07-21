package part

import (
	def "github.com/H1dEx/go-rocket/inventory/internal/service"
	"github.com/H1dEx/go-rocket/inventory/internal/repository"
)

var _ def.PartService = (*service)(nil)

type service struct {
	repo repository.InventoryRepository
}

func NewService(repo repository.InventoryRepository) *service {
	return &service{
		repo: repo,
	}
}
