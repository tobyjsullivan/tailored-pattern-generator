package geometry

import (
	"math"
	"fmt"
)

type SpiralCurve struct {
	Start  *Point
	Length float64
	Scale float64
	StartingAngle float64
}

func (c *SpiralCurve) StraightLines() []*StraightLine {
	fmt.Println("Computing spiral curve.")
	lines := []*StraightLine{}

	l := 0.0

	origin := &Point{}
	lastPoint := c.Start

	pieces := 400

	stepLength := c.Length / float64(pieces)
	t := c.StartingAngle
	for l < c.Length {
		// Step circle
		// Reduce radius
		r := c.radiusAt(l)

		fmt.Printf("Radius is now %.14f, l is %.5f\n", r, l)

		// Update origin
		x := origin.X + r * math.Cos(t)
		y := origin.Y + r * math.Sin(t)
		offsetX := lastPoint.X - x
		offsetY := lastPoint.Y - y
		origin = origin.Move(offsetX, offsetY)

		// Step angle
		circ := 2 * math.Pi * r
		stepPortion := stepLength / circ
		t += (2.0 * math.Pi) * stepPortion

		// Draw line
		p := &Point{
			X: origin.X + r * math.Cos(t),
			Y: origin.Y + r * math.Sin(t),
		}
		line := &StraightLine{
			Start: lastPoint,
			End: p,
		}

		lines = append(lines, line)

		l += line.Start.DistanceTo(line.End)
		lastPoint = p
	}

	fmt.Println("Finished computing spiral curve.")
	return lines
}

func (c *SpiralCurve) radiusAt(l float64) float64 {
	if l < 0.001 {
		return c.Scale
	}
	return math.Pow(c.Scale, 1.0 / l) - 1.0
}

func (c *SpiralCurve) BoundingBox() *BoundingBox {
	ls := c.StraightLines()
	lines := make([]BoundedShape, 0, len(ls))
	for _, l := range ls {
		lines = append(lines, l)
	}
	return CollectiveBoundingBox(lines...)
}

//func FitToPoints(p0, p1 *Point, m0, m1 float64) *SpiralCurve {
//	e
//}