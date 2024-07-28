package util

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"

	"gocv.io/x/gocv"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

func ComposeColors(cellHeight, cellWidth int, rows, cols int, colors []color.RGBA) (image.Image, error) {
	img := image.NewRGBA(image.Rect(0, 0, cellWidth*cols, cellHeight*rows))

	for i, c := range colors {
		row := i / cols
		col := i % cols
		src := image.NewUniform(c)
		draw.Over.Draw(
			img,
			image.Rect(
				col*cellWidth, row*cellHeight,
				(col+1)*cellWidth, (row+1)*cellHeight,
			),
			src,
			image.Point{},
		)
	}

	f, err := os.Create("colors.out.jpg")
	if err != nil {
		return nil, err
	}

	err = jpeg.Encode(f, img, nil)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func MustColorMat(rows, cols int, clr color.Color) gocv.Mat {
	mat, err := gocv.ImageToMatRGB(ColorImage(rows, cols, clr))
	if err != nil {
		panic(err)
	}

	return mat
}

func ColorImage(rows, cols int, clr color.Color) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, cols, rows))
	for row := range rows {
		for col := range cols {
			img.Set(col, row, clr)
		}
	}

	return img
}

func LoadJPEG(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	img, err := jpeg.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}
