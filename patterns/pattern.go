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
	FOLD_LINES = "Fold Lines"
	GRAIN_LINES = "Grain Lines"
	NOTATIONS = "Notations"
)

type Pattern interface {
	GetPoints() map[string]*geometry.Point
	GetCutLines() []geometry.Line
	GetFoldLines() []geometry.Line
	GetGrainLines() []geometry.Line
	GetLabels() []geometry.Drawable
}

func findOrCreateLayer(d *drawing.Drawing, name string, cl color.ColorNumber, lt *table.LineType) error {
	layer, _ := d.Layer(name, true)
	if layer == nil {
		_, err := d.AddLayer(name, cl, lt, true)
		if err != nil {
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

	for label, point := range p.GetPoints() {
		err := point.DrawDXF(label, d)
		if err != nil {
			return err
		}
	}

	for _, line := range p.GetCutLines() {
		err := line.DrawDXF(d)
		if err != nil {
			return err
		}
	}

	err = findOrCreateLayer(d, FOLD_LINES, dxf.DefaultColor, table.LT_DASHDOT)
	if err != nil {
		return err
	}

	for _, line := range p.GetFoldLines() {
		err := line.DrawDXF(d)
		if err != nil {
			return err
		}
	}

	err = findOrCreateLayer(d, GRAIN_LINES, color.Red, dxf.DefaultLineType)
	if err != nil {
		return err
	}

	for _, line := range p.GetGrainLines() {
		err := line.DrawDXF(d)
		if err != nil {
			return err
		}
	}

	err = findOrCreateLayer(d, NOTATIONS, dxf.DefaultColor, dxf.DefaultLineType)
	if err != nil {
		return err
	}

	for _, label := range p.GetLabels() {
		err := label.DrawDXF(d)
		if err != nil {
			return err
		}
	}

	return nil
}
