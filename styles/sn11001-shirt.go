package styles

import (
	"github.com/tailored-style/pattern-generator/pieces"
)

type SN11001Shirt struct {
	*pieces.Measurements
	pieces []pieces.Piece
}

func (p *SN11001Shirt) Pieces() []pieces.Piece {
	return []pieces.Piece{
		&pieces.PN4TorsoFront{
			Measurements: p.Measurements,
		},
		&pieces.PN5TorsoBack{
			Measurements: p.Measurements,
		},
		&pieces.PN6Yoke{
			Measurements: p.Measurements,
		},
	}
}

func (p *SN11001Shirt) Details() *Details {
	return &Details{
		Description: "Tailored Shirt - v3.0 TEST",
		StyleNumber: "11001",
		Measurements: p.Measurements,
	}
}
