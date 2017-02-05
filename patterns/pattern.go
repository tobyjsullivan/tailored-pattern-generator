package patterns

import (
	"github.com/tobyjsullivan/dxf/drawing"
	"github.com/tobyjsullivan/dxf"
	"github.com/tobyjsullivan/dxf/table"
	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tobyjsullivan/dxf/color"
)

const (
	CUT_LINES = "Cut Lines"
	STITCH_LINES = "Stitch Lines"
	FOLD_LINES = "Fold Lines"
	GRAIN_LINES = "Grain Lines"
	NOTATIONS = "Notations"
)

type Pattern interface {
	CutLines() []geometry.Line
	FoldLines() []geometry.Line
	StitchLines() []geometry.Line
	GrainLines() []geometry.Line
	Notations() []geometry.Drawable
}

func findOrCreateLayer(d *drawing.Drawing, name string, cl color.ColorNumber, lt *table.LineType) error {
	layer, _ := d.Layer(name, true)
	if layer == nil {

		// Check if linetype exists
		existingLType, _ := d.LineType(lt.Name())
		if existingLType == nil {
			d.Sections[drawing.TABLES].(table.Tables)[table.LTYPE].Add(lt)
		}

		if _, err := d.AddLayer(name, cl, lt, true); err != nil {
			return err
		}
	}

	return nil
}

func DrawDXF(p Pattern, d *drawing.Drawing) error {
	var err error
	err = findOrCreateLayer(d, CUT_LINES, dxf.DefaultColor, dxf.DefaultLineType)
	if err != nil {
		return err
	}

	for _, line := range p.CutLines() {
		err := line.DrawDXF(d)
		if err != nil {
			return err
		}
	}


	err = findOrCreateLayer(d, STITCH_LINES, dxf.DefaultColor, table.NewLineType("Dotted", "Dot . . . .", 0.2, -0.1))
	if err != nil {
		return err
	}

	for _, line := range p.StitchLines() {
		err := line.DrawDXF(d)
		if err != nil {
			return err
		}
	}

	err = findOrCreateLayer(d, FOLD_LINES, dxf.DefaultColor, table.LT_DASHDOT)
	if err != nil {
		return err
	}

	for _, line := range p.FoldLines() {
		err := line.DrawDXF(d)
		if err != nil {
			return err
		}
	}

	err = findOrCreateLayer(d, GRAIN_LINES, color.Red, dxf.DefaultLineType)
	if err != nil {
		return err
	}

	for _, line := range p.GrainLines() {
		err := line.DrawDXF(d)
		if err != nil {
			return err
		}
	}

	err = findOrCreateLayer(d, NOTATIONS, dxf.DefaultColor, dxf.DefaultLineType)
	if err != nil {
		return err
	}

	for _, label := range p.Notations() {
		err := label.DrawDXF(d)
		if err != nil {
			return err
		}
	}

	return nil
}
