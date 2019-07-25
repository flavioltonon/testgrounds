package google_cloud_storage

import (
	"context"
	"flavioltonon/testgrounds"

	"cloud.google.com/go/storage"
)

type GoogleCloudStorageBucket struct {
	context context.Context
	option  *testgrounds.BucketOption

	*storage.BucketHandle
}

func (b GoogleCloudStorageBucket) Create() error {
	return b.BucketHandle.Create(b.context, b.option.Project, &storage.BucketAttrs{})
}

func (b GoogleCloudStorageBucket) Attributes() (testgrounds.BucketAttributes, error) {
	attributes, err := b.Attrs(b.context)
	return GoogleCloudStorageBucketAttributes{attributes}, err
}
