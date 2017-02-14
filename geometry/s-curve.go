package geometry

type SCurve struct {
	Start         *Point
	End           *Point
	StartingAngle *Angle
	FinishAngle   *Angle
	MaxAngle      *Angle
}

func (l *SCurve) StraightLines() []*StraightLine {
	mid := l.Start.MidpointTo(l.End)

	start := &EllipseCurve{
		Start:         l.Start,
		End:           mid,
		StartingAngle: l.StartingAngle,
		ArcAngle:      l.MaxAngle,
	}

	end := &EllipseCurve{
		Start:         l.End,
		End:           mid,
		StartingAngle: l.FinishAngle.Opposite(),
		ArcAngle:      l.MaxAngle,
	}

	out := []*StraightLine{}
	out = append(out, start.StraightLines()...)
	out = append(out, end.StraightLines()...)

	return out
}

func (c *SCurve) BoundingBox() *BoundingBox {
	return boundingBoxOfLine(c)
}

func (c *SCurve) Length() float64 {
	return lengthOfLine(c)
}

func (c *SCurve) PointAt(dist float64) *Point {
	return pointOnLine(c, dist)
}

func (c *SCurve) AngleAt(dist float64) *Angle {
	return angleAtPointOnLine(c, dist)
}