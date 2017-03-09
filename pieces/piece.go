package pieces

import (
	"github.com/tailored-style/pattern-generator/geometry"
)

const SEAM_ALLOWANCE = 1.0

type Piece interface {
	Stitch() *geometry.Block
	OuterCut() *geometry.Polyline
	InnerCut() *geometry.Block
	Ink() *geometry.Block
	Reference() *geometry.Block
	OnFold() bool
	Mirrored() bool
	Details() *Details
	CutCount() int
}

func AddSeamAllowance(l geometry.Line, opposite bool) geometry.Line {
	numPieces := 50
	segmentLength := l.Length() / float64(numPieces)

	result := &geometry.Polyline{}

	for i := 0; i < numPieces; i++ {
		s := segmentLength * float64(i)
		e := segmentLength * float64(i + 1)

		sAngle := l.AngleAt(s).Perpendicular()
		eAngle := l.AngleAt(e).Perpendicular()

		if opposite {
			sAngle = sAngle.Opposite()
			eAngle = eAngle.Opposite()
		}

		sPoint := l.PointAt(s).DrawAt(sAngle, SEAM_ALLOWANCE)
		ePoint := l.PointAt(e).DrawAt(eAngle, SEAM_ALLOWANCE)

		line := &geometry.StraightLine{
			Start: sPoint,
			End: ePoint,
		}

		result.AddLine(line)
	}

	return result
}

func Notch(stitch geometry.Line, dist float64, opp bool) geometry.Line {
	// Get the point on the stitch line
	t := geometry.TangentAt(stitch, dist)

	seamDirection := t.Direction.Perpendicular()
	if opp {
		seamDirection = seamDirection.Opposite()
	}

	// Move out to the seam line
	t.Origin = t.Origin.DrawAt(seamDirection, SEAM_ALLOWANCE)

	// Point direction inward
	t.Direction = seamDirection.Opposite()

	return &geometry.StraightLine{
		Start: t.Origin,
		End: t.Origin.DrawAt(t.Direction, SEAM_ALLOWANCE / 2.0),
	}
}

func BoundingBox(p Piece) *geometry.BoundingBox {
	return geometry.CollectiveBoundingBox(
		p.OuterCut(),
		p.InnerCut(),
		p.Stitch(),
		p.Ink(),
	)
}
