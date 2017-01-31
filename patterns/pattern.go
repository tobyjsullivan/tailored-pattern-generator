package patterns

import (
	"github.com/yofu/dxf/drawing"
	"io"
)

type Pattern interface {
	PrintInstructions(io.Writer)
	DrawDXF(*drawing.Drawing) error
}
