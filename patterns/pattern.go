package patterns

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"io"
)

type Pattern interface {
	GetPoints() map[string]geometry.Point
	GetLines() []geometry.Line
	PrintInstructions(io.Writer)
}
