package geometry

import (
	"math"
)

type StraightLine struct {
	Start *Point
	End   *Point
}

func (l *StraightLine) StraightLines() []*StraightLine {
	return []*StraightLine{l}
}

func (l *StraightLine) Angle() float64 {
	run := (l.End.X - l.Start.X)

	angle := math.Atan((l.End.Y - l.Start.Y) / run)

	if run < 0.0 {
		angle += math.Pi
	}

	return angle
}

func (l *StraightLine) PerpendicularAngle() float64 {
	return l.Angle() - (math.Pi / 2.0)
}

func (l *StraightLine) Resize(length float64) *StraightLine {
	return &StraightLine{
		Start: l.Start,
		End:   l.Start.DrawAt(l.Angle(), length),
	}
}

func (l *StraightLine) Move(x, y float64) *StraightLine {
	return &StraightLine{
		Start: l.Start.Move(x, y),
		End:   l.End.Move(x, y),
	}
}

func (l *StraightLine) BoundingBox() *BoundingBox {
	return CollectiveBoundingBox(l.Start, l.End)
}