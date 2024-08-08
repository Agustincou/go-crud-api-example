package create

import (
	"github.com/Agustincou/go-crud-api-example/internal/errors"
	"github.com/Agustincou/go-crud-api-example/internal/models"
)

type headerCheck struct {
	Validator
}

func (h *headerCheck) Check(req models.CreateProductRequest) error {
	if len(req.Owner) == 0 || req.Owner[0] == "" {
		return errors.InvalidHeader
	}

	return h.Chain.CheckNext(req)
}
