package rendering

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
)

type OpenOnFold struct {
	pieces.Piece
}

func (p OpenOnFold) x() float64 {
	return p.Piece.InnerCut().BoundingBox().Left
}

func (p *OpenOnFold) Stitch() *geometry.Block {
	layer := p.Piece.Stitch()

	if !p.Piece.OnFold() {
		return layer
	}

	mirrored := layer.MirrorHorizontally(p.x())
	layer.AddBlock(mirrored)

	return layer
}

func (p *OpenOnFold) InnerCut() *geometry.Block {
	layer := p.Piece.InnerCut()

	if !p.Piece.OnFold() {
		return layer
	}

	mirrored := layer.MirrorHorizontally(p.x())
	layer.AddBlock(mirrored)

	return layer
}

func (p *OpenOnFold) Ink() *geometry.Block {
	layer := p.Piece.Ink()

	if !p.Piece.OnFold() {
		return layer
	}

	mirrored := layer.MirrorHorizontally(p.x())
	layer.AddBlock(mirrored)

	return layer
}

func (p *OpenOnFold) OuterCut() *geometry.Polyline {
	out := &geometry.Polyline{}
	orig := p.Piece.OuterCut()

	out.AddLine(
		orig,
		geometry.MirrorLineHorizontally(orig, p.x()),
	)

	return out
}
