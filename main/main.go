package main

import (
	"context"
	"flavioltonon/testgrounds"
	gcs "flavioltonon/testgrounds/google_cloud_storage"
	"fmt"
	"log"
)

func main() {
	var options = testgrounds.Options{
		testgrounds.StorageOption{
			Name: gcs.STORAGE_NAME,
		},
		testgrounds.StorageOption{
			Credentials: []byte{},
		},
		testgrounds.BucketOption{
			Name:    "testgrounds",
			Project: "test1",
		},
	}

	storageOptions := options.ByType(testgrounds.OPTION_TYPE_STORAGE)

	storage, err := testgrounds.NewStorage(storageOptions[len(storageOptions)-1].(testgrounds.StorageOption).Name)
	if err != nil {
		log.Fatal("failed to set storage: ", err)
	}

	client, err := storage.Connect(context.Background(), nil)
	if err != nil {
		log.Fatal("failed to connect to storage: ", err)
	}

	bucket, err := client.NewBucket(options.ByType(testgrounds.OPTION_TYPE_BUCKET)...)
	if err != nil {
		log.Fatal("failed to create new bucket handler: ", err)
	}

	err = bucket.Create()
	if err != nil {
		log.Fatal("failed to create bucket: ", err)
	}

	attributes, err := bucket.Attributes()
	if err != nil {
		log.Fatal("failed to get bucket attributes: ", err)
	}

	fmt.Println(attributes)

	// ctx := context.Background()
	// client, err := storage.NewClient(ctx)
	// if err != nil {
	// 	// TODO: Handle error.
	// }

	// fmt.Println(client, err)

}
