package geometry3d

import "math"

type Point struct {
	X float64
	Y float64
	Z float64
}

func (p *Point) Shift(x, y, z float64) *Point {
	return &Point{
		X: p.X + x,
		Y: p.Y + y,
		Z: p.Z + z,
	}
}

func (p *Point) ShiftX(x float64) *Point {
	return p.Shift(x, 0.0, 0.0)
}

func (p *Point) ShiftY(y float64) *Point {
	return p.Shift(0.0, y, 0.0)
}

func (p *Point) ShiftZ(z float64) *Point {
	return p.Shift(0.0, 0.0, z)
}

func (p *Point) EuclideanDistanceTo(other *Point) float64 {
	x := other.X - p.X
	y := other.Y - p.Y
	z := other.Z - p.Z

	return math.Sqrt(math.Pow(x, 2.0) + math.Pow(y, 2.0) + math.Pow(z, 2.0))
}