package symbols

import (
	"github.com/tailored-style/pattern-generator/geometry"
)

type Button struct {
	Centre   *geometry.Point
	Diameter float64
}

func (b *Button) Block() *geometry.Block {
	blk := &geometry.Block{}

	r := b.Diameter / 2.0

	c := geometry.Circle{
		Origin: b.Centre,
		Radius: r,
	}

	for _, sl := range c.StraightLines() {
		blk.AddLine(sl)
	}
	blk.AddLine(
		&geometry.StraightLine{
			Start: b.Centre.SquareLeft(r),
			End: b.Centre.SquareRight(r),
		},
		&geometry.StraightLine{
			Start: b.Centre.SquareUp(r),
			End: b.Centre.SquareDown(r),
		},
	)

	return blk
}
