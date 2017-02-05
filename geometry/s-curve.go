package geometry

import (
	"github.com/tobyjsullivan/dxf/drawing"
	"math"
)

type SCurve struct {
	Start             *Point
	End               *Point
	StartingAngle float64
	FinishAngle float64
	MaxAngle          float64
}

func (l *SCurve) DrawDXF(d *drawing.Drawing) error {
	mid := l.Start.MidpointTo(l.End)

	start := &EllipseCurve{
		Start: l.Start,
		End: mid,
		StartingAngle: l.StartingAngle,
		ArcAngle: l.MaxAngle,
	}

	end := &EllipseCurve{
		Start: l.End,
		End: mid,
		StartingAngle: l.FinishAngle - math.Pi,
		ArcAngle: l.MaxAngle,
	}

	err := start.DrawDXF(d)
	if err != nil {
		return err
	}

	err = end.DrawDXF(d)

	return err
}