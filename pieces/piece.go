package pieces

import "github.com/tailored-style/pattern-generator/geometry"

type Piece interface {
	StitchLayer() *geometry.Block
	CutLayer() *geometry.Block
	NotationLayer() *geometry.Block
	OnFold() bool
}
