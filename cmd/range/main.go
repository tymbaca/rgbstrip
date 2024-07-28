package main

import (
	"log"

	"github.com/tymbaca/rgbstrip/internal/dominant/cenkalti"
	"github.com/tymbaca/rgbstrip/internal/model"
	"github.com/tymbaca/rgbstrip/internal/service/leds"

	"github.com/tymbaca/rgbstrip/internal/util"
	"gocv.io/x/gocv"
)

const (
	_segCount  = 120
	_segOffset = 10
	_segLength = 120
	_segWidth  = 80
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("got webcam")

	window := gocv.NewWindow("Hello")
	// w2 := gocv.NewWindow("w2")
	imgMat := gocv.NewMat()
	svc := leds.Service{
		SegCount:          _segCount,
		SegOffset:         _segOffset,
		SegLength:         _segLength,
		SegWidth:          _segWidth,
		DominantColorFunc: cenkalti.Find,
	}

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

		svc.Screen = model.Resolution{
			Width:  origImg.Bounds().Dx(),
			Height: origImg.Bounds().Dy(),
		}

		//--------------------------------------------------------------------------------------------------

		segments, err := svc.GetColorsWithInfo(origImg)
		if err != nil {
			panic(err)
		}
		resultImg := util.DrawSegments(origImg, segments...)
		imgMat, err = gocv.ImageToMatRGB(resultImg)
		if err != nil {
			panic(err)
		}

		//--------------------------------------------------------------------------------------------------

		window.IMShow(imgMat)
		window.WaitKey(1)
	}
}
