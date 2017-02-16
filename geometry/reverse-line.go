package geometry

type ReverseLine struct {
	InnerLine Line
}

func (l *ReverseLine) StraightLines() []*StraightLine {
	ls := l.InnerLine.StraightLines()

	for i := 0; i < (len(ls) / 2); i++ {
		tail := len(ls) - (i + 1)

		tmp := ls[tail]
		ls[tail] = ls[i].Reverse()
		ls[i] = tmp.Reverse()
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