package body

import (
	"github.com/tobyjsullivan/dxf/drawing"
	"math"
)

type Chest struct {
	Circumference float64
	Height float64
}

func (c *Chest) DrawDXF(d *drawing.Drawing) error {
	_, err := d.Circle(0.0, 0.0, c.Height, c.Circumference / (2.0 * math.Pi))
	return err
}
