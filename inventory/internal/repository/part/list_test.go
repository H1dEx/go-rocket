package part

import (
	"fmt"

	"github.com/H1dEx/go-rocket/inventory/internal/model"
	repoModel "github.com/H1dEx/go-rocket/inventory/internal/repository/model"
	"github.com/brianvoe/gofakeit/v6"
)

func (s *ServiceSuite) TestListPartsEmpty() {
	notExistingName := gofakeit.UUID()
	parts, err := s.repo.ListParts(s.ctx, model.FilterParts{Names: []string{notExistingName}})

	s.NoError(err)
	s.Empty(parts)
}

func (s *ServiceSuite) TestValidatePartAllSuccess() {
	var (
		uuid    = gofakeit.UUID()
		name    = gofakeit.Product().Name
		country = gofakeit.Country()
		tagOne  = gofakeit.City()
		tagTwo  = gofakeit.City()
		part    = repoModel.Part{
			UUID:     uuid,
			Name:     name,
			Category: string(model.CategoryEngine),
			Manufacturer: repoModel.Manufacturer{
				Country: country,
			},
			Tags: []string{tagOne, tagTwo},
		}
		filter = &model.FilterParts{
			Uuids:                 []string{uuid, gofakeit.UUID()},
			Names:                 []string{name, gofakeit.Product().Name},
			Categories:            []model.Category{model.CategoryEngine, model.CategoryFuel},
			ManufacturerCountries: []string{country, gofakeit.Country()},
			Tags:                  []string{tagTwo},
		}
	)

	res := validatePart(&part, filter)
	s.True(res)
}
func (s *ServiceSuite) TestValidatePartEmptySuccess() {
	var (
		uuid    = gofakeit.UUID()
		name    = gofakeit.Product().Name
		country = gofakeit.Country()
		tagOne  = gofakeit.City()
		tagTwo  = gofakeit.City()
		part    = repoModel.Part{
			UUID:     uuid,
			Name:     name,
			Category: string(model.CategoryEngine),
			Manufacturer: repoModel.Manufacturer{
				Country: country,
			},
			Tags: []string{tagOne, tagTwo},
		}
		filter = &model.FilterParts{}
	)

	res := validatePart(&part, filter)
	s.True(res)
}
func (s *ServiceSuite) TestValidatePartIdFailure() {
	var (
		uuid       = gofakeit.UUID()
		uuidFilter = fmt.Sprintf("%s filter", uuid)
		part       = repoModel.Part{
			UUID: uuid,
		}
		filter = &model.FilterParts{
			Uuids: []string{uuidFilter},
		}
	)

	res := validatePart(&part, filter)
	s.False(res)
}

func (s *ServiceSuite) TestValidatePartNameFailure() {
	var (
		name       = gofakeit.Product().Name
		nameFilter = fmt.Sprintf("%s filter", name)
		part       = repoModel.Part{
			Name: name,
		}
		filter = &model.FilterParts{
			Names: []string{nameFilter},
		}
	)

	res := validatePart(&part, filter)
	s.False(res)
}

func (s *ServiceSuite) TestValidatePartCategoryFailure() {
	var (
		part = repoModel.Part{
			Category: string(model.CategoryEngine),
		}
		filter = &model.FilterParts{
			Categories: []model.Category{model.CategoryFuel},
		}
	)

	res := validatePart(&part, filter)
	s.False(res)
}
func (s *ServiceSuite) TestValidatePartCountryFailure() {
	var (
		country       = gofakeit.Country()
		countryFilter = fmt.Sprintf("%s filter", country)

		part = repoModel.Part{
			Manufacturer: repoModel.Manufacturer{
				Country: country,
			},
		}
		filter = &model.FilterParts{
			ManufacturerCountries: []string{countryFilter},
		}
	)

	res := validatePart(&part, filter)
	s.False(res)
}

func (s *ServiceSuite) TestValidatePartTagsFailure() {
	var (
		tagOne    = gofakeit.City()
		tagTwo    = gofakeit.Country()
		tagFilter = gofakeit.Car().Brand
		part      = repoModel.Part{
			Tags: []string{tagOne, tagTwo},
		}
		filter = &model.FilterParts{
			Tags: []string{tagFilter},
		}
	)

	res := validatePart(&part, filter)
	s.False(res)
}
