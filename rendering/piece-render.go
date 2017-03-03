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
	err := drawPiece(pdf, p.Piece, 0.0, 0.0)
	if err != nil {
		return err
	}

	return pdf.SaveAs(filepath)
}

