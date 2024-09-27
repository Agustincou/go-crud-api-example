package update

import (
	"github.com/Agustincou/go-crud-api-example/internal/models"
	"github.com/Agustincou/go-crud-api-example/internal/validators"
)

type Validator struct {
	validators.Chain[models.UpdateProductRequestBody]
}

func (v *Validator) Check(req models.UpdateProductRequestBody) error {
	checks := new(quantityCheck)

	checks.LinkWith(new(priceCheck))

	return checks.Check(req)
}
