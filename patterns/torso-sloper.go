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

	backNeckWidth := p.ChestCircumference / 12.0 + 0.6
	ba := a.DrawLeft(backNeckWidth)
	bb := ba.DrawUp(a.DistanceTo(ba) / 3.0)
	bc := j.DrawDown(1.6)

	// We need an extra distance of 1.3 cm
	bd := drawPointOnLine(bb, bc, bb.DistanceTo(bc) + 1.3)

	be := bd.SquareToVerticalLine(a.X)
	bf := bc.MidpointTo(i)
	bg := i.DrawUp(i.DistanceTo(j) / 4.0).DrawLeft(0.6)

	bh := i.DrawAt(math.Pi * (3.0 / 4.0), i.DistanceTo(m) / 2)

	fi := d.DrawDown(backNeckWidth + 0.6)
	fj := fi.DrawRight(backNeckWidth - 0.3)
	fk := fj.SquareToHorizontalLine(d.Y)
	fl := fi.MidpointTo(fk)

	// Draw a 2.2 cm perpendicular line at fl
	hyp := 2.2
	ang := math.Atan((fl.Y - fi.Y) / (fl.X - fi.X))
	down := hyp * math.Sin(ang)
	right := hyp * math.Cos(ang)
	fm := fl.DrawDown(down).DrawRight(right)

	fn := l.DrawDown(3.8)

	backShoulderLength := bb.DistanceTo(bd)
	fo := drawPointOnLine(fk, fn, backShoulderLength)

	fp := k.DrawUp(k.DistanceTo(fn) / 3.0 + 1.6)
	fq := k.DrawAt(math.Pi / 4.0, i.DistanceTo(bh) - 0.6)

	fr := c.DrawDown(1.6)

	za := a.DrawLeft(a.DistanceTo(ba) / 3.0)

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
		"FI": fi,
		"FJ": fj,
		"FK": fk,
		"FL": fl,
		"FM": fm,
		"FN": fn,
		"FO": fo,
		"FP": fp,
		"FQ": fq,
		"FR": fr,
		"ZA": za,
	}
}

func angleOfLine(p0 *geometry.Point, p1 *geometry.Point) float64 {
	run := (p1.X - p0.X)

	angle := math.Atan((p1.Y - p0.Y)/run)

	if run < 0.0 {
		angle += math.Pi
	}

	return angle
}

func drawPointOnLine(p0 *geometry.Point, p1 *geometry.Point, distanceFromP0 float64) *geometry.Point {
	run := (p1.X - p0.X)

	angle := math.Atan((p1.Y - p0.Y)/run)

	if run < 0.0 {
		angle += math.Pi
	}

	return p0.DrawAt(angle, distanceFromP0)
}

func (p *TorsoSloper) GetCutLines() []geometry.Line {
	points := p.GetPoints()

	backNecklineStart := &geometry.StraightLine{
		Start: points["A"],
		End: points["ZA"],
	}
	backShoulderAngle := angleOfLine(points["BD"], points["BB"])
	backNecklineCurve := &geometry.EllipseCurve{
		Start: points["ZA"],
		End: points["BB"],
		StartingAngleRads: math.Pi / 2.0,
		ArcAngle: backShoulderAngle + math.Pi / 4.0,
	}
	backShoulderLength := &geometry.StraightLine{
		Start: points["BB"],
		End: points["BD"],
	}
	backArmholeA := &geometry.EllipseCurve{
		Start: points["BF"],
		End: points["BD"],
		StartingAngleRads: 0.0,
		ArcAngle: backShoulderAngle,
	}
	backArmholeB := &geometry.EllipseCurve {
		Start: points["M"],
		End: points["BF"],
		StartingAngleRads: math.Pi * (3.0 / 2.0),
		ArcAngle: math.Pi / 2.0,
	}
	backSideSeam := &geometry.StraightLine{
		Start: points["M"],
		End: points["O"],
	}

	backHemLine := &geometry.StraightLine{
		Start: points["O"],
		End: points["B"],
	}
	centreBack := &geometry.StraightLine{
		Start: points["A"],
		End: points["B"],
	}

	frontShoulderAngle := angleOfLine(points["FK"], points["FO"])
	frontNeckline := &geometry.EllipseCurve{
		Start: points["FI"],
		End: points["FK"],
		StartingAngleRads: math.Pi * (3.0 / 2.0),
		ArcAngle: -frontShoulderAngle,
	}

	frontArmholeA := &geometry.EllipseCurve {
		Start: points["FP"],
		End: points["FO"],
		StartingAngleRads: math.Pi,
		ArcAngle: frontShoulderAngle,
	}

	frontArmholeB := &geometry.EllipseCurve{
		Start: points["FP"],
		End: points["M"],
		StartingAngleRads: math.Pi,
		ArcAngle: math.Pi / 2.0,
	}

	frontShoulderLength := &geometry.StraightLine{
		Start: points["FK"],
		End: points["FO"],
	}

	centreFront := &geometry.StraightLine{
		Start: points["FI"],
		End: points["FR"],
	}

	return []geometry.Line{
		backNecklineStart,
		backNecklineCurve,
		backShoulderLength,
		backArmholeA,
		backArmholeB,
		backSideSeam,
		backHemLine,
		centreBack,
		frontNeckline,
		frontArmholeA,
		frontArmholeB,
		frontShoulderLength,
		centreFront,
	}
}

func (p *TorsoSloper) GetFoldLines() []geometry.Line {
	return []geometry.Line{}
}

func (p *TorsoSloper) GetGrainLines() []geometry.Line {
	return []geometry.Line{}
}


func (p *TorsoSloper) GetLabels() []geometry.Drawable {
	points := p.GetPoints()
	back := NewLabel("BACK", points["M"].MidpointTo(points["F"]))
	back.Size = TEXT_SIZE_LARGE

	centreBack := NewLabel("Centre Back", points["F"].MidpointTo(points["B"]))
	centreBack.Rotation = math.Pi / 2.0
	centreBack.Size = TEXT_SIZE_SMALL

	foldLine := NewLabel("FOLD", points["F"].MidpointTo(points["E"]))
	foldLine.Rotation = math.Pi / 2.0
	foldLine.Size = TEXT_SIZE_SMALL

	front := NewLabel("FRONT", points["N"].MidpointTo(points["G"]))
	front.Size = TEXT_SIZE_LARGE

	return []geometry.Drawable{
		back,
		centreBack,
		foldLine,
		front,
	}
}
