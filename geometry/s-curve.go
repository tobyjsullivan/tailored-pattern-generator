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

	// Reverse array of lines for second half
	reversed := end.StraightLines()
	for i := 0; i < (len(reversed) / 2); i++ {
		tmp := reversed[i]
		tail := len(reversed) - (i + 1)
		reversed[i] = reversed[tail].Reverse()
		reversed[tail] = tmp.Reverse()
	}
	out = append(out, reversed...)

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