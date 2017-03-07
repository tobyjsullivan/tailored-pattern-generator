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

func (p *MirroredPiece) StitchLayer() *geometry.Block {
	layer := p.Piece.Stitch()

	layer.AddBlock(
		layer.MirrorHorizontally(p.x()),
	)

	return layer
}

func (p *MirroredPiece) CutLayer() *geometry.Block {
	layer := p.Piece.InnerCut()

	layer.AddBlock(
		layer.MirrorHorizontally(p.x()),
	)

	return layer
}

func (p *MirroredPiece) NotationLayer() *geometry.Block {
	layer := p.Piece.Ink()

	layer.AddBlock(
		layer.MirrorHorizontally(p.x()),
	)

	return layer
}