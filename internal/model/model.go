package model

import (
	"image"
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
