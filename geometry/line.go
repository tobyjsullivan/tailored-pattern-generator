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

func pointOnLine(l Line, dist float64) *Point {
	accruedLen := 0.0

	for _, sl := range l.StraightLines() {
		curLen := sl.Length()
		if accruedLen + curLen > dist || math.Abs((accruedLen + curLen) - dist) <= 0.001 {
			return sl.PointAt(dist - accruedLen)
		}

		accruedLen += curLen
	}

	fmt.Printf("Tried to get point at %.3f of %v but length is only %.3f\n", dist, l, l.Length())
	panic("Cannot return a point that is beyond the end of the line.");
}

func angleAtPointOnLine(l Line, dist float64) *Angle {
	accruedLen := 0.0

	for _, sl := range l.StraightLines() {
		curLen := sl.Length()
		if accruedLen + curLen > dist || math.Abs((accruedLen + curLen) - dist) <= 0.001 {
			return sl.AngleAt(0.0)
		}

		accruedLen += curLen
	}

	fmt.Printf("Tried to get angle at %.3f of %v but length is only %.3f\n", dist, l, l.Length())
	panic("Cannot return an angle that is beyond the end of the line.");
}
