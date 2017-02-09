package pieces

import "github.com/tailored-style/pattern-generator/geometry"

type Piece interface {
	StitchLayer() *geometry.Block
	CutLayer() *geometry.Block
	NotationLayer() *geometry.Block
	OnFold() bool
	Details() *Details
}

func addAnchors(b *geometry.Block, anchors map[string]*geometry.Point) {
	for k, p := range anchors {
		b.AddPoint(p)
		b.AddText(&geometry.Text{
			Content:  k,
			Position: p.Move(-1.5, -1.0),
		})
	}
}
