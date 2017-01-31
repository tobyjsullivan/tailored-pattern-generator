package patterns

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tobyjsullivan/dxf/drawing"
	"math"
)

const (
	TEXT_SIZE_LARGE = 1.3
	TEXT_SIZE_NORMAL = 1.0
	TEXT_SIZE_SMALL = 0.5
)

type Label struct {
	content  string
	location *geometry.Point
	Rotation float64
	Size 	 float64
}

func NewLabel(content string, location *geometry.Point) *Label {
	return &Label{
		content: content,
		location: location,
		Rotation: 0.0,
		Size: TEXT_SIZE_NORMAL,
	}
}

func (l *Label) DrawDXF(d *drawing.Drawing) error {
	t, err := d.Text(l.content, l.location.X, l.location.Y, 0.0, 1.0)
	degrees := 360 * (l.Rotation / (2 * math.Pi))
	t.Rotate = degrees
	t.Height = l.Size

	return err
}