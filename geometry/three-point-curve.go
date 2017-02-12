package geometry

import "math"

type ThreePointCurve struct {
	Start *Point
	Middle *Point
	End *Point
}

func (c *ThreePointCurve) innerCurves() []*EllipseCurve {
	curves := make([]*EllipseCurve, 2, 2)

	endpointAngle := math.Pi * (3.0 / 2.0)
	midpointAngle := math.Pi / 4.0

	curves[0] = &EllipseCurve{
		Start: c.Start,
		End: c.Middle,
		StartingAngle: endpointAngle,
		ArcAngle: midpointAngle,
	}

	curves[1] = &EllipseCurve{
		Start: c.End,
		End: c.Middle,
		StartingAngle: endpointAngle - math.Pi,
		ArcAngle: math.Pi / 2.0 - midpointAngle,
	}

	return curves
}

func (c *ThreePointCurve) StraightLines() []*StraightLine {
	out := []*StraightLine{}

	for _, inner := range c.innerCurves() {
		out = append(out, inner.StraightLines()...)
	}

	return out
}

func (c *ThreePointCurve) BoundingBox() *BoundingBox {
	ls := c.StraightLines()
	lines := make([]BoundedShape, 0, len(ls))
	for _, l := range ls {
		lines = append(lines, l)
	}
	return CollectiveBoundingBox(lines...)
}
