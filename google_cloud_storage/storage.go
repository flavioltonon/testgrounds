package google_cloud_storage

import (
	"context"
	"flavioltonon/testgrounds"

	"golang.org/x/oauth2/google"

	"google.golang.org/api/option"
)

const STORAGE_NAME = "google-cloud-storage"

type GoogleCloudStorage struct{}

func init() {
	testgrounds.RegisterStorage(&GoogleCloudStorage{})
}

func (s GoogleCloudStorage) Name() string {
	return STORAGE_NAME
}

func (s GoogleCloudStorage) Connect(ctx context.Context, opts ...testgrounds.Option) (testgrounds.Client, error) {
	options := testgrounds.Options(opts)

	mainOption, err := options.ByType(testgrounds.OPTION_TYPE_STORAGE).Merge()
	if err != nil {
		return nil, err
	}

	clientOptions := make([]option.ClientOption, 0)

	clientOptions = append(clientOptions, option.WithCredentials(&google.Credentials{JSON: mainOption.(*testgrounds.StorageOption).Credentials}))

	return NewGoogleCloudStorageClient(ctx, clientOptions...)
}
