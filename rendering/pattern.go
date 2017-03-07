package rendering

import (
	"errors"
	"fmt"

	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/styles"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/drawing"
	"time"
	"math"
)

const (
	LAYER_CUT       = "CUT LINES"
	LAYER_STITCH    = "STITCH LINES"
	LAYER_NOTATIONS = "NOTATIONS"

	LINE_SPACING = 1.4

	PATTERN_PIECE_MARGIN = 1.0
	PATTERN_PAGE_WIDTH = 91.44 // 36"
	PATTERN_PAGE_HEIGHT = 182.88 // 72 inches
)

type Pattern struct {
	Style styles.Style
}

func (pf *Pattern) SaveDXF(filepath string) error {
	dxf := drawing.NewDXF(PATTERN_PAGE_WIDTH)
	err := pf.DrawPattern(dxf, pf.Style)
	if err != nil {
		return err
	}

	return dxf.SaveAs(filepath)
}

func (pf *Pattern) SavePDF(filepath string) error {
	pdf := drawing.NewPDF(PATTERN_PAGE_WIDTH, PATTERN_PAGE_HEIGHT)
	err := pf.DrawPattern(pdf, pf.Style)
	if err != nil {
		return err
	}

	return pdf.SaveAs(filepath)
}

func (pf *Pattern) SetLayer(d drawing.Drawing, layer string) error {
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

func (pf *Pattern) DrawPattern(d drawing.Drawing, s styles.Style) error {
	var err error

	err = pf.SetLayer(d, LAYER_CUT)
	if err != nil {
		return err
	}
	startingY := 0.0
	details := s.Details()
	if details != nil {
		detailText := []string{
			fmt.Sprintf("Style Number: %s", details.StyleNumber),
			details.Description,
			fmt.Sprintf("Generated: %s", time.Now().Format("2006-01-02 15:04 MST")),
		}

		detailPosition := &geometry.Point {
			X: 0.0,
			Y: -1.0,
		}
		err = pf.drawMultilineText(d, detailText, detailPosition)
		if err != nil {
			return err
		}

		startingY = -(float64(len(detailText)) * LINE_SPACING + 2.0)

		if details.Measurements != nil {
			measurementText := []string{
				fmt.Sprintf("Height: %.1f cm", details.Measurements.Height),
				fmt.Sprintf("Neck: %.1f cm", details.Measurements.NeckCircumference),
				fmt.Sprintf("Chest: %.1f cm", details.Measurements.ChestCircumference),
				fmt.Sprintf("Waist: %.1f cm", details.Measurements.WaistCircumference),
				fmt.Sprintf("Hip: %.1f cm", details.Measurements.HipCircumference),
				fmt.Sprintf("Sleeve Length: %.1f cm", details.Measurements.SleeveLength),
				fmt.Sprintf("Wrist: %.1f cm", details.Measurements.WristCircumference),
			}
			measurementPosition := &geometry.Point {
				X: d.DrawableWidth() / 2.0,
				Y: -1.0,
			}
			err = pf.drawMultilineText(d, measurementText, measurementPosition)
			if err != nil {
				return err
			}

			startingY = math.Min(startingY, -(float64(len(measurementText)) * LINE_SPACING + 2.0))
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
	cornerX, cornerY := 0.0, startingY
	drawableWidth := d.DrawableWidth()
	j := 0
	for i := 0; i < len(piecesOnFold); i++ {
		// Draw first piece on fold
		p := piecesOnFold[i]
		fmt.Printf("Want to draw %v\n", p)

		fmt.Printf("Drawing %v\n", p)
		pf.drawPiece(d, p, cornerX, cornerY)

		bbox := pieces.BoundingBox(p)
		rowMaxHeight := bbox.Height()
		cornerX += bbox.Width() + PATTERN_PIECE_MARGIN


		for ; cornerX < drawableWidth && j < len(piecesOffFold); j++ {
			p = piecesOffFold[j]
			bbox = pieces.BoundingBox(p)

			if cornerX + bbox.Width() > drawableWidth {
				break
			}

			pf.drawPiece(d, p, cornerX, cornerY)

			cornerX += bbox.Width() + PATTERN_PIECE_MARGIN
			height := bbox.Height()
			if height > rowMaxHeight {
				rowMaxHeight = height
			}
		}

		cornerX = 0.0
		cornerY -= rowMaxHeight + PATTERN_PIECE_MARGIN
	}

	fmt.Println("Done drawing all pieces on fold")
	rowMaxHeight := 0.0
	for ; cornerX < drawableWidth && j < len(piecesOffFold); j++ {
		p := piecesOffFold[j]
		fmt.Printf("Want to draw %v at (%.2f, %.2f)\n", p, cornerX, cornerY)

		bbox := pieces.BoundingBox(p)
		if cornerX + bbox.Width() > drawableWidth {
			fmt.Println("Piece won't fit on this line")
			cornerX = 0.0
			cornerY -= rowMaxHeight + PATTERN_PIECE_MARGIN
			rowMaxHeight = 0.0
		}

		fmt.Printf("Drawing %v\n", p)
		pf.drawPiece(d, p, cornerX, cornerY)

		cornerX += bbox.Width() + PATTERN_PIECE_MARGIN
		height := bbox.Height()
		if height > rowMaxHeight {
			rowMaxHeight = height
		}
	}

	return nil
}

func (pf *Pattern) drawMultilineText(d drawing.Drawing, lines []string, pos *geometry.Point) error {
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

func (pf *Pattern) drawPiece(d drawing.Drawing, p pieces.Piece, cornerX, cornerY float64) error {
	bbox := pieces.BoundingBox(p)

	pieceOffset := &geometry.Point{
		X: cornerX - bbox.Left,
		Y: cornerY - bbox.Top,
	}

	err := pf.SetLayer(d, LAYER_CUT)
	if err != nil {
		return err
	}

	err = pf.DrawBlock(d, p.InnerCut(), pieceOffset)
	if err != nil {
		return err
	}

	err = pf.SetLayer(d, LAYER_STITCH)
	if err != nil {
		return err
	}

	err = pf.DrawBlock(d, p.Stitch(), pieceOffset)
	if err != nil {
		return err
	}

	err = pf.SetLayer(d, LAYER_NOTATIONS)
	if err != nil {
		return err
	}

	err = pf.DrawBlock(d, p.Ink(), pieceOffset)
	if err != nil {
		return err
	}

	// Stamp piece details
	pieceCorner := &geometry.Point{X: cornerX, Y: cornerY}
	pieceCentre := pieceCorner.MidpointTo(pieceCorner.Move(bbox.Width(), -bbox.Height()))
	count := p.CutCount()
	if p.Mirrored() {
		count *= 2
	}
	cutInstructions := fmt.Sprintf("Cut %d", count)
	if p.Mirrored() {
		cutInstructions += " mirrored"
	}
	if p.OnFold() {
		cutInstructions += " on fold"
	}
	details := p.Details()
	lines := []string {
		fmt.Sprintf("PN: %d", details.PieceNumber),
		details.Description,
		cutInstructions,
	}
	err = pf.drawMultilineText(d, lines, pieceCentre)

	return err
}

func (pf *Pattern) DrawBlock(d drawing.Drawing, b *geometry.Block, offset *geometry.Point) error {
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
		err = pf.DrawBlock(d, block, &geometry.Point{X: 0.0, Y: 0.0})
		if err != nil {
			return err
		}
	}

	return nil
}

func (pf *Pattern) drawPoint(d drawing.Drawing, p *geometry.Point) error {
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
