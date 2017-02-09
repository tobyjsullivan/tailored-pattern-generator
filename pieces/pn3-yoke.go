package pieces

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
)

type PN3Yoke struct {
	anchors map[string]*geometry.Point
}

func (p *PN3Yoke) OnFold() bool {
	return true
}

func (p *PN3Yoke) populateAnchors() error {
	if p.anchors != nil {
		return nil
	}

	a := make(map[string]*geometry.Point)

	a["0"] = &geometry.Point{X: 0.0, Y: 0.0}
	a["1"] = a["0"].SquareUp(9.5)
	a["2"] = a["0"].SquareRight(24.0)
	a["3"] = a["1"].SquareRight(8.9)
	a["4"] = a["2"].SquareToHorizontalLine(a["1"].Y)
	a["5"] = a["3"].SquareUp(4.8)
	frontShoulderLength, err := p.shoulderLength()
	if err != nil {
		return err
	}
	a["6"] = (&geometry.StraightLine{Start: a["5"], End: a["4"]}).Resize(frontShoulderLength).End

	p.anchors = a

	return nil
}

func (p *PN3Yoke) shoulderLength() (float64, error) {
	return (&PN1TorsoFront{}).shoulderLength()
}

func (p *PN3Yoke) StitchLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	frontSeam := &geometry.StraightLine{
		Start: p.anchors["5"],
		End: p.anchors["6"],
	}

	backSeam := &geometry.StraightLine{
		Start: p.anchors["0"],
		End: p.anchors["2"],
	}

	neckLine := &geometry.EllipseCurve{
		Start: p.anchors["1"],
		End: p.anchors["5"],
		StartingAngle: math.Pi * (3.0 / 2.0),
		ArcAngle: math.Pi * (3.0 / 8.0),
	}

	scyeSeam := &geometry.StraightLine{
		Start: p.anchors["2"],
		End: p.anchors["6"],
	}

	layer := &geometry.Block{}
	layer.AddLine(
		frontSeam,
		backSeam,
		neckLine,
		scyeSeam,
	)

	return layer
}

func (p *PN3Yoke) CutLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	centreBack := &geometry.StraightLine{
		Start: p.anchors["0"],
		End: p.anchors["1"],
	}

	layer := &geometry.Block{}
	layer.AddLine(centreBack)

	return layer
}


func (p *PN3Yoke) NotationLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	layer := &geometry.Block{}

	// Draw all points (DEBUG)
	addAnchors(layer, p.anchors)

	return layer
}
