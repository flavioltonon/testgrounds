package testgrounds

const OPTION_SUBTYPE_STORAGE_NAME = "storage-name"

type StorageOptionName struct {
	Name string `json:"name"`
}

func (o StorageOptionName) Type() string {
	return OPTION_TYPE_STORAGE
}

func (o StorageOptionName) Subtype() string {
	return OPTION_SUBTYPE_STORAGE_NAME
}
