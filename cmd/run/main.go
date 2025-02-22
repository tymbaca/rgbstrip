package main

import (
	"fmt"
	"log"

	"github.com/tymbaca/rgbstrip/internal/dominant/cenkalti"
	"github.com/tymbaca/rgbstrip/internal/model"
	"github.com/tymbaca/rgbstrip/internal/service/leds"
	"github.com/tymbaca/rgbstrip/internal/util"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	img, err := util.LoadJPEG("selfie.jpg")
	if err != nil {
		return err
	}

	svc := leds.Service{
		Screen: model.Resolution{
			Width:  img.Bounds().Dx(),
			Height: img.Bounds().Dy(),
		},
		SegCount:          120,
		SegOffset:         100,
		SegLength:         60,
		SegWidth:          80,
		DominantColorFunc: cenkalti.Find,
	}

	colors, err := svc.GetColors(img)
	if err != nil {
		return err
	}

	// err = util.ComposeColors(100, 100, 16, 10, colors)
	// if err != nil {
	// 	return err
	// }

	fmt.Println(colors)

	return nil
}
