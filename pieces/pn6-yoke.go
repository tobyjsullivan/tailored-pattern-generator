package pieces

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
)

type PN6Yoke struct {
	*Measurements
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

func (p *PN6Yoke) a() *geometry.Point {
	return &geometry.Point{X: 0.0, Y: 0.0}
}

func (p *PN6Yoke) b() *geometry.Point {
	return p.a().SquareDown(9.5)
}

func (p *PN6Yoke) c() *geometry.Point {
	return p.b().SquareRight(p.ChestCircumference/6.0 + 6.2)
}

func (p *PN6Yoke) d() *geometry.Point {
	return p.c().SquareToHorizontalLine(p.a().Y)
}

func (p *PN6Yoke) e() *geometry.Point {
	return p.a().SquareRight(p.NeckCircumference/8.0 + 3.7)
}

func (p *PN6Yoke) f() *geometry.Point {
	e := p.e()
	return e.SquareUp(p.a().DistanceTo(e)/2.0 + 0.3)
}

func (p *PN6Yoke) g() *geometry.Point {
	return  (&geometry.StraightLine{Start: p.f(), End: p.d()}).Resize(p.shoulderSeamLength()).End
}

func (p *PN6Yoke) shoulderSeamLength() float64 {
	return (&PN4TorsoFront{Measurements: p.Measurements}).shoulderStitch().Length()
}

func (p *PN6Yoke) necklineStitch() geometry.Line {
	return &geometry.EllipseCurve{
		Start:         p.a(),
		End:           p.f(),
		StartingAngle: &geometry.Angle{Rads: math.Pi * (3.0 / 2.0)},
		ArcAngle:      &geometry.Angle{Rads: math.Pi * (7.0 / 16.0)},
	}
}

func (p *PN6Yoke) frontStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.f(),
		End:   p.g(),
	}
}

func (p *PN6Yoke) armscyeStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.g(),
		End:   p.c(),
	}
}

func (p *PN6Yoke) backStitch() geometry.Line {
	return &geometry.StraightLine{
		Start: p.b(),
		End:   p.c(),
	}
}

func (p *PN6Yoke) centreBack() geometry.Line {
	return &geometry.StraightLine{
		Start: p.a(),
		End:   p.b(),
	}
}

func (p *PN6Yoke) StitchLayer() *geometry.Block {

	layer := &geometry.Block{}
	layer.AddLine(
		p.necklineStitch(),
		p.frontStitch(),
		p.armscyeStitch(),
		p.backStitch(),
	)

	return layer
}

func (p *PN6Yoke) CutLayer() *geometry.Block {
	layer := &geometry.Block{}
	layer.AddLine(
		p.centreBack(),
		addSeamAllowance(p.necklineStitch(), false),
		addSeamAllowance(p.frontStitch(), false),
		addSeamAllowance(p.armscyeStitch(), false),
		addSeamAllowance(p.backStitch(), true),
	)

	return layer
}

func (p *PN6Yoke) NotationLayer() *geometry.Block {
	layer := &geometry.Block{}

	// Draw all points (DEBUG)
	anchors := make(map[string]*geometry.Point)
	anchors["A"] = p.a()
	anchors["B"] = p.b()
	anchors["C"] = p.c()
	anchors["D"] = p.d()
	anchors["E"] = p.e()
	anchors["F"] = p.f()
	anchors["G"] = p.g()
	addAnchors(layer, anchors)

	return layer
}
