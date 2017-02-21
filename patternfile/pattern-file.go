package patternfile

import (
	"errors"
	"fmt"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/styles"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/drawing"
	"time"
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
}

func (pf *PatternFile) SaveDXF(filepath string) error {
	dxf := drawing.NewDXF(MAX_WIDTH)
	err := pf.DrawPattern(dxf, pf.Style)
	if err != nil {
		return err
	}

	return dxf.SaveAs(filepath)
}

func (pf *PatternFile) SavePDF(filepath string) error {
	pdf := drawing.NewPDF(MAX_WIDTH)
	err := pf.DrawPattern(pdf, pf.Style)
	if err != nil {
		return err
	}

	return pdf.SaveAs(filepath)
}

func (pf *PatternFile) SetLayer(d drawing.Drawing, layer string) error {
	var err error

	switch layer {
	case LAYER_CUT:
		err = d.SetLayer(drawing.LAYER_NORMAL)
	case LAYER_STITCH:
		err = d.SetLayer(drawing.LAYER_ALT1)
	case LAYER_NOTATIONS:
		err = d.SetLayer(drawing.LAYER_ALT2)
	default:
		err = errors.New("The requested layer does not exist")
	}

	return err
}

func (pf *PatternFile) DrawPattern(d drawing.Drawing, s styles.Style) error {
	var err error

	err = pf.SetLayer(d, LAYER_CUT)
	if err != nil {
		return err
	}
	startingY := 0.0
	details := s.Details()
	if details != nil {
		lines := []string{
			fmt.Sprintf("Style Number: %s", details.StyleNumber),
			details.Description,
			fmt.Sprintf("Generated: %s", time.Now().Format("2006-01-02 15:04 MST")),
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
			Y: -1.0,
		}
		err = pf.drawMultilineText(d, lines, detailPosition)
		if err != nil {
			return err
		}

		startingY = -(float64(len(lines)) * LINE_SPACING + 2.0)
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
	cornerX, cornerY := 0.0, startingY
	j := 0
	for i := 0; i < len(piecesOnFold); i++ {
		// Draw first piece on fold
		p := piecesOnFold[i]

		pf.drawPiece(d, p, cornerX, cornerY)

		bbox := pieceBoundingBox(p)
		rowMaxHeight := bbox.Height()
		cornerX += bbox.Width() + PIECE_MARGIN

		for ; cornerX < MAX_WIDTH && j < len(piecesOffFold); j++ {
			p = piecesOffFold[j]
			bbox = pieceBoundingBox(p)

			if cornerX + bbox.Width() > MAX_WIDTH {
				break
			}

			pf.drawPiece(d, p, cornerX, cornerY)

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

		pf.drawPiece(d, p, cornerX, cornerY)

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

func (pf *PatternFile) drawMultilineText(d drawing.Drawing, lines []string, pos *geometry.Point) error {
	detailsBlk := &geometry.Block{}
	for _, l := range lines {
		detailsBlk.AddText(&geometry.Text{
			Content: l,
			Position: pos,
		})

		pos = pos.Move(0, -LINE_SPACING)
	}

	return pf.DrawBlock(d, detailsBlk, &geometry.Point{})
}

func pieceBoundingBox(p pieces.Piece) *geometry.BoundingBox {
	cl := p.CutLayer()
	sl := p.StitchLayer()
	nl := p.NotationLayer()

	return geometry.CollectiveBoundingBox(cl, sl, nl)
}

func (pf *PatternFile) drawPiece(d drawing.Drawing, p pieces.Piece, cornerX, cornerY float64) error {
	bbox := pieceBoundingBox(p)

	pieceOffset := &geometry.Point{
		X: cornerX - bbox.Left,
		Y: cornerY - bbox.Top,
	}

	err := pf.SetLayer(d, LAYER_CUT)
	if err != nil {
		return err
	}

	err = pf.DrawBlock(d, p.CutLayer(), pieceOffset)
	if err != nil {
		return err
	}

	err = pf.SetLayer(d, LAYER_STITCH)
	if err != nil {
		return err
	}

	err = pf.DrawBlock(d, p.StitchLayer(), pieceOffset)
	if err != nil {
		return err
	}

	err = pf.SetLayer(d, LAYER_NOTATIONS)
	if err != nil {
		return err
	}

	err = pf.DrawBlock(d, p.NotationLayer(), pieceOffset)
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
	err = pf.drawMultilineText(d, lines, pieceCentre)

	return err
}

func (pf *PatternFile) DrawBlock(d drawing.Drawing, b *geometry.Block, offset *geometry.Point) error {
	movedBlk := b.Move(offset.X, offset.Y)

	var err error
	for _, l := range movedBlk.StraightLines {
		err = d.StraightLine(l.Start, l.End)
		if err != nil {
			return err
		}
	}

	for _, p := range movedBlk.Points {
		err = pf.drawPoint(d, p)
		if err != nil {
			return err
		}
	}

	for _, t := range movedBlk.Text {
		err = d.Text(t.Position, t.Content, t.Rotation)
		if err != nil {
			return err
		}
	}

	for _, block := range movedBlk.Blocks {
		err = pf.DrawBlock(d, block, offset)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pf *PatternFile) drawPoint(d drawing.Drawing, p *geometry.Point) error {
	err := d.StraightLine(
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

	return d.StraightLine(
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
