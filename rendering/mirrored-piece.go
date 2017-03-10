package rendering

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/pieces"
)

type MirroredPiece struct {
	pieces.Piece
}

const MIRRORED_PIECE_MARGIN = 1.0

func (p MirroredPiece) x() float64 {
	bbox := geometry.CollectiveBoundingBox(
		p.Piece.InnerCut(),
		p.Piece.Stitch(),
		p.Piece.Ink(),
	)
	return bbox.Right + (MIRRORED_PIECE_MARGIN / 2.0)
}

func (p *MirroredPiece) Stitch() *geometry.Block {
	layer := p.Piece.Stitch()
	out := &geometry.Block{}

	out.AddBlock(
		layer.MirrorHorizontally(p.x()),
	)

	return out
}

func (p *MirroredPiece) InnerCut() *geometry.Block {
	layer := p.Piece.InnerCut()
	out := &geometry.Block{}

	out.AddBlock(
		layer.MirrorHorizontally(p.x()),
	)

	return out
}

func (p *MirroredPiece) Ink() *geometry.Block {
	layer := p.Piece.Ink()
	out := &geometry.Block{}

	out.AddBlock(
		layer.MirrorHorizontally(p.x()),
	)

	return out
}

func (p *MirroredPiece) Reference() *geometry.Block {
	layer := p.Piece.Reference()
	out := &geometry.Block{}

	out.AddBlock(
		layer.MirrorHorizontally(p.x()),
	)

	return out
}

func (p *MirroredPiece) OuterCut() *geometry.Polyline {
	out := &geometry.Polyline{}
	orig := p.Piece.OuterCut()

	out.AddLine(
		geometry.MirrorLineHorizontally(orig, p.x()),
	)

	return out
}