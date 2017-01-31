package geometry

import "github.com/tobyjsullivan/dxf/drawing"

type Drawable interface {
	DrawDXF(d *drawing.Drawing) error
}
