package geometry

import (
	"fmt"
	"math"
	"github.com/yofu/dxf/drawing"
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

func (p *Point) DrawDXF(label string, d *drawing.Drawing) error {
	_, err := d.Line(p.X - 0.5, p.Y, 0.0, p.X + 0.5, p.Y, 0.0)
	if err != nil {
		return err
	}

	_, err = d.Line(p.X, p.Y - 0.5, 0.0, p.X, p.Y + 0.5, 0.0)
	if err != nil {
		return err
	}

	_, err = d.Text(label, p.X - 1.0, p.Y + 1.0, 0.0, 1.0)

	return err
}