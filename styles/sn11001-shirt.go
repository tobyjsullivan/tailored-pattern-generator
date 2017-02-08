package styles

import (
	"github.com/tailored-style/pattern-generator/pieces"
)

type SN11001Shirt struct {
	pieces []pieces.Piece
}

func (p *SN11001Shirt) Pieces() []pieces.Piece {
	return []pieces.Piece{
		&pieces.PN1TorsoFront{},
		&pieces.PN2TorsoBack{},
		&pieces.PN3Yoke{},
	}
}

func (p *SN11001Shirt) Details() *Details {
	return &Details{
		Description: "Tailored Shirt - v3.0",
		StyleNumber: "11001",
	}
}
