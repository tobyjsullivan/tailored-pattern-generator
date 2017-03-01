package pieces

import (
	"github.com/tailored-style/pattern-generator/geometry"
)

type OpenOnFold struct {
	Piece
}

func (p OpenOnFold) x() float64 {
	return p.Piece.CutLayer().BoundingBox().Left
}

func (p *OpenOnFold) StitchLayer() *geometry.Block {
	layer := p.Piece.StitchLayer()

	if !p.Piece.OnFold() {
		return layer
	}

	mirrored := layer.MirrorHorizontally(p.x())
	layer.AddBlock(mirrored)

	return layer
}

func (p *OpenOnFold) CutLayer() *geometry.Block {
	layer := p.Piece.CutLayer()

	if !p.Piece.OnFold() {
		return layer
	}

	mirrored := layer.MirrorHorizontally(p.x())
	layer.AddBlock(mirrored)

	return layer
}

func (p *OpenOnFold) NotationLayer() *geometry.Block {
	layer := p.Piece.NotationLayer()

	if !p.Piece.OnFold() {
		return layer
	}

	mirrored := layer.MirrorHorizontally(p.x())
	layer.AddBlock(mirrored)

	return layer
}

func (p *OpenOnFold) OnFold() bool {
	return false
}

func (p *OpenOnFold) Details() *Details {
	return p.Piece.Details()
}