package geometry

import (
	"math"
	"fmt"
)

type Tangent struct {
	Origin *Point
	Direction *Angle
}

func (t *Tangent) Intersects(other *Tangent) bool {
	eps := 0.00001

	m0 := t.m()
	m1 := other.m()

	// Equal slopes implies parallel lines
	if math.Abs(m0 - m1) < eps {
		return false
	}

	if math.IsInf(m0, 0) && math.IsInf(m1, 0) {
		return false
	}

	return true
}

func (t *Tangent) m() float64 {// Special case for vertical line (Tan(π/2) == NaN)
	eps := 0.000001
	if math.Abs(t.Direction.Radians() - math.Pi / 2.0) < eps || math.Abs(t.Direction.Radians() - -math.Pi / 2.0) < eps {
		return math.Inf(0)
	}

	return t.Direction.Tan()
}

func (t *Tangent) b() float64  {
	return t.Origin.Y - t.m() * t.Origin.X
}

func (t *Tangent) Intersection(other *Tangent) *Point {
	if !t.Intersects(other) {
		return nil
	}

	m0 := t.m()
	b0 := t.b()
	m1 := other.m()
	b1 := other.b()

	x := (b1 - b0) / (m0 - m1)
	y := m0 * x + b0

	// If either line is vertical, just find the point that the other line intersects x
	if math.IsInf(m0, 0) {
		x = t.Origin.X
		y = m1 * x + b1
	}
	if math.IsInf(m1, 0) {
		x = other.Origin.X
		y = m0 * x + b0
	}

	p := &Point{
		X: x,
		Y: y,
	}

	return p
}

func (t *Tangent) String() string {
	return fmt.Sprintf("[%.2f, %.2f, -> %.2f]", t.Origin.X, t.Origin.Y, t.Direction.Radians())
}
