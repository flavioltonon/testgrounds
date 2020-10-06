package structures_test

import (
	"testing"

	"testgrounds/structures"

	"github.com/stretchr/testify/assert"
)

func TestHashTable(t *testing.T) {
	input := []string{
		"Brazil", "Czech Republic", "Canada", "United States of America", "Italy", "Austria", "France",
		"Netherlands", "United Kingdom", "Russia", "Portugal", "Chile", "Argentina", "Paraguay", "Peru",
		"South Africa", "Japan", "China", "Australia", "New Zealand", "Spain", "Mexico", "Madagascar",
		"South Korea", "Switzerland", "Germany", "Poland", "Iceland", "Norway", "Sweden", "Belgium",
	}

	hashTable := structures.NewHashTable()

	for _, v := range input {
		hashTable.Insert(v)
	}

	assert.Equal(t, len(input), hashTable.Size())

	for _, elem := range input {
		assert.Equal(t, true, hashTable.Search(elem))
	}

	// for i, bucket := range hashTable.Buckets {
	// 	if bucket == nil {
	// 		continue
	// 	}

	// 	fmt.Printf("Bucket: %d | Len: %d | Elements: %v\n", i, bucket.Len(), bucket.Elements())
	// }
}
