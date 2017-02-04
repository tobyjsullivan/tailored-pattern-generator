package sections

import (
	"math"
	"github.com/tobyjsullivan/dxf/drawing"
)

type NeckHole struct {
	Circumference float64
	Height float64
}

func (n *NeckHole) DrawDXF(d *drawing.Drawing) error {
	_, err := d.Circle(0.0, 0.0, n.Height, n.Circumference / (2.0 * math.Pi))
	return err
}
