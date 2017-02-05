package v3

import "github.com/tailored-style/pattern-generator/geometry"

type Pattern struct {
	pieces []Piece
}

func (p *Pattern) populatePieces() error {
	if len(p.pieces) > 0 {
		return nil
	}

	p.pieces = []Piece{
		&TorsoFront{},
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