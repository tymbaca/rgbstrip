package main

import (
	"log"

	"github.com/tymbaca/rgbstrip/internal/dominant/cenkalti"
	"github.com/tymbaca/rgbstrip/internal/model"
	"github.com/tymbaca/rgbstrip/internal/service/leds"
	. "github.com/tymbaca/rgbstrip/internal/util"
	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("got webcam")

	window := gocv.NewWindow("Hello")
	w2 := gocv.NewWindow("w2")
	imgMat := gocv.NewMat()

	for {
		// order of colors is BGR, not RGB
		webcam.Read(&imgMat)
		rows, cols, matType := imgMat.Rows(), imgMat.Cols(), imgMat.Type()
		origImg, err := imgMat.ToImage()
		if err != nil {
			panic(err)
		}

		_, _, _ = rows, cols, matType
		_ = origImg

		//--------------------------------------------------------------------------------------------------

		svc := leds.Service{
			Screen: model.Resolution{
				Width:  origImg.Bounds().Dx(),
				Height: origImg.Bounds().Dy(),
			},
			SegCount:          120,
			SegOffset:         100,
			SegLength:         60,
			SegWidth:          80,
			DominantColorFunc: cenkalti.Find,
		}

		colors, err := svc.GetColors(origImg)
		if err != nil {
			panic(err)
		}
		debugImg := Must(ComposeColors(50, 30, 1, 120, colors))

		debugMat, err := gocv.ImageToMatRGB(debugImg)
		if err != nil {
			panic(err)
		}

		// for row := range rows {
		// 	for col := range cols {
		// 		r, g, b, a := img.At(col, row).RGBA()
		// 		_, _, _, _ = r, g, b, a
		// 		// log.Printf("RGBA: %d, %d, %d, %d\n", r, g, b, a)
		// 	}
		// }

		// break

		//--------------------------------------------------------------------------------------------------

		// mat, err = gocv.NewMatFromBytes(rows, cols, matType, buf.Bytes())
		// if err != nil {
		// 	panic(err)
		// }

		w2.IMShow(debugMat)

		window.IMShow(imgMat)
		window.WaitKey(1)
	}
}
