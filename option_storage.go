package testgrounds

const OPTION_TYPE_STORAGE = "storage-option"

func init() {
	RegisterOption(&StorageOption{})
}

type StorageOption struct {
	Name        string `json:"name"`
	Credentials []byte `json:"credentials"`
}

func (o StorageOption) Type() string {
	return OPTION_TYPE_STORAGE
}
