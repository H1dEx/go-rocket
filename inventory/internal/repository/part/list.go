package part

import (
	"context"
	"slices"

	"github.com/H1dEx/go-rocket/inventory/internal/model"
	"github.com/H1dEx/go-rocket/inventory/internal/repository/converter"
	repoModel "github.com/H1dEx/go-rocket/inventory/internal/repository/model"
)

func validatePart(p *repoModel.Part, filters *model.FilterParts) bool {
	if len(filters.Uuids) > 0 && !slices.Contains(filters.Uuids, p.UUID) {
		return false
	}

	if len(filters.Names) > 0 && !slices.Contains(filters.Names, p.Name) {
		return false
	}

	if len(filters.Categories) > 0 {
		match := false
		for _, category := range filters.Categories {
			if string(category) == p.Category {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}

	if len(filters.ManufacturerCountries) > 0 && !slices.Contains(filters.ManufacturerCountries, p.Manufacturer.Country) {
		return false
	}

	if len(filters.Tags) > 0 {
		match := false
		for _, tag := range filters.Tags {
			if slices.Contains(p.Tags, tag) {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}

	return true
}

func (r *repository) ListParts(context context.Context, filters model.FilterParts) ([]model.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var filteredParts []model.Part
	for _, part := range r.parts {
		if validatePart(&part, &filters) {
			formatted := converter.PartToModel(part)
			filteredParts = append(filteredParts, formatted)
		}
	}

	return filteredParts, nil
}
