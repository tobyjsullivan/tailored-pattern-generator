package geometry

import (
	"github.com/tobyjsullivan/dxf/drawing"
)

type Line interface {
	DrawDXF(d *drawing.Drawing) error
}
