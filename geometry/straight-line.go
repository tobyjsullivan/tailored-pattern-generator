package geometry

import (
	"math"
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
	run := (l.End.X - l.Start.X)

	angle := math.Atan((l.End.Y - l.Start.Y) / run)

	if run < 0.0 {
		angle += math.Pi
	}

	return &Angle{
		Rads: angle,
	}
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

func (l *StraightLine) PointAt(dist float64) *Point {
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