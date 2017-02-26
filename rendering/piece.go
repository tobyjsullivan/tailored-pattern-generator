package rendering

import (
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/drawing"
)

const (
	PIECE_PAGE_WIDTH = 91.44 // 36"
)

type Piece struct {
	pieces.Piece
}

func (p *Piece) SavePDF(filepath string) error {
	pdf := drawing.NewPDF(PIECE_PAGE_WIDTH)
	err := drawPiece(pdf, p.Piece, 0.0, 0.0)
	if err != nil {
		return err
	}

	return pdf.SaveAs(filepath)
}

