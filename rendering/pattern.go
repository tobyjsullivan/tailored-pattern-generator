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
	"github.com/tailored-style/pattern-generator/nesting"
	"reflect"
)

const (
	LAYER_CUT       = "CUT LINES"
	LAYER_STITCH    = "STITCH LINES"
	LAYER_NOTATIONS = "NOTATIONS"

	LINE_SPACING = 1.4

	PATTERN_PIECE_MARGIN = 1.0
	PATTERN_PAGE_WIDTH = 91.44 // 36"
	PATTERN_PAGE_HEIGHT = 182.88 // 72 inches
	PATTERN_PAGE_MAX_HEIGHT = 1000.0 // 10 metres
)

type Pattern struct {
	Style styles.Style
}

func (pf *Pattern) SaveDXF(filepath string) error {
	placements := pf.layoutPieces(pf.Style.Pieces(), PATTERN_PAGE_WIDTH - (2.0 * drawing.PDF_PAGE_MARGIN))

	dxf := drawing.NewDXF(PATTERN_PAGE_WIDTH)
	err := pf.DrawPattern(dxf, pf.Style, placements, 0.0)
	if err != nil {
		return err
	}

	return dxf.SaveAs(filepath)
}

func (pf *Pattern) SavePDF(filepath string) error {
	placements := pf.layoutPieces(pf.Style.Pieces(), PATTERN_PAGE_WIDTH - (2.0 * drawing.PDF_PAGE_MARGIN))

	// Get height of all layouts
	bottom := 0.0
	for _, p := range placements {
		if p.Position == nil {
			panic(fmt.Sprintf("Failed to place a piece! %v", reflect.ValueOf(p.Piece)))
		}

		itemTop := p.Position.Y
		itemBottom := pieces.BoundingBox(p.Piece).Height() + itemTop

		if itemBottom > bottom {
			bottom = itemBottom
		}
	}

	estHeaderSize := 15.0
	pdf := drawing.NewPDF(PATTERN_PAGE_WIDTH, math.Abs(bottom) + estHeaderSize)

	var err error
	h, err := pf.drawHeaderInfo(pdf, pf.Style)
	err = pf.DrawPattern(pdf, pf.Style, placements, h)
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

func (pf *Pattern) drawHeaderInfo(d drawing.Drawing, s styles.Style) (height float64, err error) {
	err = pf.SetLayer(d, LAYER_CUT)
	if err != nil {
		return
	}
	height = 0.0
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
			return
		}

		height = -(float64(len(detailText)) * LINE_SPACING + 2.0)

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
				return
			}

			height = math.Min(height, -(float64(len(measurementText)) * LINE_SPACING + 2.0))
		}
	}

	return
}

func (pf *Pattern) DrawPattern(d drawing.Drawing, s styles.Style, placements []*piecePlacement, offsetY float64) error {
	for _, placement := range placements {
		pos := placement.Position
		err := pf.drawPiece(d, placement.Piece, pos.X, offsetY - pos.Y)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pf *Pattern) layoutPieces(pcs []pieces.Piece, width float64) []*piecePlacement {
	// Compute nesting layout
	items := make(map[int]*nesting.Rectangle)
	for i, p := range pcs {
		bbox := pieces.BoundingBox(p)
		items[i] = &nesting.Rectangle{
			Width: bbox.Width(),
			Height: bbox.Height(),
		}
	}
	cont := &nesting.Container{
		Width: width,
		Height: MARKER_PAGE_MAX_HEIGHT,
	}
	placements := cont.Pack(items)

	// Copy all pieces into output slice
	out := make([]*piecePlacement, len(pcs))
	for i, p := range pcs {
		pos := placements[i]
		out[i] = &piecePlacement{
			Piece: p,
			Position: pos,
		}
	}

	return out
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

	return pf.DrawBlock(d, detailsBlk)
}

func (pf *Pattern) drawPiece(d drawing.Drawing, p pieces.Piece, cornerX, cornerY float64) error {
	bbox := pieces.BoundingBox(p)

	offsetX := cornerX - bbox.Left
	offsetY := cornerY - bbox.Top

	err := pf.SetLayer(d, LAYER_CUT)
	if err != nil {
		return err
	}

	outer := p.OuterCut().Move(offsetX, offsetY)
	err = pf.drawPolyline(d, outer)
	if err != nil {
		return err
	}

	inner := p.InnerCut().Move(offsetX, offsetY)
	err = pf.DrawBlock(d, inner)
	if err != nil {
		return err
	}

	err = pf.SetLayer(d, LAYER_STITCH)
	if err != nil {
		return err
	}

	ink := p.Ink().Move(offsetX, offsetY)
	err = pf.DrawBlock(d, ink)
	if err != nil {
		return err
	}

	stitch := p.Stitch().Move(offsetX, offsetY)
	err = pf.DrawBlock(d, stitch)
	if err != nil {
		return err
	}

	err = pf.SetLayer(d, LAYER_NOTATIONS)
	if err != nil {
		return err
	}

	ref := p.Reference().Move(offsetX, offsetY)
	err = pf.DrawBlock(d, ref)
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

func (pf *Pattern) drawPolyline(d drawing.Drawing, poly *geometry.Polyline) error {
	var err error
	for _, l := range poly.StraightLines() {
		err = d.StraightLine(l.Start, l.End)
		if err != nil {
			return err
		}
	}

	return nil

}

func (pf *Pattern) DrawBlock(d drawing.Drawing, b *geometry.Block) error {
	var err error
	for _, l := range b.StraightLines {
		err = d.StraightLine(l.Start, l.End)
		if err != nil {
			return err
		}
	}

	for _, p := range b.Points {
		err = pf.drawPoint(d, p)
		if err != nil {
			return err
		}
	}

	for _, t := range b.Text {
		err = d.Text(t.Position, t.Content, t.Rotation)
		if err != nil {
			return err
		}
	}

	for _, block := range b.Blocks {
		err = pf.DrawBlock(d, block)
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
