package geometry

type Polyline struct {
	Lines []*StraightLine
}

func (p *Polyline) StraightLines() []*StraightLine {
	return p.Lines
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

func (p *Polyline) AddLine(ls ...Line) {
	for _, l := range ls {
		p.Lines = append(p.Lines, l.StraightLines()...)
	}
}
