package model

import (
	"image"
	"image/color"
)

type DominantColorFunc func(image.Image) color.RGBA

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
