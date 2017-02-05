package geometry

import (
	"github.com/tobyjsullivan/dxf/drawing"
	"math"
)

type StraightLine struct {
	Start *Point
	End   *Point
}

func (l *StraightLine) DrawDXF(d *drawing.Drawing) error {
	_, err := d.Line(l.Start.X, l.Start.Y, 0.0, l.End.X, l.End.Y, 0.0)
	//var err error = nil

	return err
}

func (l *StraightLine) Angle() float64 {
	run := (l.End.X - l.Start.X)

	angle := math.Atan((l.End.Y - l.Start.Y)/run)

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
		End: l.Start.DrawAt(l.Angle(), length),
	}
}