package google_cloud_storage

import (
	"context"
	"flavioltonon/testgrounds"
	"fmt"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type GoogleCloudStorageClient struct {
	context context.Context

	*storage.Client
}

func NewGoogleCloudStorageClient(ctx context.Context, opts ...option.ClientOption) (GoogleCloudStorageClient, error) {
	client, err := storage.NewClient(ctx, opts...)
	return GoogleCloudStorageClient{context: ctx, Client: client}, err
}

func (c GoogleCloudStorageClient) NewBucket(options ...testgrounds.Option) (testgrounds.Bucket, error) {
	option, err := testgrounds.Options(options).Merge()
	if err != nil {
		return GoogleCloudStorageBucket{}, err
	}

	if option.Type() != testgrounds.OPTION_TYPE_BUCKET {
		return GoogleCloudStorageBucket{}, fmt.Errorf("invalid bucket option type: %s", option.Type())
	}

	return GoogleCloudStorageBucket{
		context:      c.context,
		option:       option.(*testgrounds.BucketOption),
		BucketHandle: c.Bucket(option.(*testgrounds.BucketOption).Name),
	}, nil
}
