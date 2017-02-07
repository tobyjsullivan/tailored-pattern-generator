package styles

import (
	"github.com/tailored-style/pattern-generator/pieces"
)

type Style interface {
	Details() *Details
	Pieces() []pieces.Piece
}

type Details struct {
	Description string
	StyleNumber string
}
