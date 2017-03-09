package geometry

type Polyline struct {
	points []*Point
}

func (p *Polyline) Points() []*Point {
	out := make([]*Point, len(p.points))
	copy(out, p.points)

	return out
}

func (p *Polyline) StraightLines() []*StraightLine {
	lines := []*StraightLine{}

	for i := 1; i < len(p.points); i++ {
		lines = append(lines, &StraightLine{
			Start: p.points[i - 1],
			End: p.points[i],
		})
	}

	return lines
}

func (p *Polyline) BoundingBox() *BoundingBox {
	return boundingBoxOfLine(p)
}

func (p *Polyline) Length() float64 {
	return lengthOfLine(p)
}

func (p *Polyline) PointAt(dist float64) *Point {
	return pointOnLine(p, dist)
}

func (p *Polyline) AngleAt(dist float64) *Angle {
	return angleAtPointOnLine(p, dist)
}

func (p *Polyline) lastPoint() *Point {
	if len(p.points) == 0 {
		return nil
	}

	return p.points[len(p.points) - 1]
}

func (p *Polyline) AddLine(ls ...Line) {
	for _, l := range ls {
		straightLines := l.StraightLines()
		if len(straightLines) == 0 {
			continue
		}

		newPoints := make([]*Point, 0, len(straightLines))

		if lastPoint := p.lastPoint(); lastPoint == nil || !lastPoint.Equals(straightLines[0].Start) {
			newPoints = append(newPoints, straightLines[0].Start)
		}

		for _, sl := range straightLines {
			newPoints = append(newPoints, sl.End)
		}

		p.points = append(p.points, newPoints...)
	}
}

func (p *Polyline) Move(x, y float64) *Polyline {
	points := make([]*Point, len(p.points))

	for i := range points {
		points[i] = p.points[i].Move(x, y)
	}

	return &Polyline{
		points: points,
	}
}
