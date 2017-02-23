package geometry

import (
	"github.com/gonum/matrix/mat64"
	"math"
)

// A PolyNCurve interpolates a Nth-degree polynomial to n-1 points and two end angles
type PolyNCurve struct {
	Points []*Point
	StartAngle *Angle
	EndAngle *Angle

	xCache mat64.Matrix
}

func (c *PolyNCurve) n() int {
	return len(c.Points) + 1
}

func (c *PolyNCurve) x() mat64.Matrix {
	if c.xCache == nil {
		n := c.n()
		s := n+1
		initA := make([]float64, s*s)
		initB := make([]float64, s)

		for i := 0; i < n - 1; i++ {
			for j := 0; j <= n; j++ {
				initA[i*s + j] = math.Pow(c.Points[i].X, float64(j))
			}
			initB[i] = c.Points[i].Y
		}

		initA[(n-1)*s] = 0.0
		initA[n*s] = 0.0
		for i := 1; i <= n; i++ {
			initA[(n-1)*s + i] = float64(i) * math.Pow(c.Points[0].X, float64(i - 1))
			initA[n*s + i] = float64(i) * math.Pow(c.Points[len(c.Points) - 1].X, float64(i - 1))
		}
		initB[n-1] = c.StartAngle.Tan()
		initB[n] = c.EndAngle.Tan()

		a := mat64.NewDense(s, s, initA)

		b := mat64.NewDense(s, 1, initB)

		inv := mat64.NewDense(s, s, nil)
		inv.Inverse(a)

		x := mat64.NewDense(s, 1, nil)
		x.Mul(inv, b)

		c.xCache = x
	}

	return c.xCache
}

func (c *PolyNCurve) a(i int) float64 {
	return c.x().At(i, 0)
}

func (c *PolyNCurve) p(x float64) float64 {
	sum := 0.0
	for i := 0; i <= c.n(); i++ {
		sum += c.a(i) * math.Pow(x, float64(i))
	}

	return sum
}

func (c *PolyNCurve) StraightLines() []*StraightLine {
	pieces := 20 * len(c.Points)
	p0 := c.Points[0]
	pieceLength := (c.Points[len(c.Points) - 1].X - p0.X) / float64(pieces)

	lines := make([]*StraightLine, 0, pieces)
	for i := 0; i < pieces; i++ {
		sx := p0.X + pieceLength * float64(i)
		start := &Point{
			X: sx,
			Y: c.p(sx),
		}
		ex := p0.X + pieceLength * float64(i+1)
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


func (p *PolyNCurve) BoundingBox() *BoundingBox {
	return boundingBoxOfLine(p)
}

func (p *PolyNCurve) Length() float64 {
	return lengthOfLine(p)
}

func (p *PolyNCurve) PointAt(dist float64) *Point {
	return pointOnLine(p, dist)
}

func (p *PolyNCurve) AngleAt(dist float64) *Angle {
	return angleAtPointOnLine(p, dist)
}