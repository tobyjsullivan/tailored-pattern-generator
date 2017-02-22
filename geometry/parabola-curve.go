package geometry

import (
	"math"
)

type ParabolaCurve struct {
	Start *Point
	End *Point
	StartingAngle *Angle
	ArcAngle *Angle
}

func (c *ParabolaCurve) m() float64 {
	return c.ArcAngle.Tan()
}

func (c *ParabolaCurve) a() float64 {
	return math.Pow(c.m(), 2.0) / (4.0 * (c.End.Y - c.Start.Y))
}

func (c *ParabolaCurve) h() float64 {
	return c.End.X - c.m() / (2.0 * c.a())
}

func (c *ParabolaCurve) f(x float64) float64 {
	return c.a() * math.Pow(x - c.h(), 2.0) + c.Start.Y
}

func (c *ParabolaCurve) rotatedStraightLines() []*StraightLine {
	normalized := &ParabolaCurve{
		Start: c.Start,
		End: c.End.RotateAround(c.Start, c.StartingAngle.Neg()),
		StartingAngle: &Angle{Rads: 0.0},
		ArcAngle: c.ArcAngle,
	}

	normLines := normalized.StraightLines()

	outLines := make([]*StraightLine, 0, len(normLines))

	for _, l := range normLines {
		outLines = append(outLines, l.RotateAround(c.Start, c.StartingAngle))
	}

	return outLines
}

func (c *ParabolaCurve) StraightLines() []*StraightLine {
	if c.StartingAngle != nil && c.StartingAngle.Radians() != 0.0 {
		return c.rotatedStraightLines()
	}

	lines := []*StraightLine{}

	h := c.h()
	if h != c.Start.X {
		lines = append(lines, &StraightLine{
			Start: c.Start,
			End: &Point{
				X: h,
				Y: c.Start.Y,
			},
		})
	}

	pieces := 20
	pieceSize := (c.End.X - h) / float64(pieces)

	for i := 0; i < pieces; i++ {
		sx := h + pieceSize * float64(i)
		ex := h + pieceSize * float64(i + 1)

		lines = append(lines, &StraightLine{
			Start: &Point{
				X: sx,
				Y: c.f(sx),
			},
			End: &Point{
				X: ex,
				Y: c.f(ex),
			},
		})
	}

	return lines
}

func (c *ParabolaCurve) BoundingBox() *BoundingBox {
	return boundingBoxOfLine(c)
}

func (c *ParabolaCurve) Length() float64 {
	return lengthOfLine(c)
}

func (c *ParabolaCurve) PointAt(dist float64) *Point {
	return pointOnLine(c, dist)
}

func (c *ParabolaCurve) AngleAt(dist float64) *Angle {
	return angleAtPointOnLine(c, dist)
}