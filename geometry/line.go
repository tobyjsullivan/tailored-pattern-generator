package geometry

import (
    "github.com/yofu/dxf/drawing"
)

type Line interface {
    GetStart() *Point
    GetEnd() *Point

    ToEnglish() string
    ToAutoCAD() string
    DrawDXF(d *drawing.Drawing) error
}

