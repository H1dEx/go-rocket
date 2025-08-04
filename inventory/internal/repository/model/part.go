package model

import "time"

type Dimensions struct {
	Length uint64 `bson:"length"`
	Width  uint64 `bson:"width"`
	Height uint64 `bson:"height"`
	Weight uint64 `bson:"weight"`
}

type Manufacturer struct {
	Name    string `bson:"name"`
	Country string `bson:"country"`
	Website string `bson:"website"`
}

type MetadataValue struct {
	StringValue *string
	Int64Value  *int64
	DoubleValue *float64
	BoolValue   *bool
}

type Part struct {
	UUID          string                   `bson:"_id"`
	Name          string                   `bson:"name"`
	Description   string                   `bson:"description, omitempty"`
	Price         uint64                   `bson:"price"`
	StockQuantity int64                    `bson:"quanity"`
	Category      string                   `bson:"category"`
	Dimensions    Dimensions               `bson:"dimensions, omitempty"`
	Manufacturer  Manufacturer             `bson:"manufacturer, omitempty"`
	Tags          []string                 `bson:"tags, omitempty"`
	Metadata      map[string]MetadataValue `bson:"metadata, omitempty"`
	CreatedAt     time.Time                `bson:"created_at"`
	UpdatedAt     *time.Time               `bson:"updated_at, omitempty"`
}
