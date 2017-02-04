package space

import "github.com/tailored-style/pattern-generator/geometry"

type Anchor interface {
	GetPosition() *geometry.Point
}
