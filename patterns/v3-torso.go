package patterns

import (
	"math"
	"github.com/tailored-style/pattern-generator/patterns/body"
	"github.com/tailored-style/pattern-generator/patterns/sections"
	"github.com/tailored-style/pattern-generator/geometry3d"
	"github.com/tobyjsullivan/dxf/drawing"
	"github.com/tobyjsullivan/dxf"
)

type V3Torso struct {
	v3Body
}

type V3Measurements struct {
	ChestCircumference float64
	Height float64
	WaistCircumference float64
	NeckCircumference float64
	HipCircumference float64
	BiceptCircumference float64
}

type v3Body struct {
	neck *body.Neck
	chest *body.Chest
	waist *body.Waist
	hip *body.Hip
	bicept *body.Bicept
}

func NewV3Torso(m *V3Measurements) *V3Torso {
	torso := &v3Body{}

	torso.chest = &body.Chest{
		Circumference: m.ChestCircumference,
		Height: m.Height * (3.0 / 4.0),
	}

	torso.neck = &body.Neck{
		Circumference: m.NeckCircumference,
		Height: torso.chest.Height + (m.ChestCircumference / 4.0),
	}

	torso.waist = &body.Waist{
		Circumference: m.WaistCircumference,
		Height: torso.neck.Height - (m.Height / 4.0) + 1.0,
	}

	torso.hip = &body.Hip{
		Circumference: m.HipCircumference,
		Height: torso.waist.Height - (m.Height / 8.0),
	}

	torso.bicept = &body.Bicept{
		Circumference: m.BiceptCircumference,
	}

	return &V3Torso{*torso}
}

func (t *V3Torso) armhole() *sections.Armhole {
	chestCentre := &geometry3d.Point{
		X: 0.0,
		Y: 0.0,
		Z: t.chest.Height,
	}

	chestRadius := radiusFromCircumference(t.chest.Circumference)

	circ := t.bicept.Circumference + 5.1
	armholeRadius := circ / (2.0 * math.Pi)
	centre := chestCentre.ShiftX(chestRadius).ShiftZ(armholeRadius)

	return &sections.Armhole{
		Circumference: circ,
		Centre: centre,
	}
}

func (t *V3Torso) neckHole() *sections.NeckHole {
	return &sections.NeckHole{
		Circumference: t.neck.Circumference + 2.5,
		Height: t.neck.Height,
	}
}

func (t *V3Torso) shoulderLine() *sections.ShoulderLine {
	// Get edge of neck hole
	neckCentre := &geometry3d.Point{
		X: 0.0,
		Y: 0.0,
		Z: t.neck.Height,
	}
	neckHoleRadius := radiusFromCircumference(t.neckHole().Circumference)
	neckEdge :=  neckCentre.ShiftX(neckHoleRadius)

	// Get top edge of armhole
	armholeTop := t.armhole().Centre.ShiftZ(t.armhole().Radius())

	return &sections.ShoulderLine{
		NeckHolePoint: neckEdge,
		ArmholePoint: armholeTop,
	}
}

func (t *V3Torso) DrawDXF(d *drawing.Drawing) error {
	_, err := d.AddLayer("Default", dxf.DefaultColor, dxf.DefaultLineType, true)
	if err != nil {
		return err
	}

	// Draw Chest
	err = t.chest.DrawDXF(d)
	if err != nil {
		return err
	}

	// Draw waist
	err = t.waist.DrawDXF(d)
	if err != nil {
		return err
	}

	// Draw hip
	err = t.hip.DrawDXF(d)
	if err != nil {
		return err
	}

	// Draw neck hole
	err = t.neckHole().DrawDXF(d)
	if err != nil {
		return err
	}

	// Draw armhole
	err = t.armhole().DrawDXF(d)
	if err != nil {
		return err
	}

	// Draw shoulder line
	err = t.shoulderLine().DrawDXF(d)
	if err != nil {
		return err
	}

	return nil
}

func radiusFromCircumference(circ float64) float64 {
	return circ / (2.0 * math.Pi)
}

