package geometry

import "math"

type ParabolaCurve struct {
	Start *Point
	End *Point
	StartingAngle float64
	ArcAngle float64
}

func (c *ParabolaCurve) m() float64 {
	return math.Tan(c.ArcAngle)
}

func (c *ParabolaCurve) xPrime() float64 {
	return 2.0 * c.End.Y / c.m()
}


func (c *ParabolaCurve) a() float64 {
	return c.m() / (2.0 * c.xPrime())
}

func (c *ParabolaCurve) h() float64 {
	return c.End.X - math.Sqrt(c.End.Y / c.a())
}

func (c *ParabolaCurve) f(x float64) float64 {
	if x < c.h() {
		return 0
	}

	return c.a() * math.Pow(x - c.h(), 2.0)
}

func (c *ParabolaCurve) StraightLines() []*StraightLine {
	// Move start and end points so that start is at 0,0
	offsetX := -c.Start.X
	offsetY := -c.Start.Y

	tStart := c.Start.Move(offsetX, offsetY)
	tEnd := c.End.Move(offsetX, offsetY)

	// Rotate end point around start/origin so that starting angle can be 0.0 rads
	rotationAng := -c.StartingAngle
	tEnd = rotatePointAboutPoint(tStart, tEnd, rotationAng)

	// Compute lines of initial tangent and arc
	lines := []*StraightLine{}

	tangent := &StraightLine{
		Start: tStart,
		End: &Point{
			X: c.h(),
			Y: tStart.Y,
		},
	}
	tangent = rotateStraightLineAboutPoint(tStart, tangent, -rotationAng)
	tangent = tangent.Move(-offsetX, -offsetY)

	pieces := int(math.Ceil(tEnd.X - tStart.X))
	xIncr := tEnd.X - tStart.X / float64(pieces)
	for i := 0; i < pieces; i++ {
		ls := xIncr * float64(i)
		le := xIncr * float64(i + 1)

		l := &StraightLine{
			Start: &Point{
				X: ls,
				Y: c.f(ls),
			},
			End: &Point{
				X: le,
				Y: c.f(le),
			},
		}

		// Rotate all lines to match original starting angle
		l = rotateStraightLineAboutPoint(tStart, l, -rotationAng)

		// Move all lines back to original start coordinates
		l = l.Move(-offsetX, -offsetY)

		lines = append(lines, l)
	}

	return lines
}

func (c *ParabolaCurve) BoundingBox() *BoundingBox {
	ls := c.StraightLines()
	lines := make([]BoundedShape, 0, len(ls))
	for _, l := range ls {
		lines = append(lines, l)
	}
	return CollectiveBoundingBox(lines...)
}

func rotateStraightLineAboutPoint(origin *Point, l *StraightLine, ang float64) *StraightLine {
	return &StraightLine{
		Start: rotatePointAboutPoint(origin, l.Start, ang),
		End: rotatePointAboutPoint(origin, l.End, ang),
	}
}

func rotatePointAboutPoint(origin, p *Point, ang float64) *Point {
	currentAngle := math.Atan(p.Y - origin.Y / p.X - origin.X)

	if p.Y - origin.Y < 0.0 {
		currentAngle += math.Pi
	}

	newAngle := currentAngle + ang

	r := origin.DistanceTo(p)

	return &Point{
		X: r * math.Cos(newAngle),
		Y: r * math.Sin(newAngle),
	}
}