package models

import (
	"time"

	"github.com/Agustincou/go-crud-api-example/pkg/enums"
)

type Product struct {
	ID              string
	Name            *string
	Quantity        *int
	Price           *float64
	Owner           string
	Status          enums.ProductStatus
	DateCreated     time.Time
	DateLastUpdated time.Time
}
