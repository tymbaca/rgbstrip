package leds

import (
	"testing"

	"github.com/tymbaca/rgbstrip/internal/dominant/cenkalti"
	"github.com/tymbaca/rgbstrip/internal/model"
	"github.com/tymbaca/rgbstrip/internal/util"
)

const (
	_segCount  = 60
	_segOffset = 10
	_segLength = 60
	_segWidth  = 80
)

// v0
// BenchmarkGetColors/full-8         	      34	  34371533 ns/op	 6945179 B/op	 1955915 allocs/op
// BenchmarkGetColors/get_points-8   	 4604133	       261.9 ns/op	    3072 B/op	       1 allocs/op
// BenchmarkGetColors/get_rects-8    	 2299638	       520.4 ns/op	    4096 B/op	       1 allocs/op
// BenchmarkGetColors/get_segs-8     	  310189	      3817 ns/op	   17408 B/op	     121 allocs/op
// BenchmarkGetColors/get_dominants-8         	      33	  34803347 ns/op	 6921222 B/op	 1955796 allocs/op

// v0 (хз что произошло)
// BenchmarkGetColors/full-8         	      48	  25716189 ns/op	 5336138 B/op	 1453877 allocs/op

// v0 (сравнил два разных алгоритма)
// BenchmarkGetColors/cenkalti-8         	      31	  36939751 ns/op	12823987 B/op	 1984273 allocs/op
// BenchmarkGetColors/nykakin-8          	       2	 990887896 ns/op	810642268 B/op	38866795 allocs/op

// v1 - shrink x0.1 eash segment (120 segments)
// BenchmarkGetColors/jpg-8         	     288	   4127682 ns/op	 3267593 B/op	   16741 allocs/op
// BenchmarkGetColors/rgba-8        	     385	   3069641 ns/op	 1330276 B/op	   16209 allocs/op

// v1 - shrink x0.1 eash segment (60 segments), resize: "github.com/nfnt/resize" (ClosestNeighbor)
// BenchmarkGetColors/jpg-8         	     573	   2050433 ns/op	 1631915 B/op	    8608 allocs/op
// BenchmarkGetColorsWithInfo/jpg-8          580	   2023037 ns/op	 1634434 B/op	    8609 allocs/op
// BenchmarkGetColors/rgba-8        	     721	   1737113 ns/op	  665841 B/op	    8171 allocs/op
// BenchmarkGetColorsWithInfo/rgba-8         718	   1596750 ns/op	  668575 B/op	    8172 allocs/op

// v2 - resize: "github.com/anthonynsimon/bild/transform" (ClosestNeighbor)
// BenchmarkGetColors/jpg-8         	     926	   1260019 ns/op	 1593614 B/op	    4953 allocs/op
// BenchmarkGetColorsWithInfo/jpg-8          928	   1264194 ns/op	 1596288 B/op	    4954 allocs/op
// BenchmarkGetColors/rgba-8        	    1826	    595702 ns/op	  377386 B/op	    4833 allocs/op
// BenchmarkGetColorsWithInfo/rgba-8        1852	    595068 ns/op	  379989 B/op	    4834 allocs/op

func BenchmarkGetColors(b *testing.B) {
	b.Run("jpg_GetColors", func(b *testing.B) {
		img, err := util.LoadJPEG("testdata/selfie.jpg")
		if err != nil {
			b.Fatal(err)
		}

		svc := Service{
			Screen: model.Resolution{
				Width:  img.Bounds().Dx(),
				Height: img.Bounds().Dy(),
			},
			SegCount:          _segCount,
			SegOffset:         _segOffset,
			SegLength:         _segLength,
			SegWidth:          _segWidth,
			DominantColorFunc: cenkalti.Find,
		}

		for range b.N {
			_, err := svc.GetColors(img)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("jpg_WithInfo", func(b *testing.B) {
		img, err := util.LoadJPEG("testdata/selfie.jpg")
		if err != nil {
			b.Fatal(err)
		}

		svc := Service{
			Screen: model.Resolution{
				Width:  img.Bounds().Dx(),
				Height: img.Bounds().Dy(),
			},
			SegCount:          _segCount,
			SegOffset:         _segOffset,
			SegLength:         _segLength,
			SegWidth:          _segWidth,
			DominantColorFunc: cenkalti.Find,
		}

		for range b.N {
			_, err := svc.GetColorsWithInfo(img)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("rgba_GetColors", func(b *testing.B) {
		img, err := util.LoadJPEG("testdata/selfie.jpg")
		if err != nil {
			b.Fatal(err)
		}

		img = util.ImageToRGBA(img)

		svc := Service{
			Screen: model.Resolution{
				Width:  img.Bounds().Dx(),
				Height: img.Bounds().Dy(),
			},
			SegCount:          _segCount,
			SegOffset:         _segOffset,
			SegLength:         _segLength,
			SegWidth:          _segWidth,
			DominantColorFunc: cenkalti.Find,
		}

		for range b.N {
			_, err := svc.GetColors(img)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("rgba_WithInfo", func(b *testing.B) {
		img, err := util.LoadJPEG("testdata/selfie.jpg")
		if err != nil {
			b.Fatal(err)
		}

		img = util.ImageToRGBA(img)

		svc := Service{
			Screen: model.Resolution{
				Width:  img.Bounds().Dx(),
				Height: img.Bounds().Dy(),
			},
			SegCount:          _segCount,
			SegOffset:         _segOffset,
			SegLength:         _segLength,
			SegWidth:          _segWidth,
			DominantColorFunc: cenkalti.Find,
		}

		for range b.N {
			_, err := svc.GetColorsWithInfo(img)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkGetColorsWithInfo(b *testing.B) {
}

func BenchmarkImageFormats(b *testing.B) {
	b.Run("JPEG", func(b *testing.B) {
		img, err := util.LoadJPEG("testdata/selfie.jpg")
		if err != nil {
			b.Fatal(err)
		}

		for range b.N {
			for y := range img.Bounds().Dy() {
				for x := range img.Bounds().Dx() {
					c := img.At(x, y)
					_ = c
				}
			}
		}
	})

	b.Run("RGBA", func(b *testing.B) {
		img, err := util.LoadJPEG("testdata/selfie.jpg")
		if err != nil {
			b.Fatal(err)
		}

		img = util.ImageToRGBA(img)
		for range b.N {
			for y := range img.Bounds().Dy() {
				for x := range img.Bounds().Dx() {
					c := img.At(x, y)
					_ = c
				}
			}
		}
	})
}
