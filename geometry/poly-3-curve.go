package geometry

import (
	"math"
	"github.com/gonum/matrix/mat64"
)

// A Poly3Curve interpolates a 3rd-degree polynomial to two points and two angles
type Poly3Curve struct {
	P0     *Point
	P1     *Point
	A0     *Angle
	A1     *Angle

	xCache mat64.Matrix
}

func (c *Poly3Curve) x() mat64.Matrix {
	if c.xCache == nil {
		a := mat64.NewDense(4, 4, []float64{
			1, c.P0.X, math.Pow(c.P0.X, 2.0), math.Pow(c.P0.X, 3.0),
			1, c.P1.X, math.Pow(c.P1.X, 2.0), math.Pow(c.P1.X, 3.0),
			0, 1, 2.0 * c.P0.X, math.Pow(c.P0.X, 2.0),
			0, 1, 2.0 * c.P1.X, math.Pow(c.P1.X, 2.0),
		})

		b := mat64.NewDense(4, 1, []float64{
			c.P0.Y,
			c.P1.Y,
			c.A0.Tan(),
			c.A1.Tan(),
		})

		inv := mat64.NewDense(4, 4, nil)
		inv.Inverse(a)

		x := mat64.NewDense(4, 1, nil)
		x.Mul(inv, b)

		c.xCache = x
	}

	return c.xCache
}

func (c *Poly3Curve) a() float64 {
	x := c.x()
	return x.At(3, 0)
}

func (c *Poly3Curve) b() float64 {
	return c.x().At(2, 0)
}

func (c *Poly3Curve) c() float64 {
	return c.x().At(1, 0)
}

func (c *Poly3Curve) d() float64 {
	return c.x().At(0, 0)
}

func (c *Poly3Curve) p(x float64) float64 {
	return c.a() * math.Pow(x, 3.0) + c.b() * math.Pow(x, 2.0) + c.c() * x + c.d()
}

func (c *Poly3Curve) StraightLines() []*StraightLine {
	pieces := 40
	pieceLength := (c.P1.X - c.P0.X) / float64(pieces)

	lines := make([]*StraightLine, 0, pieces)
	for i := 0; i < pieces; i++ {
		sx := c.P0.X + pieceLength * float64(i)
		start := &Point{
			X: sx,
			Y: c.p(sx),
		}
		ex := c.P0.X + pieceLength * float64(i+1)
		end := &Point{
			X: ex,
			Y: c.p(ex),
		}
		lines = append(lines, &StraightLine{
			Start: start,
			End: end,
		})
	}

	return lines
}


func (p *Poly3Curve) BoundingBox() *BoundingBox {
	return boundingBoxOfLine(p)
}

func (p *Poly3Curve) Length() float64 {
	return lengthOfLine(p)
}

func (p *Poly3Curve) PointAt(dist float64) *Point {
	return pointOnLine(p, dist)
}

func (p *Poly3Curve) AngleAt(dist float64) *Angle {
	return angleAtPointOnLine(p, dist)
}
