package drawing

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tobyjsullivan/dxf"
	yofu_drawing "github.com/tobyjsullivan/dxf/drawing"
	"github.com/yofu/dxf/drawing"
	"github.com/tobyjsullivan/dxf/color"
	"github.com/tobyjsullivan/dxf/table"
)

const DXF_MARGIN = 1.0

type dxfDrawing struct {
	dxf *yofu_drawing.Drawing
	drawableWd float64
}

func NewDXF(widthCm float64) Drawing {
	dxf := dxf.NewDrawing()

	return &dxfDrawing{
		dxf: dxf,
		drawableWd: widthCm - 2.0 * DXF_MARGIN,
	}
}

func (d *dxfDrawing) DrawableWidth() float64 {
	return d.drawableWd
}

func (d *dxfDrawing) StraightLine(p0 *geometry.Point, p1 *geometry.Point) error {
	_, err := d.dxf.Line(p0.X, p0.Y, 0.0, p1.X, p1.Y, 0.0)
	return err
}

func (d *dxfDrawing) Text(position *geometry.Point, content string, rotation *geometry.Angle) error {
	text, err := d.dxf.Text(content, position.X, position.Y, 0.0, 1.0)
	if err != nil {
		return err
	}
	if rotation != nil {
		text.Rotate = rotation.Degrees()
	}
	return nil
}

func (d *dxfDrawing) SetLayer(layer string) error {
	switch layer {
	case LAYER_NORMAL:
		return d.findOrCreateLayer(LAYER_NORMAL, dxf.DefaultColor, dxf.DefaultLineType)
	case LAYER_ALT1:
		return d.findOrCreateLayer(LAYER_ALT1, dxf.DefaultColor, table.NewLineType("Dotted", "Dot . . . .", 0.1, -0.5))
	case LAYER_ALT2:
		return d.findOrCreateLayer(LAYER_ALT2, dxf.DefaultColor, table.LT_DASHDOT)
	default:
		panic("Unknown layer")
	}
}

func (d *dxfDrawing) SaveAs(filepath string) error {
	return d.dxf.SaveAs(filepath)
}

func (d *dxfDrawing) findOrCreateLayer(name string, cl color.ColorNumber, lt *table.LineType) error {
	layer, _ := d.dxf.Layer(name, true)
	if layer == nil {
		// Check if linetype exists
		existingLType, _ := d.dxf.LineType(lt.Name())
		if existingLType == nil {
			d.dxf.Sections[drawing.TABLES].(table.Tables)[table.LTYPE].Add(lt)
		}

		if _, err := d.dxf.AddLayer(name, cl, lt, true); err != nil {
			return err
		}
	}

	return nil
}