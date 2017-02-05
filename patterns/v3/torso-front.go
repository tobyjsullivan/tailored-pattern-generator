package v3

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
)

type TorsoFront struct {
	anchors map[string]*geometry.Point
}

func (p *TorsoFront) populateAnchors() error {
	if p.anchors != nil {
		return nil
	}

	a := make(map[string]*geometry.Point)

	a["0"] = &geometry.Point{X: 0.0, Y: 0.0}
	a["1"] = a["0"].SquareLeft(28.1)
	a["2"] = a["0"].SquareUp(21.0)
	a["3"] = a["0"].SquareLeft(22.9)
	a["4"] = a["1"].SquareToHorizontalLine(a["2"].Y)
	a["5"] = a["3"].SquareToHorizontalLine(a["2"].Y)
	a["6"] = a["3"].SquareUp(8.7)
	a["7"] = a["0"].SquareDown(15.9)
	a["8"] = a["7"].SquareDown(12.7)
	a["9"] = a["8"].SquareDown(9.5)
	a["10"] = a["9"].SquareDown(11.4)
	a["11"] = a["7"].SquareLeft(26.0)
	a["12"] = a["8"].SquareLeft(27.1)
	a["13"] = a["9"].SquareLeft(27.9)
	a["14"] = a["13"].SquareDown(7.0)
	a["15"] = a["2"].SquareLeft(7.5)
	a["16"] = a["15"].SquareUp(5.7)
	shoulderSlope := &geometry.StraightLine{Start: a["16"], End: a["5"]}
	shoulderLength := shoulderSlope.Start.DistanceTo(shoulderSlope.End) + 1.0
	a["17"] = shoulderSlope.Resize(shoulderLength).End
	a["18"] = a["1"].MidpointTo(a["16"])
	a["19"] = a["18"].DrawAt((&geometry.StraightLine{Start: a["1"], End: a["18"]}).PerpendicularAngle(), 1.6)

	p.anchors = a
	return nil
}

func (p *TorsoFront) CutLines() []geometry.Line {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	return []geometry.Line{}
}

func (p *TorsoFront) StitchLines() []geometry.Line {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	centreFront := &geometry.StraightLine{Start: p.anchors["2"], End: p.anchors["10"]}
	neckLine := &geometry.EllipseCurve{
		Start: p.anchors["2"],
		End: p.anchors["16"],
		StartingAngle: math.Pi * (3.0 / 2.0),
		ArcAngle: math.Pi / 3.0,
	}
	shoulder := &geometry.StraightLine{
		Start: p.anchors["16"],
		End: p.anchors["17"],
	}
	scyeTop := &geometry.EllipseCurve{
		Start: p.anchors["6"],
		End: p.anchors["17"],
		StartingAngle: 0.0,
		ArcAngle: math.Pi / 8.0,
	}
	scyeBottom := &geometry.EllipseCurve{
		Start: p.anchors["6"],
		End: p.anchors["1"],
		StartingAngle: math.Pi,
		ArcAngle: math.Pi * (2.0 / 5.0),
	}

	sideSeamA := &geometry.EllipseCurve{
		Start: p.anchors["11"],
		End: p.anchors["1"],
		StartingAngle: 0.0,
		ArcAngle: math.Pi / 4.0,
	}

	sideSeamB := &geometry.EllipseCurve{
		Start: p.anchors["11"],
		End: p.anchors["12"],
		StartingAngle: math.Pi,
		ArcAngle: math.Pi / 8.0,
	}

	sideSeamC := &geometry.EllipseCurve{
		Start: p.anchors["13"],
		End: p.anchors["12"],
		StartingAngle: 0.0,
		ArcAngle: math.Pi / 8.0,
	}

	sideSeamD := &geometry.StraightLine{
		Start: p.anchors["13"],
		End: p.anchors["14"],
	}

	hemLine := &geometry.SCurve{
		Start: p.anchors["14"],
		End: p.anchors["10"],
		StartingAngle: math.Pi * (3.0 / 2.0),
		FinishAngle: math.Pi * (3.0 / 2.0),
		MaxAngle: math.Pi / 8.0,
	}

	return []geometry.Line{
		centreFront,
		neckLine,
		shoulder,
		scyeTop,
		scyeBottom,
		sideSeamA,
		sideSeamB,
		sideSeamC,
		sideSeamD,
		hemLine,
	}
}

func (p *TorsoFront) FoldLines() []geometry.Line {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	return []geometry.Line{}
}

func (p *TorsoFront) GrainLine() *geometry.Line {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	return nil
}

func (p *TorsoFront) Notations() []geometry.Drawable {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	out := []geometry.Drawable{}

	// Draw all points (DEBUG)
	for k, p := range p.anchors {
		a := &anchor{
			Point: p,
			label: k,
		}
		out = append(out, a)
	}

	return out
}
