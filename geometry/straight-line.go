package geometry

import (
	"fmt"
)

type StraightLine struct {
	Start *Point
	End   *Point
}

func (l *StraightLine) StraightLines() []*StraightLine {
	return []*StraightLine{l}
}

func (l *StraightLine) AngleAt(_ float64) *Angle {
	return l.End.AngleRelativeTo(l.Start)
}

func (l *StraightLine) Resize(length float64) *StraightLine {
	return &StraightLine{
		Start: l.Start,
		End:   l.Start.DrawAt(l.AngleAt(0.0), length),
	}
}

func (l *StraightLine) String() string {
	return fmt.Sprintf("[%v, %v]", l.Start, l.End)
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

func (l *StraightLine) Length() float64 {
	return l.Start.DistanceTo(l.End)
}

func (l *StraightLine) TangentAt(dist float64) *Tangent {
	a := l.AngleAt(dist)

	return &Tangent{
		Origin: l.Start.DrawAt(a, dist),
		Direction: a,
	}
}

func (l *StraightLine) PointAt(dist float64) *Point {
	if l.Length() == 0.0 {
		return l.Start
	}

	ratio := dist / l.Length()

	return &Point{
		X: l.Start.X + ratio * (l.End.X - l.Start.X),
		Y: l.Start.Y + ratio * (l.End.Y - l.Start.Y),
	}
}

func (l *StraightLine) Reverse() *StraightLine {
	return &StraightLine{
		Start: l.End,
		End: l.Start,
	}
}

func (l *StraightLine) RotateAround(o *Point, a *Angle) *StraightLine {
	return &StraightLine{
		Start: l.Start.RotateAround(o, a),
		End: l.End.RotateAround(o, a),
	}
}

func (l *StraightLine) Equals(o *StraightLine) bool {
	return l.Start.Equals(o.Start) && l.End.Equals(o.End)
}
