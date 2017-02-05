package v3

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/patterns"
)

type Pattern struct {
	pieces []Piece
}

func (p *Pattern) populatePieces() error {
	if len(p.pieces) > 0 {
		return nil
	}

	curX := 0.0

	back := &TorsoBack{}
	backBB := back.BoundingBox()
	back.Origin.X = curX + backBB.TopLeft.X

	curX += (backBB.BottomRight.X - backBB.TopLeft.X) + 5.0

	front := &TorsoFront{}
	frontBB := front.BoundingBox()
	front.Origin.X = curX + frontBB.TopLeft.X

	curX += (frontBB.BottomRight.X - frontBB.TopLeft.X) + 5.0

	p.pieces = []Piece{
		front,
		back,
	}

	return nil
}

func (p *Pattern) CutLines() []geometry.Line {
	err := p.populatePieces()
	if err != nil {
		panic(err)
	}

	out := []geometry.Line{}

	for _, piece := range p.pieces {
		out = append(out, piece.CutLines()...)
	}

	return out
}

func (p *Pattern) FoldLines() []geometry.Line {
	err := p.populatePieces()
	if err != nil {
		panic(err)
	}

	out := []geometry.Line{}

	for _, piece := range p.pieces {
		out = append(out, piece.FoldLines()...)
	}

	return out
}

func (p *Pattern) StitchLines() []geometry.Line {
	err := p.populatePieces()
	if err != nil {
		panic(err)
	}

	out := []geometry.Line{}

	for _, piece := range p.pieces {
		out = append(out, piece.StitchLines()...)
	}

	return out
}

func (p *Pattern) GrainLines() []geometry.Line {
	err := p.populatePieces()
	if err != nil {
		panic(err)
	}

	out := []geometry.Line{}

	for _, piece := range p.pieces {
		gl := piece.GrainLine()
		if gl != nil {
			out = append(out, *gl)
		}
	}

	return out
}

func (p *Pattern) Notations() []geometry.Drawable {
	err := p.populatePieces()
	if err != nil {
		panic(err)
	}

	out := []geometry.Drawable{}

	for _, piece := range p.pieces {
		out = append(out, piece.Notations()...)
	}

	return out
}

func (p *Pattern) Details() *patterns.PatternDetails {
	return &patterns.PatternDetails{
		Description: "Tailored Shirt - v3.0",
		StyleNumber: "11001",
	}
}