package patternfile

import (
	"errors"
	"fmt"
	"math"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/styles"
	"github.com/tobyjsullivan/dxf/color"
	"github.com/tobyjsullivan/dxf/drawing"
	"github.com/tobyjsullivan/dxf/table"
	"github.com/tobyjsullivan/dxf"
	"github.com/tailored-style/pattern-generator/pieces"
)

const (
	LAYER_CUT       = "CUT LINES"
	LAYER_STITCH    = "STITCH LINES"
	LAYER_FOLD      = "FOLD LINES"
	LAYER_GRAIN     = "GRAIN LINES"
	LAYER_NOTATIONS = "NOTATIONS"

	LINE_SPACING = 1.4
)

type PatternFile struct {
	drawing *drawing.Drawing
}

func NewPatternFile() *PatternFile {
	drawing := dxf.NewDrawing()
	drawing.Header().LtScale = 1.0

	return &PatternFile{
		drawing: drawing,
	}
}

func (pf *PatternFile) SaveAs(filepath string) error {
	return pf.drawing.SaveAs(filepath)
}

func (d *PatternFile) SetLayer(layer string) error {
	var err error

	switch layer {
	case LAYER_CUT:
		err = d.findOrCreateLayer(LAYER_CUT, dxf.DefaultColor, dxf.DefaultLineType)
	case LAYER_STITCH:
		err = d.findOrCreateLayer(LAYER_STITCH, dxf.DefaultColor, table.NewLineType("Dotted", "Dot . . . .", 0.2, -0.1))
	case LAYER_NOTATIONS:
		err = d.findOrCreateLayer(LAYER_NOTATIONS, dxf.DefaultColor, dxf.DefaultLineType)
	case LAYER_FOLD:
		err = d.findOrCreateLayer(LAYER_FOLD, dxf.DefaultColor, table.LT_DASHDOT)
	case LAYER_GRAIN:
		err = d.findOrCreateLayer(LAYER_GRAIN, color.Red, dxf.DefaultLineType)
	default:
		err = errors.New("The requested layer does not exist")
	}

	return err
}

func (pf *PatternFile) DrawPattern(s styles.Style) error {
	var err error

	err = pf.SetLayer(LAYER_CUT)
	if err != nil {
		return err
	}
	details := s.Details()
	if details != nil {
		lines := []string{
			fmt.Sprintf("Style Number: %s", details.StyleNumber),
			details.Description,
		}

		// Write style details just above the origin mark
		detailPosition := &geometry.Point {
			X: 0.0,
			Y: float64(len(lines)) * LINE_SPACING + 2.0,
		}
		detailsBlk := &geometry.Block{}
		for _, l := range lines {
			detailsBlk.AddText(&geometry.Text{
				Content: l,
				Position: detailPosition,
			})

			detailPosition = detailPosition.Move(0, -LINE_SPACING)
		}

		err = pf.DrawBlock(detailsBlk, &geometry.Point{})
		if err != nil {
			return err
		}
	}

	// Pieces on fold should be most left.
	piecesOnFold := []pieces.Piece{}
	piecesOffFold := []pieces.Piece{}
	for _, p := range s.Pieces() {
		if p.OnFold() {
			piecesOnFold = append(piecesOnFold, p)
		} else {
			piecesOffFold = append(piecesOffFold, p)
		}
	}
	allPieces := append(piecesOnFold, piecesOffFold...)

	// Draw each piece
	cornerX, cornerY := 0.0, 0.0
	for _, p := range allPieces {
		cl := p.CutLayer()
		sl := p.StitchLayer()
		nl := p.NotationLayer()

		bbox := geometry.CollectiveBoundingBox(cl, sl, nl)

		pieceOffset := &geometry.Point{
			X: cornerX - bbox.Left,
			Y: cornerY - bbox.Top,
		}

		err = pf.SetLayer(LAYER_CUT)
		if err != nil {
			return err
		}

		err = pf.DrawBlock(cl, pieceOffset)
		if err != nil {
			return err
		}

		err = pf.SetLayer(LAYER_STITCH)
		if err != nil {
			return err
		}

		err = pf.DrawBlock(sl, pieceOffset)
		if err != nil {
			return err
		}

		err = pf.SetLayer(LAYER_NOTATIONS)
		if err != nil {
			return err
		}

		err = pf.DrawBlock(nl, pieceOffset)
		if err != nil {
			return err
		}

		cornerX += bbox.Width() + 10.0
	}

	return nil
}

func (d *PatternFile) DrawBlock(b *geometry.Block, offset *geometry.Point) error {
	movedBlk := b.Move(offset.X, offset.Y)

	var err error
	for _, l := range movedBlk.StraightLines {
		err = d.drawStraightLine(l)
		if err != nil {
			return err
		}
	}

	for _, p := range movedBlk.Points {
		err = d.drawPoint(p)
		if err != nil {
			return err
		}
	}

	for _, t := range movedBlk.Text {
		err = d.drawText(t)
		if err != nil {
			return err
		}
	}

	for _, block := range movedBlk.Blocks {
		err = d.DrawBlock(block, offset)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *PatternFile) drawStraightLine(l *geometry.StraightLine) error {
	_, err := d.drawing.Line(l.Start.X, l.Start.Y, 0.0, l.End.X, l.End.Y, 0.0)
	return err
}

func (d *PatternFile) drawPoint(p *geometry.Point) error {
	_, err := d.drawing.Line(p.X-0.5, p.Y, 0.0, p.X+0.5, p.Y, 0.0)
	if err != nil {
		return err
	}

	_, err = d.drawing.Line(p.X, p.Y-0.5, 0.0, p.X, p.Y+0.5, 0.0)
	if err != nil {
		return err
	}

	return nil
}

func (d *PatternFile) drawText(t *geometry.Text) error {
	text, err := d.drawing.Text(t.Content, t.Position.X, t.Position.Y, 0.0, 1.0)
	if err != nil {
		return err
	}
	text.Rotate = 360.0 * (t.Rotation / (2.0 * math.Pi))
	return nil
}

func (d *PatternFile) findOrCreateLayer(name string, cl color.ColorNumber, lt *table.LineType) error {
	layer, _ := d.drawing.Layer(name, true)
	if layer == nil {

		// Check if linetype exists
		existingLType, _ := d.drawing.LineType(lt.Name())
		if existingLType == nil {
			d.drawing.Sections[drawing.TABLES].(table.Tables)[table.LTYPE].Add(lt)
		}

		if _, err := d.drawing.AddLayer(name, cl, lt, true); err != nil {
			return err
		}
	}

	return nil
}
