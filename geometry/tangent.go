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

	m0 := t.Direction.Tan()
	m1 := other.Direction.Tan()

	// Equal slopes implies parallel lines
	if math.Abs(m0 - m1) < eps {
		return false
	}

	return true
}

func (t *Tangent) m() float64 {
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

	return &Point{
		X: x,
		Y: y,
	}
}

func (t *Tangent) String() string {
	return fmt.Sprintf("[%.2f, %.2f, -> %.2f]", t.Origin.X, t.Origin.Y, t.Direction.Radians())
}
