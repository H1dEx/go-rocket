package converter

import (
	"github.com/H1dEx/go-rocket/inventory/internal/model"
	repoModel "github.com/H1dEx/go-rocket/inventory/internal/repository/model"
)


func PartToRepoModel(part model.Part) repoModel.Part {
	metadata := make(map[string]repoModel.MetadataValue, len(part.Metadata)) 
	for key, v := range part.Metadata {
		metadata[key] = PartMetadataValueToRepoModel(v)
	}

	return repoModel.Part{
		UUID: part.UUID,
		Name: part.Name,
		Description: part.Description,
		Price: part.Price,
		StockQuantity: part.StockQuantity,
		Category: string(part.Category),
		Dimensions: PartDimensionsToRepoModel(part.Dimensions),
		Manufacturer: PartManufacturerToRepoModel(part.Manufacturer),
		Tags: part.Tags,
		Metadata: metadata,
		CreatedAt: part.CreatedAt,
		UpdatedAt: part.UpdatedAt,
	}
}
func PartToModel(part repoModel.Part) model.Part {
	metadata := make(map[string]model.MetadataValue, len(part.Metadata)) 
	for key, v := range part.Metadata {
		metadata[key] = PartMetadataValueToModel(v)
	}
	return model.Part{
		UUID: part.UUID,
		Name: part.Name,
		Description: part.Description,
		Price: part.Price,
		StockQuantity: part.StockQuantity,
		Category: PartCategoryToModel(part.Category),
		Dimensions: PartDimensionsToModel(part.Dimensions),
		Manufacturer: PartManufacturerToModel(part.Manufacturer),
		Tags: part.Tags,
		Metadata: metadata,
		CreatedAt: part.CreatedAt,
		UpdatedAt: part.UpdatedAt,
	}
}

func PartCategoryToModel(category string) model.Category{
	switch category {
	case string(model.CategoryEngine):
		return model.CategoryEngine
	case string(model.CategoryFuel):
		return model.CategoryFuel
	case string(model.CategoryPorthole):
		return model.CategoryPorthole
	case string(model.CategoryWing):
		return model.CategoryWing
	default:
		return model.CategoryUnknown
	}
}

func PartDimensionsToRepoModel(d model.Dimensions) repoModel.Dimensions{
	return repoModel.Dimensions{
		Length: d.Length,
		Width: d.Weight,
		Height: d.Height,
		Weight: d.Weight,
	}
}
func PartDimensionsToModel(d repoModel.Dimensions) model.Dimensions{
	return model.Dimensions{
		Length: d.Length,
		Width: d.Weight,
		Height: d.Height,
		Weight: d.Weight,
	}
}

func PartManufacturerToRepoModel(m model.Manufacturer) repoModel.Manufacturer{
	return repoModel.Manufacturer{
		Name: m.Name,
		Country: m.Country,
		Website: m.Website,
	}
}
func PartManufacturerToModel(m repoModel.Manufacturer) model.Manufacturer{
	return model.Manufacturer{
		Name: m.Name,
		Country: m.Country,
		Website: m.Website,
	}
}

func PartMetadataValueToRepoModel(m model.MetadataValue)repoModel.MetadataValue{
	return repoModel.MetadataValue{
		StringValue: m.StringValue,
		Int64Value: m.Int64Value,
		DoubleValue: m.DoubleValue,
		BoolValue: m.BoolValue,
	}
}
func PartMetadataValueToModel(m repoModel.MetadataValue)model.MetadataValue{
	return model.MetadataValue{
		StringValue: m.StringValue,
		Int64Value: m.Int64Value,
		DoubleValue: m.DoubleValue,
		BoolValue: m.BoolValue,
	}
}