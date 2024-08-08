package create

import (
	"github.com/Agustincou/go-crud-api-example/internal/errors"
	"github.com/Agustincou/go-crud-api-example/internal/models"
)

type quantityCheck struct {
	Validator
}

func (q *quantityCheck) Check(req models.CreateProductRequest) error {
	if req.Quantity == nil || *req.Quantity < 0 {
		return errors.InvalidQuantity
	}

	return q.Chain.CheckNext(req)
}
