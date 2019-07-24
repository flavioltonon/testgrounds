package testgrounds

const ANOTHER_STORAGE = "another-storage"

type AnotherStorage struct{}

func init() {
	RegisterStorage(&AnotherStorage{})
}

func (s AnotherStorage) Name() string {
	return ANOTHER_STORAGE
}

func (s AnotherStorage) Connect(options ...Option) (Client, error) {
	return nil, nil
}
