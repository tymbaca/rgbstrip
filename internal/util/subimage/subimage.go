package subimage

import (
	"fmt"
	"image"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func Get(img image.Image, rect image.Rectangle) (image.Image, error) {
	si, ok := img.(SubImager)
	if !ok {
		return nil, fmt.Errorf("can't get subimage: image doesn't implement SubImager: %#v", img)
	}

	return si.SubImage(rect), nil
}

func MustGet(img image.Image, rect image.Rectangle) image.Image {
	sub, err := Get(img, rect)
	if err != nil {
		panic(err)
	}

	return sub
}
