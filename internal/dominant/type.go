package dominant

import (
	"image"
	"image/color"
)

type DominantColorFunc = func(image.Image) (color.RGBA, error)
