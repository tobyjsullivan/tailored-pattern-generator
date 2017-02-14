package geometry

type Line interface {
	StraightLines() []*StraightLine
	BoundingBox() *BoundingBox
	Length() float64
	PointAt(dist float64) *Point
	AngleAt(dist float64) *Angle
}

func boundingBoxOfLine(l Line) *BoundingBox {
	ls := l.StraightLines()
	lines := make([]BoundedShape, 0, len(ls))
	for _, l := range ls {
		lines = append(lines, l)
	}
	return CollectiveBoundingBox(lines...)
}

func lengthOfLine(l Line) float64 {
	accrued := 0.0
	for _, cur := range l.StraightLines() {
		accrued += cur.Length()
	}

	return accrued
}

func pointOnLine(l Line, dist float64) *Point {
	accruedLen := 0.0

	for _, sl := range l.StraightLines() {
		curLen := sl.Length()
		if accruedLen + curLen >= dist {
			return sl.PointAt(dist - accruedLen)
		}

		accruedLen += curLen
	}

	panic("Cannot return a point that is beyond the end of the line.");
}

func angleAtPointOnLine(l Line, dist float64) *Angle {
	accruedLen := 0.0

	for _, sl := range l.StraightLines() {
		curLen := sl.Length()
		if accruedLen + curLen >= dist {
			return sl.AngleAt(0.0)
		}

		accruedLen += curLen
	}

	panic("Cannot return a point that is beyond the end of the line.");
}
