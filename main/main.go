package main

import (
	"flavioltonon/testgrounds"
	"fmt"
	"log"
)

func main() {
	var options = testgrounds.Options{
		testgrounds.StorageOptionName{
			Name: testgrounds.GOOGLE_CLOUD_STORAGE,
		},
		testgrounds.StorageOptionName{
			Name: testgrounds.ANOTHER_STORAGE,
		},
	}

	var optionsByType = make(map[string]testgrounds.Options, 0)

	for _, option := range options {
		optionsByType[option.Type()] = append(optionsByType[option.Type()], option)
	}

	storageOptions := optionsByType[testgrounds.OPTION_TYPE_STORAGE]

	storageOptionsName := storageOptions.BySubtype(testgrounds.OPTION_SUBTYPE_STORAGE_NAME)
	if len(storageOptionsName) > 1 {
		log.Fatal(fmt.Errorf("invalid number of storage options name: %d", len(storageOptionsName)))
	}

	storageName := storageOptionsName[0].(testgrounds.StorageOptionName).Name

	storage, err := testgrounds.RegisteredStorages().ByName(storageName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(storage.Name())

	// ctx := context.Background()
	// client, err := storage.NewClient(ctx)
	// if err != nil {
	// 	// TODO: Handle error.
	// }

	// fmt.Println(client, err)

}
