package symbols

import "github.com/tailored-style/pattern-generator/geometry"

type ButtonHole struct {
	Centre *geometry.Point
	Length float64
	Angle  *geometry.Angle
}

func (b *ButtonHole) Block() *geometry.Block {
	blk := &geometry.Block{}

	a := b.Angle
	if a == nil {
		a = &geometry.Angle{Rads: 0.0}
	}

	start := b.Centre.DrawAt(a, b.Length / 2.0)
	end := b.Centre.DrawAt(a.Opposite(), b.Length / 2.0)

	blk.AddLine(
		&geometry.StraightLine{
			Start: start,
			End: end,
		},
		&geometry.StraightLine{
			Start: start.DrawAt(a.Perpendicular(), b.Length / 4.0),
			End: start.DrawAt(a.Perpendicular().Opposite(), b.Length / 4.0),
		},
		&geometry.StraightLine{
			Start: end.DrawAt(a.Perpendicular(), b.Length / 4.0),
			End: end.DrawAt(a.Perpendicular().Opposite(), b.Length / 4.0),
		},
	)

	return blk
}
