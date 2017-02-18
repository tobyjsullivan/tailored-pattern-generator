package patternfile

import (
	"errors"
	"fmt"

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

	PIECE_MARGIN = 5.00
	MAX_WIDTH = 90.0
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
		err = d.findOrCreateLayer(LAYER_STITCH, dxf.DefaultColor, table.NewLineType("Dotted", "Dot . . . .", 0.1, -0.5))
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

		if details.Measurements != nil {
			lines = append(lines, []string{
				fmt.Sprintf("Height: %.1f cm", details.Measurements.Height),
				fmt.Sprintf("Neck: %.1f cm", details.Measurements.NeckCircumference),
				fmt.Sprintf("Chest: %.1f cm", details.Measurements.ChestCircumference),
				fmt.Sprintf("Waist: %.1f cm", details.Measurements.WaistCircumference),
				fmt.Sprintf("Hip: %.1f cm", details.Measurements.HipCircumference),
				fmt.Sprintf("Sleeve Length: %.1f cm", details.Measurements.SleeveLength),
			}...)
		}

		// Write style details just above the origin mark
		detailPosition := &geometry.Point {
			X: 0.0,
			Y: float64(len(lines)) * LINE_SPACING + 2.0,
		}
		err = pf.drawMultilineText(lines, detailPosition)
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

	// Draw each piece
	cornerX, cornerY := 0.0, 0.0
	j := 0
	for i := 0; i < len(piecesOnFold); i++ {
		// Draw first piece on fold
		p := piecesOnFold[i]

		pf.drawPiece(p, cornerX, cornerY)

		bbox := pieceBoundingBox(p)
		rowMaxHeight := bbox.Height()
		cornerX += bbox.Width() + PIECE_MARGIN

		for ; cornerX < MAX_WIDTH && j < len(piecesOffFold); j++ {
			p = piecesOffFold[j]
			bbox = pieceBoundingBox(p)

			if cornerX + bbox.Width() > MAX_WIDTH {
				break
			}

			pf.drawPiece(p, cornerX, cornerY)

			cornerX += bbox.Width() + PIECE_MARGIN
			height := bbox.Height()
			if height > rowMaxHeight {
				rowMaxHeight = height
			}
		}

		cornerX = 0.0
		cornerY -= rowMaxHeight + PIECE_MARGIN
	}

	rowMaxHeight := 0.0
	for ; cornerX < MAX_WIDTH && j < len(piecesOffFold); j++ {
		p := piecesOffFold[j]

		pf.drawPiece(p, cornerX, cornerY)

		bbox := pieceBoundingBox(p)
		cornerX += bbox.Width() + PIECE_MARGIN
		height := bbox.Height()
		if height > rowMaxHeight {
			rowMaxHeight = height
		}

		if cornerX > MAX_WIDTH {
			cornerX = 0.0
			cornerY -= bbox.Height() + PIECE_MARGIN
			rowMaxHeight = 0.0
		}
	}

	return nil
}

func (pf *PatternFile) drawMultilineText(lines []string, pos *geometry.Point) error {
	detailsBlk := &geometry.Block{}
	for _, l := range lines {
		detailsBlk.AddText(&geometry.Text{
			Content: l,
			Position: pos,
		})

		pos = pos.Move(0, -LINE_SPACING)
	}

	return pf.DrawBlock(detailsBlk, &geometry.Point{})
}

func pieceBoundingBox(p pieces.Piece) *geometry.BoundingBox {
	cl := p.CutLayer()
	sl := p.StitchLayer()
	nl := p.NotationLayer()

	return geometry.CollectiveBoundingBox(cl, sl, nl)
}

func (pf *PatternFile) drawPiece(p pieces.Piece, cornerX, cornerY float64) error {
	bbox := pieceBoundingBox(p)

	pieceOffset := &geometry.Point{
		X: cornerX - bbox.Left,
		Y: cornerY - bbox.Top,
	}

	err := pf.SetLayer(LAYER_CUT)
	if err != nil {
		return err
	}

	err = pf.DrawBlock(p.CutLayer(), pieceOffset)
	if err != nil {
		return err
	}

	err = pf.SetLayer(LAYER_STITCH)
	if err != nil {
		return err
	}

	err = pf.DrawBlock(p.StitchLayer(), pieceOffset)
	if err != nil {
		return err
	}

	err = pf.SetLayer(LAYER_NOTATIONS)
	if err != nil {
		return err
	}

	err = pf.DrawBlock(p.NotationLayer(), pieceOffset)
	if err != nil {
		return err
	}

	// Stamp piece details
	pieceCorner := &geometry.Point{X: cornerX, Y: cornerY}
	pieceCentre := pieceCorner.MidpointTo(pieceCorner.Move(bbox.Width(), -bbox.Height()))
	details := p.Details()
	lines := []string {
		fmt.Sprintf("PN: %s", details.PieceNumber),
		details.Description,
	}
	err = pf.drawMultilineText(lines, pieceCentre)

	return err
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
	if t.Rotation != nil {
		text.Rotate = t.Rotation.Degrees()
	}
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
