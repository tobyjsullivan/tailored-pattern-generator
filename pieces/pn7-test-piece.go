package pieces

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
)

type PN7TestPiece struct {
	anchors map[string]*geometry.Point
}

func (p *PN7TestPiece) Details() *Details {
	return &Details{
		PieceNumber: "7",
		Description: "TEST",
	}
}


func (p *PN7TestPiece) OnFold() bool {
	return true
}

func (p *PN7TestPiece) populateAnchors() error {
	if p.anchors != nil {
		return nil
	}

	a := make(map[string]*geometry.Point)

	a["A"] = &geometry.Point{X: 0.0, Y: 0.0}
	a["B"] = a["A"].Move(10.0, 2.0)
	a["C"] = a["B"].Move(10.0, 2.0)

	a["D"] = a["A"].SquareUp(5.0)
	a["E"] = a["D"].Move(10.0, 2.0)
	a["F"] = a["E"].Move(8.0, 4.0)

	a["G"] = a["D"].SquareUp(10.0)

	a["H"] = a["G"].SquareUp(10.0)
	a["I"] = a["H"].Move(15.0, 8.0)

	a["J"] = a["H"].SquareUp(20.0)
	a["K"] = a["J"].SquareUp(20.0)

	a["Z"] = a["A"].Move(60.0, 60.0)

	p.anchors = a
	return nil
}

func (p *PN7TestPiece) CutLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	layer := &geometry.Block{}

	e1 := &geometry.Ellipse{
		XRadius: 2.0,
		YRadius: 1.0,
		Origin: p.anchors["G"],
	}

	// Draw the arc
	lines := []geometry.Line{}
	for _, l :=  range e1.Arc(0.0, math.Pi / 2.0) {
		lines = append(lines, l)
	}
	layer.AddLine(lines...)

	// Choose a tangent point
	t := e1.FindT(-3.0)

	// Add the point as an anchor
	p.anchors["P"] = e1.PointAt(t)

	tangent := func (p *geometry.Point, slope float64) geometry.Line {
		start := p.Move(1, slope)
		end := p.Move(-1, -slope)

		return &geometry.StraightLine{
			Start: start,
			End: end,
		}
	}

	layer.AddLine(tangent(p.anchors["P"], e1.Slope(t)))

	// Fit an ellipse to to points
	//a0 := 0.0
	//a1 := math.Pi / 4.0
	//e2, t0, t1 := geometry.FitEllipseToPoints(p.anchors["H"], p.anchors["I"], a0, a1)
	//// Draw the arc
	//lines = []geometry.Line{}
	//for _, l :=  range e2.Arc(t0, t1) {
	//	lines = append(lines, l)
	//}
	//layer.AddLine(lines...)
	//layer.AddLine(tangent(p.anchors["H"], math.Atan(a0)))
	//layer.AddLine(tangent(p.anchors["I"], math.Atan(a1)))

	spiral := &geometry.SpiralCurve{
		Start: p.anchors["J"],
		Length: 50.0,
		Scale: 3.0,
		StartingAngle: 0.0,
	}
	layer.AddLine(spiral)

	spiralB := &geometry.SpiralCurve{
		Start: p.anchors["K"],
		Length: 50.0,
		Scale: 12.0,
		StartingAngle: 0.0,
	}
	layer.AddLine(spiralB)

	layer.AddLine(
		&geometry.ThreePointCurve{
			Start: p.anchors["A"],
			Middle: p.anchors["B"],
			End: p.anchors["C"],
		},
		&geometry.ThreePointCurve{
			Start: p.anchors["D"],
			Middle: p.anchors["E"],
			End: p.anchors["F"],
		},
	)

	return layer
}


func (p *PN7TestPiece) StitchLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	return &geometry.Block{}
}

func (p *PN7TestPiece) NotationLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	layer := &geometry.Block{}

	// Draw all points (DEBUG)
	addAnchors(layer, p.anchors)

	return layer
}