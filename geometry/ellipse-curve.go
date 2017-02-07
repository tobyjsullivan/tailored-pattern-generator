package geometry

import (
	"math"
)

const PIECES_PER_LINE = 20

type EllipseCurve struct {
	Start         *Point
	End           *Point
	StartingAngle float64
	ArcAngle      float64
}

func (c *EllipseCurve) StraightLines() []*StraightLine {
	out := []*StraightLine{}

	arcStartAngle := c.StartingAngle
	arcEndAngle := arcStartAngle + c.ArcAngle

	// Simulate a simple circle until segment angle >= end angle
	sx := c.Start.X + math.Cos(arcStartAngle)
	sy := c.Start.Y + math.Sin(arcStartAngle)

	ex := c.Start.X + math.Cos(arcEndAngle)
	ey := c.Start.Y + math.Sin(arcEndAngle)

	// Shift is how much the arc start is offset from curved line start
	shiftX := c.Start.X - sx
	shiftY := c.Start.Y - sy

	// Compute the x and y scaling required to fit end point
	pointDistX := c.End.X - c.Start.X
	pointDistY := c.End.Y - c.Start.Y

	arcEndDistX := ex - sx
	arcEndDistY := ey - sy

	scaleX := pointDistX / arcEndDistX
	scaleY := pointDistY / arcEndDistY

	// Draw out the transform
	numPieces := PIECES_PER_LINE
	chunkSize := c.ArcAngle / float64(numPieces)

	for i := 0; i < numPieces; i++ {
		sRad := arcStartAngle + chunkSize*float64(i)
		eRad := arcStartAngle + chunkSize*float64(i+1)

		p1 := &Point{
			X: c.Start.X + scaleX*(math.Cos(sRad)+shiftX),
			Y: c.Start.Y + scaleY*(math.Sin(sRad)+shiftY),
		}

		p2 := &Point{
			X: c.Start.X + scaleX*(math.Cos(eRad)+shiftX),
			Y: c.Start.Y + scaleY*(math.Sin(eRad)+shiftY),
		}

		out = append(out, &StraightLine{Start: p1, End: p2})
	}

	return out
}
