package update

import (
	"github.com/Agustincou/go-crud-api-example/internal/errors"
	"github.com/Agustincou/go-crud-api-example/internal/models"
)

type priceCheck struct {
	Validator
}

func (p *priceCheck) Check(req models.UpdateProductRequestBody) error {
	if req.Price == nil || *req.Price < 0.0 {
		return errors.InvalidPrice
	}

	return p.Chain.CheckNext(req)
}
