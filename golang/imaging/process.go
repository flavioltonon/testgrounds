package imaging

import (
	"fmt"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
)

func Process(path string, opts ...Option) error {
	src, err := imaging.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open image: %v", err)
	}

	for _, opt := range opts {
		src = opt.apply(src)
	}

	out := fmt.Sprintf("../output/imaging/%s.jpg", uuid.New())

	err = imaging.Save(src, out)
	if err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	fmt.Printf("Output image successfully processed and saved at %s\n", out)

	return nil
}
