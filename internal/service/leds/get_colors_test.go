package leds

import (
	"testing"

	nykakin "github.com/tymbaca/rgbstrip/internal/dominant/Nykakin"
	"github.com/tymbaca/rgbstrip/internal/dominant/cenkalti"
	"github.com/tymbaca/rgbstrip/internal/model"
	"github.com/tymbaca/rgbstrip/internal/util"
)

// v0
// BenchmarkGetColors/full-8         	      34	  34371533 ns/op	 6945179 B/op	 1955915 allocs/op
// BenchmarkGetColors/get_points-8   	 4604133	       261.9 ns/op	    3072 B/op	       1 allocs/op
// BenchmarkGetColors/get_rects-8    	 2299638	       520.4 ns/op	    4096 B/op	       1 allocs/op
// BenchmarkGetColors/get_segs-8     	  310189	      3817 ns/op	   17408 B/op	     121 allocs/op
// BenchmarkGetColors/get_dominants-8         	      33	  34803347 ns/op	 6921222 B/op	 1955796 allocs/o

// v0 (хз что произошло)
// BenchmarkGetColors/full-8         	      48	  25716189 ns/op	 5336138 B/op	 1453877 allocs/o

// v0 (сравнил два разных алгоритма)
// BenchmarkGetColors/cenkalti-8         	      31	  36939751 ns/op	12823987 B/op	 1984273 allocs/op
// BenchmarkGetColors/nykakin-8          	       2	 990887896 ns/op	810642268 B/op	38866795 allocs/o

func BenchmarkGetColors(b *testing.B) {
	img, err := util.LoadJPEG("testdata/selfie.jpg")
	if err != nil {
		b.Fatal(err)
	}

	b.Run("cenkalti", func(b *testing.B) {
		svc := Service{
			Screen: model.Resolution{
				Width:  img.Bounds().Dx(),
				Height: img.Bounds().Dy(),
			},
			SegCount:          120,
			SegOffset:         10,
			SegLength:         60,
			SegWidth:          80,
			DominantColorFunc: cenkalti.Find,
		}

		for range b.N {
			_, err := svc.GetColors(img)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("nykakin", func(b *testing.B) {
		svc := Service{
			Screen: model.Resolution{
				Width:  img.Bounds().Dx(),
				Height: img.Bounds().Dy(),
			},
			SegCount:          120,
			SegOffset:         10,
			SegLength:         60,
			SegWidth:          80,
			DominantColorFunc: nykakin.New().Find,
		}

		for range b.N {
			_, err := svc.GetColors(img)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	// b.Run("get points", func(b *testing.B) {
	// 	for range b.N {
	// 		svc.getPoints()
	// 	}
	// })
	// b.Run("get rects", func(b *testing.B) {
	// 	points := svc.getPoints()
	// 	for range b.N {
	// 		svc.getSegmentsRects(points)
	// 	}
	// })
	// b.Run("get segs", func(b *testing.B) {
	// 	points := svc.getPoints()
	// 	segRects := svc.getSegmentsRects(points)
	// 	for range b.N {
	// 		svc.getSegments(img, segRects)
	// 	}
	// })
	// b.Run("get dominants", func(b *testing.B) {
	// 	points := svc.getPoints()
	// 	segRects := svc.getSegmentsRects(points)
	// 	segments := svc.getSegments(img, segRects)
	// 	for range b.N {
	// 		svc.getDominants(segments)
	// 	}
	// })
}
