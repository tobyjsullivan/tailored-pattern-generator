package drawing

import (
	"github.com/tailored-style/pattern-generator/geometry"
)

const (
	LAYER_NORMAL = "Normal"
	LAYER_ALT1 = "Alt1"
	LAYER_ALT2 = "Alt2"
)

type Drawing interface {
	StraightLine(start *geometry.Point, end *geometry.Point) error
	Text(position *geometry.Point, content string, rotation *geometry.Angle) error
	SaveAs(filepath string) error
	DrawableWidth() float64
	SetLayer(layer string) error
}

