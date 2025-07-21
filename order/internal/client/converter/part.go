package converter

import (
	"time"

	"github.com/H1dEx/go-rocket/order/internal/model"
	"github.com/samber/lo"

	inventory_v1 "github.com/H1dEx/go-rocket/shared/pkg/proto/inventory/v1"
)

func PartMetadataValueToModel(m map[string]*inventory_v1.Value) map[string]model.MetadataValue {
	if m == nil {
		return nil
	}
	acc := make(map[string]model.MetadataValue)

	for key, value := range m {
		switch v := value.GetValueType().(type) {
		case *inventory_v1.Value_StringValue:
			acc[key] = model.MetadataValue{StringValue: &v.StringValue}
		case *inventory_v1.Value_Int64Value:
			acc[key] = model.MetadataValue{Int64Value: &v.Int64Value}
		case *inventory_v1.Value_BoolValue:
			acc[key] = model.MetadataValue{BoolValue: &v.BoolValue}
		case *inventory_v1.Value_DoubleValue:
			acc[key] = model.MetadataValue{DoubleValue: &v.DoubleValue}
		default:
			continue
		}
	}
	return acc
}

func PartManufacturerToModel(m *inventory_v1.Manufacturer) model.Manufacturer {
	return model.Manufacturer{
		Name:    m.Name,
		Country: m.Country,
		Website: m.Website,
	}
}

func PartDimensionsToModel(d *inventory_v1.Dimensions) model.Dimensions {
	return model.Dimensions{
		Length: d.Length,
		Width:  d.Weight,
		Height: d.Height,
		Weight: d.Weight,
	}
}

func PartToModel(p *inventory_v1.Part) model.Part {
	var createdAt time.Time
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt.AsTime()
	}
	var updatedAt *time.Time
	if p.UpdatedAt != nil {
		updatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}
	return model.Part{
		UUID:          p.Uuid,
		Name:          p.Name,
		Description:   p.Description,
		Price:         p.Price,
		StockQuantity: p.StockQuantity,
		Category:      string(p.Category),
		Dimensions:    PartDimensionsToModel(p.Dimensions),
		Manufacturer:  PartManufacturerToModel(p.Manufacturer),
		Tags:          p.Tags,
		Metadata:      PartMetadataValueToModel(p.Metadata),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func PartListToModel(partList []*inventory_v1.Part) []model.Part {
	res := make([]model.Part, 0, len(partList))

	for _, v := range partList {
		res = append(res, PartToModel(v))
	}

	return res
}

func PartCategoryToProto(category string) inventory_v1.Category {
	switch category {
	case string(inventory_v1.Category_CATEGORY_ENGINE):
		return inventory_v1.Category_CATEGORY_ENGINE
	case string(inventory_v1.Category_CATEGORY_FUEL):
		return inventory_v1.Category_CATEGORY_FUEL
	case string(inventory_v1.Category_CATEGORY_PORTHOLE):
		return inventory_v1.Category_CATEGORY_PORTHOLE
	case string(inventory_v1.Category_CATEGORY_WING):
		return inventory_v1.Category_CATEGORY_WING
	default:
		return inventory_v1.Category_CATEGORY_UNKNOWN
	}
}

func PartsFilterToProto(filters model.FilterParts) *inventory_v1.PartsFilter {
	categories := make([]inventory_v1.Category, 0, len(filters.Categories))
	for _, v := range filters.Categories {
		categories = append(categories, PartCategoryToProto(v))
	}
	return &inventory_v1.PartsFilter{
		Uuids:                 filters.Uuids,
		Names:                 filters.Names,
		Categories:            categories,
		ManufacturerCountries: filters.ManufacturerCountries,
		Tags:                  filters.Tags,
	}
}
