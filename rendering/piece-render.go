package rendering

import (
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/drawing"
)

const (
	PIECE_PAGE_WIDTH = 91.44 // 36"
	PIECE_PAGE_HEIGHT = 91.44 // 36"
)

type PieceRender struct {
	pieces.Piece
}

func (p *PieceRender) SavePDF(filepath string) error {
	pdf := drawing.NewPDF(PIECE_PAGE_WIDTH, PIECE_PAGE_HEIGHT)
	err := p.drawPiece(pdf)
	if err != nil {
		return err
	}

	return pdf.SaveAs(filepath)
}

func (p *PieceRender) drawPiece(d drawing.Drawing) error {
	bbox := pieces.BoundingBox(p.Piece)

	offsetX := 0.0 - bbox.Left
	offsetY := 0.0 - bbox.Top

	var err error
	outer := p.Piece.OuterCut().Move(offsetX, offsetY)
	err = drawPolyline(d, outer)

	inner := p.Piece.InnerCut().Move(offsetX, offsetY)
	err = DrawBlock(d, inner)
	if err != nil {
		return err
	}

	ink := p.Piece.Ink().Move(offsetX, offsetY)
	err = DrawBlock(d, ink)
	if err != nil {
		return err
	}

	ref := p.Piece.Reference().Move(offsetX, offsetY)
	err = DrawBlock(d, ref)
	if err != nil {
		return err
	}

	return err
}

