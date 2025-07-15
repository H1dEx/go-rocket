package model

import "time"

type Category string

const (
	CategoryUnknown  Category = "UNKNOWN"
	CategoryEngine   Category = "ENGINE"
	CategoryFuel     Category = "FUEL"
	CategoryPorthole Category = "PORTHOLE"
	CategoryWing     Category = "WING"
)

type Dimensions struct {
	Length uint64
	Width  uint64
	Height uint64 
	Weight uint64 
}

type Manufacturer struct {
	Name    string 
	Country string 
	Website string 
}

type MetadataValue struct {
	StringValue *string  
	Int64Value  *int64   
	DoubleValue *float64 
	BoolValue   *bool    
}

type Part struct {
	UUID          string                  
	Name          string                  
	Description   string                  
	Price         uint64                   
	StockQuantity int64                   
	Category      Category                 
	Dimensions    Dimensions              
	Manufacturer  Manufacturer        
	Tags          []string                 
	Metadata      map[string]MetadataValue
	CreatedAt     time.Time               
	UpdatedAt     *time.Time               
}

type FilterParts struct {
	Uuids                 []string
	Names                 []string
	Categories            []Category
	ManufacturerCountries []string
	Tags                  []string
}
