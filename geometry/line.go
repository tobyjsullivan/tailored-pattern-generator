package geometry

type Line interface {
	StraightLines() []*StraightLine
	BoundingBox() *BoundingBox
}
