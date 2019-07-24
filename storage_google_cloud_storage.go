package testgrounds

import (
	"context"

	"google.golang.org/api/option"

	"cloud.google.com/go/storage"
)

const GOOGLE_CLOUD_STORAGE = "google-cloud-storage"

type GoogleCloudStorage struct{}

func init() {
	RegisterStorage(&GoogleCloudStorage{})
}

func (s GoogleCloudStorage) Name() string {
	return GOOGLE_CLOUD_STORAGE
}

func (s GoogleCloudStorage) Connect(options ...Option) (Client, error) {
	var opts = make([]option.ClientOption, 0)

	// opts = append(opts, option.WithCredentials(&google.Credentials{}))
	opts = append(opts, option.WithoutAuthentication())

	return storage.NewClient(context.Background(), opts...)
}
