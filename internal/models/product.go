package models

import (
	"time"

	"github.com/Agustincou/go-crud-api-example/pkg/enums"
)

type Product struct {
	ID              string              `json:"id"`
	Name            *string             `json:"name,omitempty"`
	Quantity        *int                `json:"quantity,omitempty"`
	Price           *float64            `json:"price,omitempty"`
	Owner           string              `json:"owner"`
	Status          enums.ProductStatus `json:"status"`
	DateCreated     time.Time           `json:"date_created"`
	DateLastUpdated time.Time           `json:"date_last_updated"`
}
