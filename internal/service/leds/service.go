package leds

import (
	"github.com/tymbaca/rgbstrip/internal/dominant"
	"github.com/tymbaca/rgbstrip/internal/model"
)

// |
// | <- segOffset -> |------- segLength -------|
// |
type Service struct {
	Screen            model.Resolution
	SegCount          int
	SegOffset         int
	SegLength         int
	SegWidth          int
	DominantColorFunc dominant.DominantColorFunc
}
