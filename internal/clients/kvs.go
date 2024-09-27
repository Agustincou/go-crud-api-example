package clients

import (
	"context"
	"encoding/json"
)

var (
	data = map[string]Item{}
)

type Item struct {
	Key   string
	Value []byte
}

type KVS interface {
	Set(context.Context, string, any) error
	Get(context.Context, string) (Item, error)
	Delete(context.Context, string) (bool, error)
}

type inMemoryKVS struct {
	KVS
}

func NewInMemoryKVS() KVS {
	return &inMemoryKVS{}
}

func (k *inMemoryKVS) Set(_ context.Context, key string, value any) error {
	valueBytes, err := json.Marshal(value)

	if err != nil {
		return err
	}

	data[key] = Item{
		Key:   key,
		Value: valueBytes,
	}

	return nil
}

func (k *inMemoryKVS) Get(_ context.Context, key string) (Item, error) {
	return data[key], nil
}

func (k *inMemoryKVS) Delete(_ context.Context, key string) (bool, error) {
	if _, exists := data[key]; exists {
		newData := map[string]Item{}

		for dataKey, dataValue := range data {
			if dataKey != key {
				newData[key] = dataValue
			}
		}

		data = newData
	}

	return true, nil
}
