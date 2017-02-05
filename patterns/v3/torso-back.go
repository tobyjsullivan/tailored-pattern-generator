package v3

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
)

type TorsoBack struct {
	Origin geometry.Point
	anchors map[string]*geometry.Point
}

func (p *TorsoBack) computeAnchors() (map[string]*geometry.Point, error) {
	a := make(map[string]*geometry.Point)

	a["0"] = &p.Origin
	a["1"] = a["0"].SquareRight(28.4)
	a["2"] = a["0"].SquareUp(18.1)
	a["3"] = a["1"].SquareToHorizontalLine(a["2"].Y)
	a["4"] = a["0"].SquareRight(24.0)
	a["5"] = a["4"].SquareToHorizontalLine(a["2"].Y)
	a["6"] = a["2"].SquareRight(15.6)
	a["7"] = a["5"].SquareDown(1.1)
	a["8"] = a["5"].SquareDown(5.7)
	a["9"] = a["8"].SquareLeft(0.5)
	a["10"] = a["4"].DrawAt(math.Pi / 4.0, 2.5)
	a["11"] = a["0"].SquareDown(16.2)
	a["12"] = a["11"].SquareRight(26.0)
	a["13"] = a["11"].SquareDown(12.5)
	a["14"] = a["13"].SquareRight(27.1)
	a["15"] = a["0"].SquareDown(38.1)
	a["16"] = a["15"].SquareRight(27.9)
	a["17"] = a["15"].SquareDown(12.7)
	a["18"] = a["16"].SquareDown(7.3)
	a["19"] = a["17"].SquareRight(7.6)
	a["20"] = a["0"].SquareRight(12.1)
	a["21"] = a["20"].SquareDown(20.3)
	a["22"] = a["21"].SquareDown(19.1)
	a["23"] = a["21"].SquareRight(1.3)
	a["24"] = a["21"].SquareLeft(1.3)

	return a, nil
}

func (p *TorsoBack) populateAnchors() error {
	if p.anchors != nil {
		return nil
	}

	var err error
	p.anchors, err = p.computeAnchors()
	return err
}

func (p *TorsoBack) Dimensions() *geometry.Point {
	a, _ := p.computeAnchors()

	var minX, maxX, minY, maxY float64
	for _, p := range a {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	return &geometry.Point{
		X: maxX - minX,
		Y: maxY - maxY,
	}
}

func (p *TorsoBack) BoundingBox() *BoundingBox {
	a, _ := p.computeAnchors()

	var minX, maxX, minY, maxY float64
	for _, p := range a {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	return &BoundingBox{
		TopLeft: &geometry.Point{
			X: maxX,
			Y: minY,
		},
		BottomRight: &geometry.Point{
			X: minX,
			Y: maxX,
		},
	}
}

func (p *TorsoBack) CutLines() []geometry.Line {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	centreBack := &geometry.StraightLine{
		Start: p.anchors["2"],
		End: p.anchors["17"],
	}

	return []geometry.Line{
		centreBack,
	}
}

func (p *TorsoBack) StitchLines() []geometry.Line {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	yokeSeamA := &geometry.StraightLine{
		Start: p.anchors["2"],
		End: p.anchors["6"],
	}

	yokeSeamB := &geometry.EllipseCurve{
		Start: p.anchors["6"],
		End: p.anchors["7"],
		StartingAngle: math.Pi / 2.0,
		ArcAngle: math.Pi / 8.0,
	}

	scyeTop := &geometry.EllipseCurve{
		Start: p.anchors["9"],
		End: p.anchors["7"],
		StartingAngle: 0.0,
		ArcAngle: math.Pi / 8.0,
	}

	scyeBottom := &geometry.EllipseCurve{
		Start: p.anchors["9"],
		End: p.anchors["1"],
		StartingAngle: math.Pi,
		ArcAngle: math.Pi * (3.0 / 8.0),
	}

	sideSeamA := &geometry.EllipseCurve{
		Start: p.anchors["12"],
		End: p.anchors["1"],
		StartingAngle: 0.0,
		ArcAngle: math.Pi / 4.0,
	}

	sideSeamB := &geometry.EllipseCurve{
		Start: p.anchors["12"],
		End: p.anchors["14"],
		StartingAngle: math.Pi,
		ArcAngle: math.Pi / 8.0,
	}

	sideSeamC := &geometry.EllipseCurve{
		Start: p.anchors["16"],
		End: p.anchors["14"],
		StartingAngle: 0.0,
		ArcAngle: math.Pi / 8.0,
	}

	sideSeamD := &geometry.StraightLine{
		Start: p.anchors["16"],
		End: p.anchors["18"],
	}

	hemLineA := &geometry.StraightLine{
		Start: p.anchors["17"],
		End: p.anchors["19"],
	}

	hemLineB := &geometry.SCurve{
		Start: p.anchors["19"],
		End: p.anchors["18"],
		StartingAngle: math.Pi * (3.0 / 2.0),
		FinishAngle: math.Pi * (3.0 / 2.0),
		MaxAngle: math.Pi / 4.0,
	}

	return []geometry.Line{
		yokeSeamA,
		yokeSeamB,
		scyeTop,
		scyeBottom,
		sideSeamA,
		sideSeamB,
		sideSeamC,
		sideSeamD,
		hemLineA,
		hemLineB,
	}
}

func (p *TorsoBack) FoldLines() []geometry.Line {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	return []geometry.Line{}
}

func (p *TorsoBack) GrainLine() *geometry.Line {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	return nil
}

func (p *TorsoBack) Notations() []geometry.Drawable {
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
