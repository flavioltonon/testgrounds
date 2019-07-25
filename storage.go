package testgrounds

import (
	"context"
	"fmt"
)

type Storage interface {
	Name() string
	Connect(ctx context.Context, options ...Option) (Client, error)
}

type Storages map[string]Storage

var storages = make(Storages, 0)

func RegisterStorage(storage Storage) {
	if _, exists := storages[storage.Name()]; exists {
		panic(fmt.Sprintf("storage %s has already been registered", storage.Name()))
	}

	storages[storage.Name()] = storage
}

func RegisteredStorages() Storages {
	return storages
}

func NewStorage(name string) (Storage, error) {
	if _, exists := storages[name]; !exists {
		return nil, fmt.Errorf("'%s' is not a registered storage type", name)
	}
	return storages[name], nil
}

func (s Storages) Contain(name string) bool {
	_, exists := s[name]
	return exists
}
