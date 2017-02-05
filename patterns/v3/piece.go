package v3

import "github.com/tailored-style/pattern-generator/geometry"

type Piece interface {
	CutLines() []geometry.Line
	StitchLines() []geometry.Line
	FoldLines() []geometry.Line
	GrainLine() *geometry.Line
	Notations() []geometry.Drawable
	Dimensions() *geometry.Point
	BoundingBox() *BoundingBox
}

type BoundingBox struct {
	TopLeft *geometry.Point
	BottomRight *geometry.Point
}
