package testgrounds

import "fmt"

type Storage interface {
	Name() string
	Connect(options ...Option) (Client, error)
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

func (s Storages) ByName(name string) (Storage, error) {
	if _, exists := s[name]; !exists {
		return nil, fmt.Errorf("'%s' is not a registered storage type", name)
	}
	return s[name], nil
}

func (s Storages) Contain(name string) bool {
	_, exists := s[name]
	return exists
}
