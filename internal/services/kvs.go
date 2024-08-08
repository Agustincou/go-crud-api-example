package services

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/Agustincou/go-crud-api-example/internal/clients"
	"github.com/Agustincou/go-crud-api-example/internal/models"
)

type KVS interface {
	SaveProduct(ctx context.Context, product models.Product) error
	GetProductByID(ctx context.Context, id string) (*models.Product, error)
}

type kvsImpl struct {
	client clients.KVS
}

func NewKVS(client clients.KVS) KVS {
	return &kvsImpl{
		client: client,
	}
}

func (k *kvsImpl) SaveProduct(ctx context.Context, product models.Product) error {
	if product.ID == "" {
		return errors.New("missing product ID")
	}

	return k.client.Set(ctx, product.ID, product)
}

func (k *kvsImpl) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	item, err := k.client.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if item.Key == "" {
		return nil, nil
	}

	product := models.Product{}
	unmErr := json.Unmarshal(item.Value, &product)
	if unmErr != nil {
		return nil, unmErr
	}

	return &product, nil
}
