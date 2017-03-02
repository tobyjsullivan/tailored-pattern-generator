package geometry

import (
	"fmt"
	"math"
)

type Line interface {
	StraightLines() []*StraightLine
	BoundingBox() *BoundingBox
	Length() float64
	PointAt(dist float64) *Point
	AngleAt(dist float64) *Angle
}

func boundingBoxOfLine(l Line) *BoundingBox {
	ls := l.StraightLines()
	lines := make([]BoundedShape, 0, len(ls))
	for _, l := range ls {
		lines = append(lines, l)
	}
	return CollectiveBoundingBox(lines...)
}

func lengthOfLine(l Line) float64 {
	accrued := 0.0
	for _, cur := range l.StraightLines() {
		accrued += cur.Length()
	}

	return accrued
}

func straightLineOnLine(l Line, dist float64) (*StraightLine, float64) {
	accruedLen := 0.0
	for _, sl := range l.StraightLines() {
		curLen := sl.Length()
		if accruedLen + curLen > dist || math.Abs((accruedLen + curLen) - dist) <= 0.001 {
			return sl, accruedLen
		}

		accruedLen += curLen
	}

	return nil, l.Length()
}

func pointOnLine(l Line, dist float64) *Point {
	sl, accruedLen := straightLineOnLine(l, dist)
	if sl != nil {
		return sl.PointAt(dist - accruedLen)
	}

	fmt.Printf("Tried to get point at %.3f of %v but length is only %.3f\n", dist, l, l.Length())
	panic("Cannot return a point that is beyond the end of the line.");
}

func angleAtPointOnLine(l Line, dist float64) *Angle {
	sl, accruedLen := straightLineOnLine(l, dist)
	if sl != nil {
		return sl.AngleAt(dist - accruedLen)
	}

	fmt.Printf("Tried to get angle at %.3f of %v but length is only %.3f\n", dist, l, l.Length())
	panic("Cannot return an angle that is beyond the end of the line.");
}

func SubLine(line Line, distStart float64, length float64) Line {
	if distStart + length > line.Length() {
		panic("Sub-line length exceeds original line length")
	}

	out := &Polyline{}

	pieces := 20
	pieceLength := length / float64(pieces)
	for i := 0; i < pieces; i++ {
		start := line.PointAt(distStart + float64(i) * pieceLength)
		end := line.PointAt(distStart + float64(i + 1) * pieceLength)

		out.AddLine(&StraightLine{
			Start: start,
			End: end,
		})
	}

	return out
}

func distanceToXIntersect(l Line, x float64) float64 {
	// Make sure point falls along line
	bbox := l.BoundingBox()
	if bbox.Left > x || bbox.Right < x {
		panic(fmt.Sprintf("X of %.2f falls outside of line %v.", x, l))
	}

	// Follow along the line until we hit x
	accruedLen := 0.0
	for _, sl := range l.StraightLines() {
		bbox := sl.BoundingBox()
		if bbox.Width() == 0.0 {
			continue
		}

		if bbox.Left <= x && bbox.Right >= x {
			p := (x - sl.Start.X) / (sl.End.X - sl.Start.X)
			h := math.Sqrt(math.Pow(p * (sl.End.X - sl.Start.X), 2.0) + math.Pow(p * (sl.End.Y - sl.Start.Y), 2.0))

			return accruedLen + h
		}

		accruedLen += sl.Length()
	}

	panic("Point on line not found")
}

func SliceLineVertically(l Line, x float64) Line {
	return SubLine(l, 0.0, distanceToXIntersect(l, x))
}

func MirrorLineHorizontally(line Line, x float64) Line {
	out := &Polyline{}

	for _, sl := range line.StraightLines() {
		out.AddLine(
			&StraightLine{
				Start: sl.Start.MirrorHorizontally(x),
				End: sl.End.MirrorHorizontally(x),
			},
		)
	}

	return out
}

func tangentAt(l Line, dist float64) *Tangent {
	sl, acc := straightLineOnLine(l, dist)
	return sl.TangentAt(dist - acc)
}

func TangentAtLineStart(l Line) *Tangent {
	return tangentAt(l, 0.0)
}

func TangentAtLineEnd(l Line) *Tangent {
	endEquiv := l.Length() - 0.001

	t := tangentAt(l, endEquiv)

	sl, _ := straightLineOnLine(l, endEquiv)

	return &Tangent{
		Origin: sl.End,
		Direction: t.Direction,
	}
}

func Connect(l0 Line, l1 Line) Line {
	t0 := TangentAtLineEnd(l0)
	t1 := TangentAtLineStart(l1)

	c0 := t0.Intersection(t1)

	poly := &Polyline{}
	if c0 == nil {
		return poly
	}

	poly.AddLine(
		&StraightLine{
			Start: t0.Origin,
			End: c0,
		},
		&StraightLine{
			Start: c0,
			End: t1.Origin,
		},
	)

	return poly
}
