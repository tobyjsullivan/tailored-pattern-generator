package pieces

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
)

type PN6Yoke struct {
	*Measurements
	anchors map[string]*geometry.Point
}

func (p *PN6Yoke) Details() *Details {
	return &Details{
		PieceNumber: "6",
		Description: "Yoke",
	}
}

func (p *PN6Yoke) OnFold() bool {
	return true
}

func (p *PN6Yoke) populateAnchors() error {
	if p.anchors != nil {
		return nil
	}

	a := make(map[string]*geometry.Point)

	a["A"] = &geometry.Point{X: 0.0, Y: 0.0}
	a["B"] = a["A"].SquareDown(9.5)
	a["C"] = a["B"].SquareRight(p.ChestCircumference/6.0 + 6.2)
	a["D"] = a["C"].SquareToHorizontalLine(a["A"].Y)
	a["E"] = a["A"].SquareRight(p.NeckCircumference/8.0 + 3.7)
	a["F"] = a["E"].SquareUp(a["A"].DistanceTo(a["E"])/2.0 + 0.3)
	a["G"] = (&geometry.StraightLine{Start: a["F"], End: a["D"]}).Resize(p.shoulderSeamLength()).End

	p.anchors = a

	return nil
}

func (p *PN6Yoke) shoulderSeamLength() float64 {
	return (&PN4TorsoFront{Measurements: p.Measurements}).shoulderStitch().Length()
}

func (p *PN6Yoke) StitchLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	neckLine := &geometry.EllipseCurve{
		Start:         p.anchors["A"],
		End:           p.anchors["F"],
		StartingAngle: &geometry.Angle{Rads: math.Pi * (3.0 / 2.0)},
		ArcAngle:      &geometry.Angle{Rads: math.Pi * (7.0 / 16.0)},
	}

	shoulderSeam := &geometry.StraightLine{
		Start: p.anchors["F"],
		End:   p.anchors["G"],
	}

	armscye := &geometry.StraightLine{
		Start: p.anchors["G"],
		End:   p.anchors["C"],
	}

	backSeam := &geometry.StraightLine{
		Start: p.anchors["B"],
		End:   p.anchors["C"],
	}

	layer := &geometry.Block{}
	layer.AddLine(
		neckLine,
		shoulderSeam,
		armscye,
		backSeam,
	)

	return layer
}

func (p *PN6Yoke) CutLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	centreBack := &geometry.StraightLine{
		Start: p.anchors["A"],
		End:   p.anchors["B"],
	}

	layer := &geometry.Block{}
	layer.AddLine(
		centreBack,
	)

	return layer
}

func (p *PN6Yoke) NotationLayer() *geometry.Block {
	err := p.populateAnchors()
	if err != nil {
		panic(err)
	}

	layer := &geometry.Block{}

	// Draw all points (DEBUG)
	addAnchors(layer, p.anchors)

	return layer
}
