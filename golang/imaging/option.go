package imaging

import (
	"image"

	"github.com/disintegration/imaging"
)

type Option interface {
	apply(img image.Image) *image.NRGBA
}

type ResizeOption struct {
	width  int
	height int
	filter imaging.ResampleFilter
}

func Resize(w int, h int) Option {
	return &ResizeOption{
		width:  w,
		height: h,
		filter: imaging.Lanczos,
	}
}

func (o *ResizeOption) apply(img image.Image) *image.NRGBA {
	if img.Bounds().Dx() > img.Bounds().Dy() {
		o.height = 0
	} else {
		o.width = 0
	}

	return imaging.Resize(img, o.width, o.height, o.filter)
}

type GrayscaleOption struct{}

func Grayscale() Option {
	return &GrayscaleOption{}
}

func (o *GrayscaleOption) apply(img image.Image) *image.NRGBA {
	return imaging.Grayscale(img)
}

type AdjustContrastOption struct {
	percentage float64
}

func AdjustContrast(percentage float64) Option {
	return &AdjustContrastOption{
		percentage: percentage,
	}
}

func (o *AdjustContrastOption) apply(img image.Image) *image.NRGBA {
	return imaging.AdjustContrast(img, o.percentage)
}

type AdjustBrightnessOption struct {
	percentage float64
}

func AdjustBrightness(percentage float64) Option {
	return &AdjustBrightnessOption{
		percentage: percentage,
	}
}

func (o *AdjustBrightnessOption) apply(img image.Image) *image.NRGBA {
	return imaging.AdjustBrightness(img, o.percentage)
}

type SharpenOption struct {
	sigma float64
}

func Sharpen(sigma float64) Option {
	return &SharpenOption{
		sigma: sigma,
	}
}

func (o *SharpenOption) apply(img image.Image) *image.NRGBA {
	return imaging.Sharpen(img, o.sigma)
}
