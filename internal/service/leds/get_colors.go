package leds

import (
	"image"
	"image/color"
	"log"

	"github.com/tymbaca/rgbstrip/internal/model"
	"github.com/tymbaca/rgbstrip/internal/util/subimage"
)

func (s *Service) GetColors(img image.Image) []color.RGBA {
	points := s.getPoints()
	segRects := s.getSegmentsRects(points)
	segments := s.getSegments(img, segRects)
	dominants := s.getDominants(segments)

	return dominants
}

func (s *Service) getPoints() []model.PathPoint {
	width := s.Screen.Width - s.SegOffset*2
	height := s.Screen.Height - s.SegOffset*2
	pathLen := width + height*2 // get left, upper and right sides

	leftEdgeEnd := float32(height)
	upperEdgeEnd := float32(height + width)
	// rightEdgeEnd := float32(height + width + height)

	points := make([]model.PathPoint, s.SegCount)
	step := float32(pathLen) / float32(s.SegCount)
	cur := float32(0)
	for i := range s.SegCount {
		var edge model.Edge
		var x int
		var y int
		if cur <= leftEdgeEnd { // left edge
			x = s.SegOffset
			y = s.Screen.Height - s.SegOffset - int(cur)
			edge = model.LeftEdge
		} else if cur <= upperEdgeEnd { // upper edge
			x = s.SegOffset + int(cur-leftEdgeEnd)
			y = s.SegOffset
			edge = model.UpperEdge
		} else { // right edge
			x = s.Screen.Width - s.SegOffset
			y = s.SegOffset + int(cur-upperEdgeEnd)
			edge = model.RightEdge
		}

		points[i] = model.PathPoint{
			Pos:  image.Point{X: x, Y: y},
			Edge: edge,
		}

		cur += step
	}

	return points
}

func (s *Service) getSegmentsRects(points []model.PathPoint) []image.Rectangle {
	rects := make([]image.Rectangle, len(points))
	for i, p := range points {
		rects[i] = rect(p.Pos, s.SegLength, s.SegWidth, p.Edge)
	}
	return rects
}

func rect(origin image.Point, length, width int, edge model.Edge) image.Rectangle {
	switch edge {
	case model.LeftEdge:
		return image.Rect(
			origin.X, origin.Y-width/2,
			origin.X+length, origin.Y+width/2,
		)
	case model.UpperEdge:
		return image.Rect(
			origin.X-width/2, origin.Y,
			origin.X+width/2, origin.Y+length,
		)
	case model.RightEdge:
		return image.Rect(
			origin.X-length, origin.Y-width/2,
			origin.X, origin.Y+width/2,
		)
	}

	log.Fatalf("got incorrect edge: %+v", edge)
	return image.Rectangle{}
}

func (s *Service) getSegments(img image.Image, rects []image.Rectangle) []image.Image {
	segments := make([]image.Image, len(rects))
	for i, r := range rects {
		segments[i] = subimage.MustGet(img, r)
	}
	return segments
}

func (s *Service) getDominants(segments []image.Image) []color.RGBA {
	dominants := make([]color.RGBA, len(segments))

	for i, seg := range segments {
		dominants[i] = s.DominantColorFunc(seg)
	}

	return dominants
}
