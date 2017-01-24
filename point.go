package main

import (
	"math"
	"fmt"
)

type point struct {
	x float64
	y float64
}

func (p point) distanceTo(other point) float64 {
a := math.Abs(p.x - other.x)
b := math.Abs(p.y - other.y)

dist := math.Sqrt(math.Pow(a, 2.0) + math.Pow(b, 2.0))

return dist
}

func (p point) drawLeft(dist float64) point {
return point{p.x - dist, p.y}
}

func (p point) drawRight(dist float64) point {
return point{p.x + dist, p.y}
}

func (p point) drawUp(dist float64) point {
return point{p.x, p.y + dist}
}

func (p point) drawDown(dist float64) point {
return point{p.x, p.y - dist}
}

func (p point) squareToVerticalLine(x float64) point {
return point{x, p.y}
}

func (p point) squareToHorizontalLine(y float64) point {
return point{p.x, y}
}

func (p point) midpointTo(other point) point {
x := p.x + ((other.x - p.x) / 2)
y := p.y + ((other.y - p.y) / 2)

return point{x, y}
}

func (p point) String() string {
return fmt.Sprintf("[%.1f, %.1f]", p.x, p.y)
}
