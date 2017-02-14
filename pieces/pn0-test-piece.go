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

	parabolaA := &geometry.ParabolaCurve{
		Start: p.anchors["B"],
		End: p.anchors["C"],
		StartingAngle: 0.0,
		ArcAngle: math.Pi / 16.0,
	}

	layer := &geometry.Block{}
	layer.AddLine(
		parabolaA,
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