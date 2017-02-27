package geometry

import "math"

type Circle struct {
	Origin *Point
	Radius float64
}

func (c *Circle) StraightLines() []*StraightLine {
	numPieces := 40

	lines := make([]*StraightLine, numPieces)

	dt := 2.0 * math.Pi / float64(numPieces)
	for i := 0; i < numPieces; i++ {
		st := dt * float64(i)
		et := dt * float64(i + 1)

		lines[i] = &StraightLine{
			Start: &Point{
				X: c.Origin.X + c.Radius * math.Cos(st),
				Y: c.Origin.Y + c.Radius * math.Sin(st),
			},
			End: &Point{
				X: c.Origin.X + c.Radius * math.Cos(et),
				Y: c.Origin.Y + c.Radius * math.Sin(et),
			},
		}
	}

	return lines
}

func (p *Circle) BoundingBox() *BoundingBox {
	return boundingBoxOfLine(p)
}

func (p *Circle) Length() float64 {
	return lengthOfLine(p)
}

func (p *Circle) PointAt(dist float64) *Point {
	return pointOnLine(p, dist)
}

func (p *Circle) AngleAt(dist float64) *Angle {
	return angleAtPointOnLine(p, dist)
}
