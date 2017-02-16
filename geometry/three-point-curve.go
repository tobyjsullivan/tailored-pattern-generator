package geometry

import (
	"math"
	"fmt"
)

type ThreePointCurve struct {
	Start *Point
	Middle *Point
	End *Point
	Rotation *Angle
}

func (c *ThreePointCurve) x0() float64 {
	return c.Start.X
}

func (c *ThreePointCurve) y0() float64 {
	return c.Start.Y
}

func (c *ThreePointCurve) x1() float64 {
	return c.Middle.X
}

func (c *ThreePointCurve) y1() float64 {
	return c.Middle.Y
}

func (c *ThreePointCurve) x2() float64 {
	return c.End.X
}

func (c *ThreePointCurve) y2() float64 {
	return c.End.Y
}

func (c *ThreePointCurve) a2() float64 {
	return (c.y1() - c.y2()) / math.Pow(c.x1() - c.h2(), 2.0)
}

func (c *ThreePointCurve) m1() float64 {
	return (c.y1() - c.y0()) / (c.x1() - c.x0())
}

func (c *ThreePointCurve) m2() float64 {
	return (c.y2() - c.y1()) / (c.x2() - c.x1())
}

func (c *ThreePointCurve) h1() float64 {
	if math.Abs(c.m1()) > math.Abs(c.m2()) {
		return c.x0()
	}

	return c.x1() - (c.y1() - c.y0()) * (c.x1() - c.x2()) / (c.y1() - c.y2())
}

func (c *ThreePointCurve) h2() float64 {
	if math.Abs(c.m1()) < math.Abs(c.m2()) {
		return c.x2()
	}

	return c.x1() - (c.y1() - c.y2()) * (c.x1() - c.x0()) / (c.y1() - c.y0())
}

func (c *ThreePointCurve) a1() float64 {
	return (c.y1() - c.y0()) / math.Pow(c.x1() - c.h1(), 2.0)
}

func (c *ThreePointCurve) f(x float64) float64 {
	return c.a1() * math.Pow(x - c.h1(), 2.0) + c.y0()
}

func (c *ThreePointCurve) g(x float64) float64 {
	return c.a2() * math.Pow(x - c.h2(), 2.0) + c.y2()
}

func (c *ThreePointCurve) rotatedStraightLines() []*StraightLine {
	normalized := &ThreePointCurve{
		Start: c.Start,
		Middle: c.Middle.RotateAround(c.Start, c.Rotation.Neg()),
		End: c.End.RotateAround(c.Start, c.Rotation.Neg()),
	}

	normLines := normalized.StraightLines()

	outLines := make([]*StraightLine, 0, len(normLines))

	for _, l := range normLines {
		outLines = append(outLines, l.RotateAround(c.Start, c.Rotation))
	}

	return outLines
}

func (c *ThreePointCurve) StraightLines() []*StraightLine {
	if c.Rotation != nil && c.Rotation.Radians() != 0.0 {
		return c.rotatedStraightLines()
	}

	pieces := 20

	fmt.Printf("x0 is %.2f\n", c.x0())
	fmt.Printf("y0 is %.2f\n", c.y0())
	x1 := c.x1()
	fmt.Printf("x1 is %.2f\n", x1)
	fmt.Printf("y1 is %.2f\n", c.y1())
	fmt.Printf("x2 is %.2f\n", c.x2())
	fmt.Printf("y2 is %.2f\n", c.y2())

	h1 := c.h1()
	h2 := c.h2()
	fmt.Printf("h1 is %.2f\n", h1)
	fmt.Printf("h2 is %.2f\n", h2)
	fmt.Printf("a1 is %.2f\n", c.a1())
	fmt.Printf("a2 is %.2f\n", c.a2())

	out := make([]*StraightLine, 2, (2*pieces) + 2)

	// Draw the initial tangent line
	out[0] = &StraightLine{
		Start: c.Start,
		End: &Point{
			X: h1,
			Y: c.Start.Y,
		},
	}
	out[1] = &StraightLine{
		Start: c.End,
		End: &Point{
			X: h2,
			Y: c.End.Y,
		},
	}

	// Draw f(x)
	fLines := drawStraightLines(h1, x1, c.f, pieces)
	out = append(out, fLines...)

	// Draw g(x)
	gLines := drawStraightLines(x1, h2, c.g, pieces)
	out = append(out, gLines...)

	return out
}

func drawStraightLines(startX, endX float64, f func(float64) float64, numPieces int) []*StraightLine {
	out := []*StraightLine{}

	xIncr := (endX - startX) / float64(numPieces)
	for i := 0; i < numPieces; i++ {
		x0 := startX + xIncr * float64(i)
		x1 := startX + xIncr * float64(i + 1)

		line := &StraightLine{
			Start: &Point{
				X: x0,
				Y: f(x0),
			},
			End: &Point{
				X: x1,
				Y: f(x1),
			},
		}

		out = append(out, line)
	}

	return out
}


func (c *ThreePointCurve) BoundingBox() *BoundingBox {
	return boundingBoxOfLine(c)
}

func (c *ThreePointCurve) Length() float64 {
	return lengthOfLine(c)
}

func (c *ThreePointCurve) PointAt(dist float64) *Point {
	return pointOnLine(c, dist)
}

func (c *ThreePointCurve) AngleAt(dist float64) *Angle {
	return angleAtPointOnLine(c, dist)
}