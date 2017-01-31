package patterns

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
)

type TorsoSloper struct {
	ChestCircumference float64
	Height float64
	BackInterscyeLength float64
	ShoulderToShoulder float64
	ArmLength float64
	BicepCircumference float64
	SloperFullLength float64
}

func (p *TorsoSloper) GetPoints() map[string]*geometry.Point {
	d := &geometry.Point{X: 0.0, Y: 0.0}
	a := d.DrawRight(p.ChestCircumference / 2.0 + 5.1)

	c := d.DrawDown((p.Height / 4.0 + 1.0) + p.Height / 8.0)

	b := a.SquareToHorizontalLine(c.Y)
	//centreBack := &geometry.StraightLine{Start:a, End:b}
	//centreFront := &geometry.StraightLine{Start:c, End:d}

	interscyeAdjustment := (p.ChestCircumference - 106.7) * 0.3175
	// Round the result to one decimal
	interscyeAdjustment = math.Floor(interscyeAdjustment * 10.0 + 0.5) / 10.0
	e := a.DrawDown(p.ChestCircumference / 4.0 + interscyeAdjustment)

	f := a.DrawDown(p.Height / 4.0 + 1.0)

	g := e.SquareToVerticalLine(d.X)
	//chestCircumferenceLine := &geometry.StraightLine{Start:e, End: g}
	h := f.SquareToVerticalLine(d.X)
	//waistCircumferenceLine := &geometry.StraightLine{Start: f, End: h}

	i := e.DrawLeft(p.ChestCircumference / 6.0 + 3.8)
	j := i.SquareToHorizontalLine(d.Y)

	//backInterscyeLengthLine := &geometry.StraightLine{Start: i, End: j}

	k := g.DrawRight(i.DistanceTo(e) - 1.3)
	l := k.SquareToHorizontalLine(d.Y)
	//frontInterscyeLengthLine := &geometry.StraightLine{Start: k, End: l}

	m := g.DrawRight(g.DistanceTo(e) / 2.0 - 0.6)
	n := m.SquareToHorizontalLine(h.Y)
	o := m.SquareToHorizontalLine(c.Y)
	//sideSeamLine := &geometry.StraightLine{Start: m, End: o}

	ba := a.DrawLeft(p.ChestCircumference / 12.0 + 0.6)
	bb := ba.DrawUp(a.DistanceTo(ba) / 3.0)
	bc := j.DrawDown(1.6)

	run := bc.X - bb.X
	rise := bc.Y - bb.Y

	// We need an extra distance of 1.3 cm
	hyp := 1.3
	angle := math.Atan(rise/run)
	opp := hyp * math.Sin(angle)
	adj := math.Sqrt(math.Pow(hyp, 2.0) - math.Pow(opp, 2.0))

	bd := bc.DrawDown(opp).DrawLeft(adj)
	be := bd.SquareToVerticalLine(a.X)
	bf := bc.MidpointTo(i)
	bg := i.DrawUp(i.DistanceTo(j) / 4.0).DrawLeft(0.6)

	adj = math.Sqrt(math.Pow(i.DistanceTo(m) / 2, 2.0) / 2)
	bh := i.DrawUp(adj).DrawLeft(adj)

	za := a.DrawLeft(a.DistanceTo(ba) / 3.0)
	zb := l.DrawUp(5.0)

	return map[string]*geometry.Point{
		"A": a,
		"B": b,
		"C": c,
		"D": d,
		"E": e,
		"F": f,
		"G": g,
		"H": h,
		"I":i,
		"J":j,
		"K":k,
		"L":l,
		"M":m,
		"N":n,
		"O":o,
		"BA": ba,
		"BB": bb,
		"BC": bc,
		"BD": bd,
		"BE": be,
		"BF": bf,
		"BG": bg,
		"BH": bh,
		"ZA": za,
		"ZB": zb,
	}
}

func (p *TorsoSloper) GetCutLines() []geometry.Line {
	points := p.GetPoints()

	backNecklineStart := &geometry.StraightLine{
		Start: points["A"],
		End: points["ZA"],
	}
	backNecklineCurve := &geometry.CurvedLine{
		Start: points["ZA"],
		End: points["BB"],
	}
	shoulderLength := &geometry.StraightLine{
		Start: points["BB"],
		End: points["BD"],
	}
	backArmholeA := &geometry.CurvedLine{
		Start: points["BD"],
		End: points["BF"],
	}
	backArmholeB := &geometry.CurvedLine{
		Start: points["M"],
		End: points["BF"],
	}

	testCurve := &geometry.CurvedLine{
		Start: points["D"],
		End: points["ZB"],
	}

	return []geometry.Line{
		backNecklineStart,
		backNecklineCurve,
		shoulderLength,
		backArmholeA,
		backArmholeB,
		testCurve,
	}
}

func (p *TorsoSloper) GetFoldLines() []geometry.Line {
	return []geometry.Line{}
}

func (p *TorsoSloper) GetGrainLines() []geometry.Line {
	return []geometry.Line{}
}
