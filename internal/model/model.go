package model

import (
	"image"
	"image/color"
)

type Resolution struct {
	Width, Height int
}

type PathPoint struct {
	Pos  image.Point
	Edge Edge
}

type Edge int

const (
	LeftEdge  Edge = 0
	UpperEdge Edge = 1
	RightEdge Edge = 2
)

type SegmentInfo struct {
	Rect  image.Rectangle
	Color color.RGBA
}
