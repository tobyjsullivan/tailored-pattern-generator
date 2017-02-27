package geometry

import "fmt"

type QuadraticBezierCurve struct {
	P0 *Point
	P1 *Point
	P2 *Point
}

type linearBezierCurve struct {
	p0 *Point
	p1 *Point
}

func (c *linearBezierCurve) b(t float64) *Point {
	if t - 1.0 > 0.001 {
		panic(fmt.Sprintf("t is out of bounds (%.3f)", t))
	}

	d := c.p0.DistanceTo(c.p1)

	return c.p0.DrawAt(c.p1.AngleRelativeTo(c.p0), d * t)
}

func (c *QuadraticBezierCurve) q0(t float64) *Point {
	return (&linearBezierCurve{
		p0: c.P0,
		p1: c.P1,
	}).b(t)
}

func (c *QuadraticBezierCurve) q1(t float64) *Point {
	return (&linearBezierCurve{
		p0: c.P1,
		p1: c.P2,
	}).b(t)
}

func (c *QuadraticBezierCurve) b(t float64) *Point {
	return (&linearBezierCurve{
		p0: c.q0(t),
		p1: c.q1(t),
	}).b(t)
}

func (c *QuadraticBezierCurve) StraightLines() []*StraightLine {
	numPieces := 20

	lines := make([]*StraightLine, numPieces)

	dt := 1.0 / float64(numPieces)
	t := 0.0
	for i := 0; i < numPieces; i++ {
		lines[i] = &StraightLine{
			Start: c.b(t),
			End: c.b(t + dt),
		}

		t += dt
	}

	return lines
}

func (p *QuadraticBezierCurve) BoundingBox() *BoundingBox {
	return boundingBoxOfLine(p)
}

func (p *QuadraticBezierCurve) Length() float64 {
	return lengthOfLine(p)
}

func (p *QuadraticBezierCurve) PointAt(dist float64) *Point {
	return pointOnLine(p, dist)
}

func (p *QuadraticBezierCurve) AngleAt(dist float64) *Angle {
	return angleAtPointOnLine(p, dist)
}


