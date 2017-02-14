package geometry

import "fmt"

const PIECES_PER_LINE = 12

type EllipseCurve struct {
	Start         *Point
	End           *Point
	StartingAngle *Angle
	ArcAngle      *Angle
}

func (c *EllipseCurve) StraightLines() []*StraightLine {
	out := []*StraightLine{}

	arcStartAngle := c.StartingAngle
	arcEndAngle := arcStartAngle.Add(c.ArcAngle)

	// Simulate a simple circle until segment angle >= end angle
	sx := c.Start.X + arcStartAngle.Cos()
	sy := c.Start.Y + arcStartAngle.Sin()

	ex := c.Start.X + arcEndAngle.Cos()
	ey := c.Start.Y + arcEndAngle.Sin()

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
	chunkSize := c.ArcAngle.Divide(float64(numPieces))

	for i := 0; i < numPieces; i++ {
		sRad := arcStartAngle.Add(chunkSize.Multiply(float64(i)))
		eRad := arcStartAngle.Add(chunkSize.Multiply(float64(i + 1)))

		p1 := &Point{
			X: c.Start.X + scaleX*(sRad.Cos() + shiftX),
			Y: c.Start.Y + scaleY*(sRad.Sin() + shiftY),
		}

		p2 := &Point{
			X: c.Start.X + scaleX*(eRad.Cos() + shiftX),
			Y: c.Start.Y + scaleY*(eRad.Sin() + shiftY),
		}

		out = append(out, &StraightLine{Start: p1, End: p2})
	}

	return out
}

func (c *EllipseCurve) BoundingBox() *BoundingBox {
	return boundingBoxOfLine(c)
}

func (c *EllipseCurve) Length() float64 {
	return lengthOfLine(c)
}

func (c *EllipseCurve) PointAt(dist float64) *Point {
	return pointOnLine(c, dist)
}

func (c *EllipseCurve) AngleAt(dist float64) *Angle {
	return angleAtPointOnLine(c, dist)
}

func (c *EllipseCurve) String() string {
	return fmt.Sprintf("EllipseCurve[Start: %v, End: %v, StartingAngle: %v, ArcAngle: %v]", c.Start, c.End, c.StartingAngle, c.ArcAngle)
}