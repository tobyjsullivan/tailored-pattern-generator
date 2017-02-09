package geometry

import (
	"math"
)

type SCurve struct {
	Start         *Point
	End           *Point
	StartingAngle float64
	FinishAngle   float64
	MaxAngle      float64
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
		StartingAngle: l.FinishAngle - math.Pi,
		ArcAngle:      l.MaxAngle,
	}

	out := []*StraightLine{}
	out = append(out, start.StraightLines()...)
	out = append(out, end.StraightLines()...)

	return out
}

func (c *SCurve) BoundingBox() *BoundingBox {
	ls := c.StraightLines()
	lines := make([]BoundedShape, 0, len(ls))
	for _, l := range ls {
		lines = append(lines, l)
	}
	return CollectiveBoundingBox(lines...)
}
