package geometry

import (
	"fmt"
	"github.com/yofu/dxf/drawing"
)

type StraightLine struct {
	Start *Point
	End   *Point
}

func (l *StraightLine) GetStart() *Point {
	return l.Start
}

func (l *StraightLine) GetEnd() *Point {
	return l.End
}

func (l *StraightLine) ToEnglish() string {
	return fmt.Sprintf("Straight line from %v to %v\n", l.Start, l.End)
}

func (l *StraightLine) ToAutoCAD() string {
	return fmt.Sprintf("(command \"LINE\" \"%.1f,%.1f\" \"%.1f,%.1f\" \"\")\n", l.Start.X, l.Start.Y, l.End.X, l.End.Y)
}

func (l *StraightLine) DrawDXF(d *drawing.Drawing) error {
	_, err := d.Line(l.Start.X, l.Start.Y, 0.0, l.End.X, l.End.Y, 0.0)
	//var err error = nil

	return err
}
