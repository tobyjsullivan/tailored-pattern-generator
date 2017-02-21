package patternfile

import (
	"errors"
	"fmt"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/styles"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/drawing"
)

const (
	LAYER_CUT       = "CUT LINES"
	LAYER_STITCH    = "STITCH LINES"
	LAYER_NOTATIONS = "NOTATIONS"

	LINE_SPACING = 1.4

	PIECE_MARGIN = 5.00
	MAX_WIDTH = 90.0
)

type PatternFile struct {
	Style styles.Style
	dxf drawing.Drawing
}

func (pf *PatternFile) SaveDXF(filepath string) error {
	pf.dxf = drawing.NewDXF(MAX_WIDTH)
	err := pf.DrawPattern(pf.Style)
	if err != nil {
		return err
	}

	return pf.dxf.SaveAs(filepath)
}

func (pf *PatternFile) SavePDF(filepath string) error {
	pf.dxf = drawing.NewPDF(MAX_WIDTH)
	err := pf.DrawPattern(pf.Style)
	if err != nil {
		return err
	}

	return pf.dxf.SaveAs(filepath)
}

func (d *PatternFile) SetLayer(layer string) error {
	var err error

	switch layer {
	case LAYER_CUT:
		err = d.dxf.SetLayer(drawing.LAYER_NORMAL)
	case LAYER_STITCH:
		err = d.dxf.SetLayer(drawing.LAYER_ALT1)
	case LAYER_NOTATIONS:
		err = d.dxf.SetLayer(drawing.LAYER_ALT2)
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
	return d.dxf.StraightLine(l.Start, l.End)
}

func (d *PatternFile) drawPoint(p *geometry.Point) error {
	err := d.dxf.StraightLine(
		&geometry.Point{
			X: p.X-0.5,
			Y: p.Y,
		},
		&geometry.Point{
			X: p.X+0.5,
			Y: p.Y,
		},
	)
	if err != nil {
		return err
	}

	return d.dxf.StraightLine(
		&geometry.Point{
			X: p.X,
			Y: p.Y-0.5,
		},
		&geometry.Point{
			X: p.X,
			Y: p.Y+0.5,
		},
	)
}

func (d *PatternFile) drawText(t *geometry.Text) error {
	return d.dxf.Text(t.Position, t.Content, t.Rotation)
}
