package sections

import (
	"math"
	"github.com/tailored-style/pattern-generator/geometry3d"
	"github.com/tobyjsullivan/dxf/drawing"
	"github.com/tobyjsullivan/dxf"
)

type Armhole struct {
	Circumference float64
	Centre *geometry3d.Point
}

func (a *Armhole) Radius() float64 {
	return a.Circumference / (2.0 * math.Pi)
}

func (a *Armhole) DrawDXF(d *drawing.Drawing) error {
	c, err := d.Circle(a.Centre.X, a.Centre.Y, a.Centre.Z, a.Radius())
	if err != nil {
		return err
	}
	dxf.SetExtrusion(c, []float64{-1.0, 0.0, 0.0})

	return nil
}