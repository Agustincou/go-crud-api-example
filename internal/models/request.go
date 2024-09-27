package models

import (
	"github.com/Agustincou/go-crud-api-example/internal/clients"
	"github.com/Agustincou/go-crud-api-example/pkg/enums"
)

type CreateProductRequestHeader struct {
	Owner []string `json:"owner"`
}

type CreateProductRequestBody struct {
	Name *string `json:"name,omitempty"`
	UpdateProductRequestBody
}

type UpdateProductRequestBody struct {
	Quantity *int     `json:"quantity,omitempty"`
	Price    *float64 `json:"price,omitempty"`
}

type CreateProductRequest struct {
	CreateProductRequestHeader
	CreateProductRequestBody
}

func (c *CreateProductRequest) MakeProduct(idGen clients.IDGenerator, time clients.Time) Product {
	return Product{
		ID:              idGen.GetID(),
		Name:            c.Name,
		Quantity:        c.Quantity,
		Price:           c.Price,
		Owner:           c.Owner[0],
		Status:          enums.Available,
		DateCreated:     time.Now(),
		DateLastUpdated: time.Now(),
	}
}

type SearchProductParams struct {
	Name   *string `json:"name,omitempty"`
	Offset int32   `json:"offset"`
	Limit  int32   `json:"limit"`
}
