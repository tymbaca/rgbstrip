package nykakin

import (
	"image"
	"image/color"

	"github.com/Nykakin/quantize"
)

type algo struct {
	q q
}

type q interface {
	Quantize(img image.Image, count int) ([]color.RGBA, error)
}

func New() *algo {
	q := quantize.NewHierarhicalQuantizer()
	return &algo{q: q}
}

func (a *algo) Find(img image.Image) (color.RGBA, error) {
	colors, err := a.q.Quantize(img, 1)
	if err != nil {
		return color.RGBA{}, err
	}

	if len(colors) != 1 {
		panic("not 1 color from algo")
	}

	return colors[0], nil
}
