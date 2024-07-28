package leds

import (
	"image"
	"image/color"
	"log"

	"github.com/anthonynsimon/bild/transform"
	"github.com/tymbaca/rgbstrip/internal/model"
	"github.com/tymbaca/rgbstrip/internal/util/subimage"
)

func (s *Service) GetColorsWithInfo(img image.Image) ([]model.SegmentInfo, error) {
	points, segRects, _, dominants, err := s.calculateAll(img)
	if err != nil {
		return nil, err
	}

	segments := make([]model.SegmentInfo, len(points))
	for i := range points {
		segments[i] = model.SegmentInfo{
			Rect:  segRects[i],
			Color: dominants[i],
		}
	}

	return segments, nil
}

func (s *Service) GetColors(img image.Image) ([]color.RGBA, error) {
	_, _, _, dominants, err := s.calculateAll(img)
	if err != nil {
		return nil, err
	}

	return dominants, nil
}

func (s *Service) calculateAll(img image.Image) ([]model.PathPoint, []image.Rectangle, []image.Image, []color.RGBA, error) {
	points := s.getPoints()
	segRects := s.getSegmentsRects(points)
	segSrcs := s.getSegments(img, segRects)

	for i := range segSrcs {
		segSrcs[i] = resizeImage(segSrcs[i], 0.1)
	}

	dominants, err := s.getDominants(segSrcs)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return points, segRects, segSrcs, dominants, nil
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

func (s *Service) getDominants(segments []image.Image) ([]color.RGBA, error) {
	dominants := make([]color.RGBA, len(segments))
	var err error

	for i, seg := range segments {
		dominants[i], err = s.DominantColorFunc(seg)
		if err != nil {
			return nil, err
		}
	}

	return dominants, nil
}

// factor must be positive number
func resizeImage(img image.Image, factor float32) image.Image {
	width := int(float32(img.Bounds().Dx()) * factor)
	height := int(float32(img.Bounds().Dy()) * factor)
	return transform.Resize(img, width, height, transform.NearestNeighbor)
}
