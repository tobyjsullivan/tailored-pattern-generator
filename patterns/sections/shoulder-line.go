package sections

import (
	"github.com/tailored-style/pattern-generator/geometry3d"
	"github.com/tobyjsullivan/dxf/drawing"
)

type ShoulderLine struct {
	NeckHolePoint *geometry3d.Point
	ArmholePoint *geometry3d.Point
}

func (sl *ShoulderLine) Length() float64 {
	return sl.ArmholePoint.EuclideanDistanceTo(sl.NeckHolePoint)
}

func (sl *ShoulderLine) DrawDXF(d *drawing.Drawing) error {
	_, err := d.Line(sl.NeckHolePoint.X, sl.NeckHolePoint.Y, sl.NeckHolePoint.Z, sl.ArmholePoint.X, sl.ArmholePoint.Y, sl.ArmholePoint.Z)

	return err
}