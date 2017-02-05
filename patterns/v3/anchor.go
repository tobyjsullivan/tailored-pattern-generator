package v3

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tobyjsullivan/dxf/drawing"
)

type anchor struct {
	*geometry.Point
	label string
}

func (a *anchor) DrawDXF(d *drawing.Drawing) error {
	return a.Point.DrawDXF(a.label, d)
}
