package space

import "github.com/tailored-style/pattern-generator/geometry"

type RootAnchor struct {}

func (a *RootAnchor) GetPosition() {
	return &geometry.Point{X: 0.0, Y: 0.0}
}
