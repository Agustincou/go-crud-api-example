package create

import (
	"github.com/Agustincou/go-crud-api-example/internal/models"
	"github.com/Agustincou/go-crud-api-example/internal/validators"
)

type Validator struct {
	validators.Chain[models.CreateProductRequest]
}

func (c *Validator) Check(req models.CreateProductRequest) error {
	creationChecks := new(nameCheck)

	creationChecks.
		LinkWith(new(quantityCheck)).
		LinkWith(new(priceCheck)).
		LinkWith(new(headerCheck))

	return creationChecks.Check(req)
}
