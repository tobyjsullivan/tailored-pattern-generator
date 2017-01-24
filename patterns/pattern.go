package patterns

import (
    "github.com/tailored-style/pattern-generator/geometry"
    "io"
)

type Pattern interface {
    GetPoints() map[string]geometry.Point
    PrintInstructions(io.Writer)
}
