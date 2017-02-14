package pieces

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
)

type PN5TorsoBack struct {
	*Measurements
	anchors map[string]*geometry.Point
}

func (p *PN5TorsoBack) Details() *Details {
	return &Details{
		PieceNumber: "5",
		Description: "Torso Back",
	}
}

func (p *PN5TorsoBack) OnFold() bool {
	return true
}

func (p *PN5TorsoBack) populateAnchors() error {
	if p.anchors != nil {
		return nil
	}

	a := make(map[string]*geometry.Point)

	a["A"] = &geometry.Point{X: 0.0, Y: 0.0}
	a["B"] = a["A"].SquareDown(p.ChestCircumference/4.0 - 8.6)
	a["C"] = a["B"].SquareRight(p.ChestCircumference/4.0 + 1.6)
	a["D"] = a["A"].SquareDown(p.Height/4.0 - 11.4)
	a["E"] = a["D"].SquareRight(p.WaistCircumference/4.0 + 3.2)
	a["F"] = a["A"].SquareDown(p.Height*(7.0/24.0) - 6.4)
	a["G"] = a["F"].SquareRight(p.HipCircumference/4.0 - 0.2)
	a["H"] = a["A"].SquareDown(p.Height*(3.0/8.0) - 4.8)
	a["I"] = a["H"].SquareRight(p.HipCircumference/4.0 + 0.6)
	a["J"] = a["I"].SquareUp(7.3)
	a["K"] = a["H"].SquareDown(5.4)
	a["L"] = a["K"].SquareRight(7.6)
	a["M"] = a["B"].SquareRight(p.ChestCircumference/6.0 + 6.2)
	a["N"] = a["M"].SquareToHorizontalLine(a["A"].Y)
	a["O"] = a["N"].SquareDown(1.1)
	a["P"] = a["M"].SquareUp(a["M"].DistanceTo(a["O"])*(2.0/3.0) + 1.3)
	a["Q"] = a["N"].SquareLeft(8.4)
	a["R"] = a["P"].SquareLeft(0.5)
	a["S"] = a["B"].MidpointTo(a["M"])
	a["T"] = a["S"].SquareDown(p.Height/8.0 - 2.5)
	a["U"] = a["S"].SquareDown(a["S"].DistanceTo(a["T"])*2.0 - 3.8)
	a["V"] = a["T"].SquareLeft(1.3)
	a["W"] = a["T"].SquareRight(1.3)

	p.anchors = a
	return nil
}

func (p *PN5TorsoBack) CutLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	layer := &geometry.Block{}

	centreBack := &geometry.StraightLine{
		Start: p.anchors["A"],
		End:   p.anchors["K"],
	}

	layer.AddLine(
		centreBack,
	)

	return layer
}

func (p *PN5TorsoBack) StitchLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	layer := &geometry.Block{}

	yokeSeamA := &geometry.StraightLine{
		Start: p.anchors["A"],
		End:   p.anchors["Q"],
	}

	yokeSeamB := &geometry.EllipseCurve{
		Start:         p.anchors["Q"],
		End:           p.anchors["O"],
		StartingAngle: &geometry.Angle{Rads: math.Pi * (3.0 / 2.0)},
		ArcAngle:      &geometry.Angle{Rads: math.Pi / 8.0},
	}

	armscyeA := &geometry.EllipseCurve{
		Start:         p.anchors["R"],
		End:           p.anchors["O"],
		StartingAngle: &geometry.Angle{Rads: 0.0},
		ArcAngle:      &geometry.Angle{Rads: math.Pi / 8.0},
	}

	armscyeB := &geometry.EllipseCurve{
		Start:         p.anchors["R"],
		End:           p.anchors["C"],
		StartingAngle: &geometry.Angle{Rads: math.Pi},
		ArcAngle:      &geometry.Angle{Rads: math.Pi * 3.0 / 8.0},
	}

	sideSeamA := &geometry.EllipseCurve{
		Start:         p.anchors["E"],
		End:           p.anchors["C"],
		StartingAngle: &geometry.Angle{Rads: 0.0},
		ArcAngle:      &geometry.Angle{Rads: math.Pi / 8.0},
	}

	sideSeamB := &geometry.EllipseCurve{
		Start:         p.anchors["E"],
		End:           p.anchors["G"],
		StartingAngle: &geometry.Angle{Rads: math.Pi},
		ArcAngle:      &geometry.Angle{Rads: math.Pi / 16.0},
	}

	sideSeamC := &geometry.EllipseCurve{
		Start:         p.anchors["J"],
		End:           p.anchors["G"],
		StartingAngle: &geometry.Angle{Rads: 0.0},
		ArcAngle:      &geometry.Angle{Rads: math.Pi / 16.0},
	}

	sideSeamD := &geometry.StraightLine{
		Start: p.anchors["J"],
		End:   p.anchors["I"],
	}

	hemLineA := &geometry.StraightLine{
		Start: p.anchors["K"],
		End:   p.anchors["L"],
	}

	hemLineB := &geometry.SCurve{
		Start:         p.anchors["L"],
		End:           p.anchors["I"],
		StartingAngle: &geometry.Angle{Rads: math.Pi * (3.0 / 2.0)},
		FinishAngle:   &geometry.Angle{Rads: math.Pi * (3.0 / 2.0)},
		MaxAngle:      &geometry.Angle{Rads: math.Pi / 4.0},
	}

	layer.AddLine(
		yokeSeamA,
		yokeSeamB,
		armscyeA,
		armscyeB,
		sideSeamA,
		sideSeamB,
		sideSeamC,
		sideSeamD,
		hemLineA,
		hemLineB,
	)

	return layer
}

func (p *PN5TorsoBack) NotationLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	layer := &geometry.Block{}

	// Draw all points (DEBUG)
	addAnchors(layer, p.anchors)

	return layer
}
