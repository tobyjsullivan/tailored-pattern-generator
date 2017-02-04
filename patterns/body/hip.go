package body

import (
	"math"
	"github.com/tobyjsullivan/dxf/drawing"
)

type Hip struct {
	Circumference float64
	Height float64
}

func (h *Hip) DrawDXF(d *drawing.Drawing) error {
	_, err := d.Circle(0.0, 0.0, h.Height, h.Circumference / (2.0 * math.Pi))
	return err
}
