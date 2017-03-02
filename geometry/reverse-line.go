package geometry

type ReverseLine struct {
	InnerLine Line
}

func (l *ReverseLine) StraightLines() []*StraightLine {
	originalLines := l.InnerLine.StraightLines()
	ls := make([]*StraightLine, len(originalLines))

	copy(ls, originalLines)

	for i := 0; i < (len(ls) / 2); i++ {
		tailIdx := len(ls) - (i + 1)

		ls[tailIdx], ls[i] = ls[i].Reverse(), ls[tailIdx].Reverse()
	}

	if len(ls) % 2 == 1 {
		midx := len(ls) / 2
		ls[midx] = ls[midx].Reverse()
	}

	return ls
}

func (l *ReverseLine) BoundingBox() *BoundingBox {
	return boundingBoxOfLine(l)
}

func (l *ReverseLine) Length() float64 {
	return lengthOfLine(l)
}

func (l *ReverseLine) PointAt(dist float64) *Point {
	return pointOnLine(l, dist)
}

func (l *ReverseLine) AngleAt(dist float64) *Angle {
	return angleAtPointOnLine(l, dist)
}