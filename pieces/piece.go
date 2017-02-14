package pieces

import "github.com/tailored-style/pattern-generator/geometry"

const SEAM_ALLOWANCE = 1.0

type Piece interface {
	StitchLayer() *geometry.Block
	CutLayer() *geometry.Block
	NotationLayer() *geometry.Block
	OnFold() bool
	Details() *Details
}

func addAnchors(b *geometry.Block, anchors map[string]*geometry.Point) {
	for k, p := range anchors {
		addAnchor(b, k, p)
	}
}

func addAnchor(b *geometry.Block, label string, p *geometry.Point) {
	b.AddPoint(p)
	b.AddText(&geometry.Text{
		Content:  label,
		Position: p.Move(-1.5, -1.0),
	})
}

func addSeamAllowance(l geometry.Line, opposite bool) geometry.Line {
	numPieces := 20
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

		result.AddLine(&geometry.StraightLine{
			Start: sPoint,
			End: ePoint,
		})
	}

	return result
}
