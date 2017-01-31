package patterns

import (
	"github.com/tobyjsullivan/dxf/drawing"
	"github.com/tobyjsullivan/dxf"
	"github.com/tobyjsullivan/dxf/table"
	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tobyjsullivan/dxf/color"
)

type Pattern interface {
	GetPoints() map[string]*geometry.Point
	GetCutLines() []geometry.Line
	GetFoldLines() []geometry.Line
	GetGrainLines() []geometry.Line
	GetLabels() []geometry.Drawable
}

func DrawDXF(p Pattern, d *drawing.Drawing) error {
	var err error
	_, err = d.AddLayer("Cut Lines", dxf.DefaultColor, dxf.DefaultLineType, true)
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

	_, err = d.AddLayer("Fold Lines", dxf.DefaultColor, table.LT_DASHDOT, true)
	if err != nil {
		return err
	}

	for _, line := range p.GetFoldLines() {
		err := line.DrawDXF(d)
		if err != nil {
			return err
		}
	}

	_, err = d.AddLayer("Grain Lines", color.Red, dxf.DefaultLineType, true)
	if err != nil {
		return err
	}

	for _, line := range p.GetGrainLines() {
		err := line.DrawDXF(d)
		if err != nil {
			return err
		}
	}

	_, err = d.AddLayer("Notations", dxf.DefaultColor, dxf.DefaultLineType, true)
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
