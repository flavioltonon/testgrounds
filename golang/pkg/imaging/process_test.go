package imaging

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	path := fmt.Sprintf("../assets/images/sample.jpg")

	err := Process(path,
		Resize(1024, 768),
		Grayscale(),
		AdjustContrast(75),
		AdjustBrightness(-15),
		Sharpen(3),
	)

	assert.NoError(t, err)
}
