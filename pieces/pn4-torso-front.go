package pieces

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
)

type PN4TorsoFront struct {
	Measurements
	anchors map[string]*geometry.Point
}

func (p *PN4TorsoFront) Details() *Details {
	return &Details{
		PieceNumber: "4",
		Description: "Torso Front",
	}
}


func (p *PN4TorsoFront) OnFold() bool {
	return false
}

func (p *PN4TorsoFront) populateAnchors() error {
	if p.anchors != nil {
		return nil
	}

	a := make(map[string]*geometry.Point)

	a["A"] = &geometry.Point{X: 0.0, Y: 0.0}
	a["B"] = a["A"].SquareDown(p.ChestCircumference / 4.0)
	a["C"] = a["B"].SquareLeft(p.ChestCircumference / 4.0 + 1.4)
	a["D"] = a["A"].SquareDown(p.Height / 4.0 - 3.2)
	a["E"] = a["D"].SquareLeft(p.WaistCircumference / 4.0 + 3.2)
	a["F"] = a["A"].SquareDown(p.Height / 3.0 - 5.7)
	a["G"] = a["F"].SquareLeft(p.HipCircumference / 4.0 - 0.6)
	a["H"] = a["A"].SquareDown(p.Height * (3.0/8.0) + 3.2)
	a["I"] = a["H"].SquareLeft(p.HipCircumference / 4.0 + 0.6)
	a["J"] = a["I"].SquareUp(7.0)
	a["K"] = a["H"].SquareDown(4.4)
	a["L"] = a["A"].SquareDown(p.NeckCircumference / 8.0 + 0.5)
	a["M"] = a["L"].SquareLeft(p.NeckCircumference / 8.0 + 2.2)
	a["N"] = a["M"].SquareToHorizontalLine(a["A"].Y)
	a["O"] = a["B"].SquareLeft(p.ChestCircumference / 6.0 + 5.1)
	a["P"] = a["O"].SquareToHorizontalLine(a["A"].Y)
	a["Q"] = a["P"].SquareDown(5.7)
	a["R"] = (&geometry.StraightLine{Start: a["N"], End: a["Q"]}).Resize(a["N"].DistanceTo(a["Q"]) + 1.3).End
	a["S"] = a["O"].SquareUp(a["O"].DistanceTo(a["Q"]) / 3.0 + 1.6)

	p.anchors = a
	return nil
}

func (p *PN4TorsoFront) CutLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	layer := &geometry.Block{}

	centreFront := &geometry.StraightLine{Start: p.anchors["L"], End: p.anchors["K"]}
	layer.AddLine(centreFront)

	return layer
}

func (p *PN4TorsoFront) StitchLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	layer := &geometry.Block{}

	neckLine := &geometry.EllipseCurve{
		Start: p.anchors["L"],
		End: p.anchors["N"],
		StartingAngle: math.Pi / 2.0,
		ArcAngle: math.Pi / 3.0,
	}

	shoulderLine := &geometry.StraightLine{
		Start: p.anchors["N"],
		End: p.anchors["R"],
	}

	armscyeTop := &geometry.EllipseCurve{
		Start: p.anchors["S"],
		End: p.anchors["R"],
		StartingAngle: 0.0,
		ArcAngle:      math.Pi / 8.0,
	}

	armscyeBottom := &geometry.EllipseCurve{
		Start: p.anchors["S"],
		End: p.anchors["C"],
		StartingAngle: math.Pi,
		ArcAngle:      math.Pi * (2.0 / 5.0),
	}

	sideSeamA := &geometry.EllipseCurve{
		Start:         p.anchors["E"],
		End:           p.anchors["C"],
		StartingAngle: 0.0,
		ArcAngle:      math.Pi / 4.0,
	}

	sideSeamB := &geometry.EllipseCurve{
		Start:         p.anchors["E"],
		End:           p.anchors["G"],
		StartingAngle: math.Pi,
		ArcAngle:      math.Pi / 8.0,
	}

	sideSeamC := &geometry.EllipseCurve{
		Start:         p.anchors["J"],
		End:           p.anchors["G"],
		StartingAngle: 0.0,
		ArcAngle:      math.Pi / 8.0,
	}

	sideSeamD := &geometry.StraightLine{
		Start: p.anchors["J"],
		End:   p.anchors["I"],
	}

	hemLine := &geometry.SCurve{
		Start: p.anchors["K"],
		End: p.anchors["I"],
		StartingAngle: math.Pi / 2.0,
		FinishAngle: math.Pi / 2.0,
		MaxAngle: math.Pi / 8.0,
	}

	layer.AddLine(
		neckLine,
		shoulderLine,
		armscyeTop,
		armscyeBottom,
		sideSeamA,
		sideSeamB,
		sideSeamC,
		sideSeamD,
		hemLine,
	)

	return layer
}

func (p *PN4TorsoFront) NotationLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	layer := &geometry.Block{}

	// Draw all points (DEBUG)
	addAnchors(layer, p.anchors)

	return layer
}

func (p*PN4TorsoFront) shoulderSeamLength() float64 {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	return p.anchors["N"].DistanceTo(p.anchors["R"])
}