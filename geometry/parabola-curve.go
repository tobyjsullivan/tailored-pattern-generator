package geometry

import (
	"math"
	"fmt"
)

type ParabolaCurve struct {
	Start *Point
	End *Point
	StartingAngle float64
	ArcAngle float64
}

func (c *ParabolaCurve) y() float64 {
	return c.End.Y - c.Start.Y
}

func (c *ParabolaCurve) m() float64 {
	m := math.Tan(c.ArcAngle)
	fmt.Printf("Computed m as %.4f\n", m)
	return m
}

func (c *ParabolaCurve) xPrime() float64 {
	x := 2.0 * c.y() / c.m()
	fmt.Printf("Computed x' as %.4f\n", x)
	return x
}


func (c *ParabolaCurve) a() float64 {
	a := c.m() / (2.0 * c.xPrime())
	fmt.Printf("Computed a as %.4f\n", a)
	return a
}

func (c *ParabolaCurve) h() float64 {
	h := c.End.X - math.Sqrt(c.y() / c.a())
	fmt.Printf("Computed h as %.4f\n", h)
	return h
}

func (c *ParabolaCurve) f(x float64) float64 {
	if x < c.h() {
		fmt.Printf("Value is less than h, defaulting to 0.0 (val: %.4f)\n", x)
		return 0.0
	}

	f := c.a() * math.Pow(x - c.h(), 2.0)
	fmt.Printf("Computed f(%.4f) as %.4f\n", x, f)

	return f
}

func (c *ParabolaCurve) StraightLines() []*StraightLine {
	fmt.Printf("Drawing a parabola curve with Start: %v; End: %v; ArcAngle: %.2f\n", c.Start, c.End, c.ArcAngle)

	// Move start and end points so that start is at 0,0
	offsetX := -c.Start.X
	offsetY := -c.Start.Y

	tStart := c.Start.Move(offsetX, offsetY)
	tEnd := c.End.Move(offsetX, offsetY)
	fmt.Printf("Normalized start is %v and end is %v\n", tStart, tEnd)

	// Rotate end point around start/origin so that starting angle can be 0.0 rads
	//rotationAng := -c.StartingAngle
	//tEnd = rotatePointAboutPoint(tStart, tEnd, rotationAng)

	// Compute lines of initial tangent and arc
	lines := []*StraightLine{}

	tangent := &StraightLine{
		Start: tStart,
		End: &Point{
			X: c.h(),
			Y: tStart.Y,
		},
	}
	fmt.Printf("Tangent pre-move is %v\n", tangent)
	//tangent = rotateStraightLineAboutPoint(tStart, tangent, -rotationAng)
	tangent = tangent.Move(-offsetX, -offsetY)
	lines = append(lines, tangent)

	fmt.Printf("Tangent line is %v\n", tangent)

	pieces := int(tEnd.X - c.h())
	fmt.Printf("Drawing curve in %d pieces.\n", pieces)
	xIncr := (tEnd.X - c.h()) / float64(pieces)
	for i := 0; i < pieces; i++ {
		ls := c.h() + xIncr * float64(i)
		le := c.h() + xIncr * float64(i + 1)

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

		fmt.Printf("Drawing curve section %v\n", l)

		// Rotate all lines to match original starting angle
		//l = rotateStraightLineAboutPoint(tStart, l, -rotationAng)

		// Move all lines back to original start coordinates
		l = l.Move(-offsetX, -offsetY)

		lines = append(lines, l)
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

func rotateStraightLineAboutPoint(origin *Point, l *StraightLine, ang float64) *StraightLine {
	return &StraightLine{
		Start: rotatePointAboutPoint(origin, l.Start, ang),
		End: rotatePointAboutPoint(origin, l.End, ang),
	}
}

func rotatePointAboutPoint(origin, p *Point, ang float64) *Point {
	if origin.DistanceTo(p) == 0.0 {
		return p
	}

	r := origin.DistanceTo(p)

	currentAngle := math.Atan(p.Y - origin.Y / r)

	if p.Y - origin.Y < 0.0 {
		currentAngle += math.Pi
	}

	newAngle := currentAngle + ang

	return &Point{
		X: r * math.Cos(newAngle),
		Y: r * math.Sin(newAngle),
	}
}