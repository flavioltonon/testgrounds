package imaging_test

import (
	"fmt"
	"testgrounds/imaging"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	path := fmt.Sprintf("../assets/images/sample.jpg")

	err := imaging.Process(path,
		imaging.Resize(1024, 768),
		imaging.Grayscale(),
		imaging.AdjustContrast(75),
		imaging.AdjustBrightness(-15),
		imaging.Sharpen(3),
	)

	assert.NoError(t, err)
}
