package part
import "github.com/H1dEx/go-rocket/inventory/internal/repository/model"
import def "github.com/H1dEx/go-rocket/inventory/internal/repository"
import "sync"

var _ def.InventoryRepository = (*repository)(nil)

type repository struct {
	mu sync.RWMutex
	parts map[string]model.Part
}

func NewPerository()*repository {
	return &repository{
		parts: make(map[string]model.Part),
	}
}