package geometry

import (
	"github.com/yofu/dxf/drawing"
)

type Line interface {
	DrawDXF(d *drawing.Drawing) error
}
