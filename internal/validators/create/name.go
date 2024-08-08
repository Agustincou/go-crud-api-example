package create

import (
	"github.com/Agustincou/go-crud-api-example/internal/errors"
	"github.com/Agustincou/go-crud-api-example/internal/models"
)

type nameCheck struct {
	Validator
}

func (n *nameCheck) Check(req models.CreateProductRequest) error {
	if req.Quantity == nil || *req.Quantity < 0 {
		return errors.InvalidQuantity
	}

	return n.Chain.CheckNext(req)
}
