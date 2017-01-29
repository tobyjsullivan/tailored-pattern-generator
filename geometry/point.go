package geometry

import (
	"math"
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) DistanceTo(other *Point) float64 {
	a := math.Abs(p.X - other.X)
	b := math.Abs(p.Y - other.Y)

	dist := math.Sqrt(math.Pow(a, 2.0) + math.Pow(b, 2.0))

	return dist
}

func (p *Point) DrawLeft(dist float64) *Point {
	return &Point{p.X - dist, p.Y}
}

func (p *Point) DrawRight(dist float64) *Point {
	return &Point{p.X + dist, p.Y}
}

func (p *Point) DrawUp(dist float64) *Point {
	return &Point{p.X, p.Y + dist}
}

func (p *Point) DrawDown(dist float64) *Point {
	return &Point{p.X, p.Y - dist}
}

func (p *Point) SquareToVerticalLine(x float64) *Point {
	return &Point{x, p.Y}
}

func (p *Point) SquareToHorizontalLine(y float64) *Point {
	return &Point{p.X, y}
}

func (p *Point) MidpointTo(other *Point) *Point {
	x := p.X + ((other.X - p.X) / 2)
	y := p.Y + ((other.Y - p.Y) / 2)

	return &Point{x, y}
}

func (p *Point) String() string {
	return fmt.Sprintf("[%.1f, %.1f]", p.X, p.Y)
}
