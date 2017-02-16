package pieces

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
)

type PN0TestPiece struct {
	anchors map[string]*geometry.Point
}

func (p *PN0TestPiece) Details() *Details {
	return &Details{
		PieceNumber: "0",
		Description: "TEST PIECE",
	}
}

func (p *PN0TestPiece) OnFold() bool {
	return true
}
func (p *PN0TestPiece) populateAnchors() error {
	if p.anchors != nil {
		return nil
	}

	a := make(map[string]*geometry.Point)

	a["A"] = &geometry.Point{X: 0.0, Y: 0.0}
	a["B"] = a["A"].SquareUp(20.0)
	a["C"] = a["B"].SquareRight(50.0).SquareUp(10.0)

	a["E"] = &geometry.Point{X: 1, Y: 3}
	a["F"] = &geometry.Point{X: 12, Y: 6}
	a["G"] = &geometry.Point{X: 17, Y: 9}

	a["H"] = &geometry.Point{X: 12, Y: 4}

	a["I"] = &geometry.Point{X: 12, Y: 0}
	a["J"] = &geometry.Point{X: 17, Y: -3}

	a["K"] = &geometry.Point{X: 6, Y: 6}

	a["L"] = &geometry.Point{X: 9, Y: 6}
	a["M"] = &geometry.Point{X: 6, Y: 0}

	a["N"] = &geometry.Point{X: -6, Y: 19}
	a["O"] = a["E"].MidpointTo(a["N"]).SquareUp(2.0)
	a["P"] = a["E"].MidpointTo(a["N"]).SquareDown(2.0)

	a["Z"] = a["A"].Move(100.0, 100.0)

	p.anchors = a

	return nil
}

func (p *PN0TestPiece) StitchLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}
	layer := &geometry.Block{}
	return layer
}

func (p *PN0TestPiece) CutLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	threePointCurveA := &geometry.ThreePointCurve{
		Start: p.anchors["E"],
		Middle: p.anchors["F"],
		End: p.anchors["G"],
	}

	threePointCurveB := &geometry.ThreePointCurve{
		Start: p.anchors["E"],
		Middle: p.anchors["H"],
		End: p.anchors["G"],
	}

	threePointCurveC := &geometry.ThreePointCurve{
		Start: p.anchors["E"],
		Middle: p.anchors["I"],
		End: p.anchors["J"],
	}

	threePointCurveD := &geometry.ThreePointCurve{
		Start: p.anchors["E"],
		Middle: p.anchors["K"],
		End: p.anchors["G"],
	}

	threePointCurveE := &geometry.ThreePointCurve{
		Start: p.anchors["E"],
		Middle: p.anchors["L"],
		End: p.anchors["G"],
	}

	threePointCurveF := &geometry.ThreePointCurve{
		Start: p.anchors["E"],
		Middle: p.anchors["M"],
		End: p.anchors["J"],
	}

	threePointCurveG := &geometry.ThreePointCurve{
		Start: p.anchors["E"],
		Middle: p.anchors["O"],
		End: p.anchors["N"],
		Rotation: &geometry.Angle{Rads: math.Pi / 2.0},
	}

	threePointCurveH := &geometry.ThreePointCurve{
		Start: p.anchors["E"],
		Middle: p.anchors["P"],
		End: p.anchors["N"],
		Rotation: &geometry.Angle{Rads: math.Pi / 2.0},
	}

	layer := &geometry.Block{}
	layer.AddLine(
		threePointCurveA,
		threePointCurveB,
		threePointCurveC,
		threePointCurveD,
		threePointCurveE,
		threePointCurveF,
		threePointCurveG,
		threePointCurveH,
	)

	return layer
}


func (p *PN0TestPiece) NotationLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	layer := &geometry.Block{}

	// Draw all points (DEBUG)
	addAnchors(layer, p.anchors)

	return layer
}