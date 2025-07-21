package converter

import (
	"time"

	"github.com/H1dEx/go-rocket/inventory/internal/model"
	inventory_v1 "github.com/H1dEx/go-rocket/shared/pkg/proto/inventory/v1"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PartCategoryToModel(category inventory_v1.Category) model.Category {
	switch category {
	case inventory_v1.Category_CATEGORY_ENGINE:
		return model.CategoryEngine
	case inventory_v1.Category_CATEGORY_FUEL:
		return model.CategoryFuel
	case inventory_v1.Category_CATEGORY_PORTHOLE:
		return model.CategoryPorthole
	case inventory_v1.Category_CATEGORY_WING:
		return model.CategoryWing
	default:
		return model.CategoryUnknown
	}
}

func PartCategoryToProto(category model.Category) inventory_v1.Category {
	switch category {
	case model.CategoryEngine:
		return inventory_v1.Category_CATEGORY_ENGINE
	case model.CategoryFuel:
		return inventory_v1.Category_CATEGORY_FUEL
	case model.CategoryPorthole:
		return inventory_v1.Category_CATEGORY_PORTHOLE
	case model.CategoryWing:
		return inventory_v1.Category_CATEGORY_WING
	default:
		return inventory_v1.Category_CATEGORY_UNKNOWN
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
func PartDimensionsToProto(d model.Dimensions) *inventory_v1.Dimensions {
	return &inventory_v1.Dimensions{
		Length: d.Length,
		Width:  d.Weight,
		Height: d.Height,
		Weight: d.Weight,
	}
}
func PartManufacturerToModel(m *inventory_v1.Manufacturer) model.Manufacturer {
	return model.Manufacturer{
		Name:    m.Name,
		Country: m.Country,
		Website: m.Website,
	}
}
func PartManufacturerToProto(m model.Manufacturer) *inventory_v1.Manufacturer {
	return &inventory_v1.Manufacturer{
		Name:    m.Name,
		Country: m.Country,
		Website: m.Website,
	}
}

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

func PartMetadataValueToProtoModel(m map[string]model.MetadataValue) map[string]*inventory_v1.Value {
	if m == nil {
		return nil
	}
	acc := map[string]*inventory_v1.Value{}
	for key, value := range m {
		switch {
		case value.StringValue != nil:
			acc[key] = &inventory_v1.Value{ValueType: &inventory_v1.Value_StringValue{StringValue: *value.StringValue}}
		case value.Int64Value != nil:
			acc[key] = &inventory_v1.Value{ValueType: &inventory_v1.Value_Int64Value{Int64Value: *value.Int64Value}}
		case value.BoolValue != nil:
			acc[key] = &inventory_v1.Value{ValueType: &inventory_v1.Value_BoolValue{BoolValue: *value.BoolValue}}
		case value.DoubleValue != nil:
			acc[key] = &inventory_v1.Value{ValueType: &inventory_v1.Value_DoubleValue{DoubleValue: *value.DoubleValue}}
		}
	}

	return acc
}

func PartToModel(p *inventory_v1.Part) *model.Part {
	if p == nil {
		return nil
	}
	var createdAt time.Time
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt.AsTime()
	}
	var updatedAt *time.Time
	if p.UpdatedAt != nil {
		updatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}
	return &model.Part{
		UUID:          p.Uuid,
		Name:          p.Name,
		Description:   p.Description,
		Price:         p.Price,
		StockQuantity: p.StockQuantity,
		Category:      PartCategoryToModel(p.Category),
		Dimensions:    PartDimensionsToModel(p.Dimensions),
		Manufacturer:  PartManufacturerToModel(p.Manufacturer),
		Tags:          p.Tags,
		Metadata:      PartMetadataValueToModel(p.Metadata),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func PartToProtoModel(p model.Part) *inventory_v1.Part {
	createdAt := timestamppb.New(p.CreatedAt)

	var updatedAt *timestamppb.Timestamp
	if p.UpdatedAt != nil {
		updatedAt = timestamppb.New(*p.UpdatedAt)
	}

	return &inventory_v1.Part{
		Uuid:          p.UUID,
		Name:          p.Name,
		Description:   p.Description,
		Price:         p.Price,
		StockQuantity: p.StockQuantity,
		Category:      PartCategoryToProto(p.Category),
		Dimensions:    PartDimensionsToProto(p.Dimensions),
		Manufacturer:  PartManufacturerToProto(p.Manufacturer),
		Tags:          p.Tags,
		Metadata:      PartMetadataValueToProtoModel(p.Metadata),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func ListPartsFilterToModel(f *inventory_v1.PartsFilter) model.FilterParts {
	if f == nil {
		return model.FilterParts{}
	}
	categories := []model.Category{}

	for _, v := range f.Categories {
		categories = append(categories, PartCategoryToModel(v))
	}
	return model.FilterParts{
		Uuids:                 f.Uuids,
		Names:                 f.Names,
		Categories:            categories,
		ManufacturerCountries: f.ManufacturerCountries,
		Tags:                  f.ManufacturerCountries,
	}
}
