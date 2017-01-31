package geometry

import (
	"github.com/yofu/dxf/drawing"
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
