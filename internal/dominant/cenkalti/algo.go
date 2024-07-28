package cenkalti

import (
	"image"
	"image/color"

	"github.com/cenkalti/dominantcolor"
)

func Find(img image.Image) (color.RGBA, error) {
	return dominantcolor.Find(img), nil
}
