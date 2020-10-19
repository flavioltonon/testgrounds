package hash

import (
	"testing"

	hash "testgrounds/structures/hash-table"

	"github.com/stretchr/testify/assert"
)

func TestHashTable(t *testing.T) {
	input := []string{
		"Brazil", "Czech Republic", "Canada", "United States of America", "Italy", "Austria", "France",
		"Netherlands", "United Kingdom", "Russia", "Portugal", "Chile", "Argentina", "Paraguay", "Peru",
		"South Africa", "Japan", "China", "Australia", "New Zealand", "Spain", "Mexico", "Madagascar",
		"South Korea", "Switzerland", "Germany", "Poland", "Iceland", "Norway", "Sweden", "Belgium",
	}

	table := hash.NewTable()

	for _, v := range input {
		table.Add(v)
	}

	assert.Equal(t, len(input), table.Size())

	for _, elem := range input {
		assert.Equal(t, true, table.Search(elem))
	}
}
