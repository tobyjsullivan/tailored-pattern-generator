package rendering

import (
	"github.com/tailored-style/pattern-generator/styles"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/drawing"
	"github.com/tailored-style/pattern-generator/nesting"
	"math"
	"fmt"
	"reflect"
)

const (
	MARKER_PAGE_WIDTH = 152.4 // 60 inches
	MARKER_PAGE_MAX_HEIGHT = 1000.0 // 10 metres
)

type Marker struct {
	Style styles.Style
}

func (m *Marker) SavePDF(filepath string) error {
	width := MARKER_PAGE_WIDTH - (2.0 * drawing.PDF_PAGE_MARGIN)
	placements := m.layoutPieces(width)

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

	// Create PDF with height
	pageHeight := math.Abs(bottom) + (2.0 * drawing.PDF_PAGE_MARGIN)
	drawing := drawing.NewPDF(MARKER_PAGE_WIDTH, pageHeight)
	m.drawPieces(drawing, placements)

	return drawing.SaveAs(filepath)
}

type piecePlacement struct {
	pieces.Piece
	Position *geometry.Point
}

func (m *Marker) openPieces() []pieces.Piece {
	originals := m.Style.Pieces()
	openedPieces := make([]pieces.Piece, 0, len(originals))

	for _, p := range originals {
		if p.OnFold() {
			p = &OpenOnFold{Piece: p}
		}

		if p.Mirrored() {
			p = &MirroredPiece{Piece: p}
		}

		for i := 0; i < p.CutCount(); i++ {
			openedPieces = append(openedPieces, p)
		}

	}

	return openedPieces
}

func (m *Marker) layoutPieces(width float64) []*piecePlacement {
	openedPieces := m.openPieces()

	// Compute nesting layout
	items := make(map[int]*nesting.Rectangle)
	for i, p := range openedPieces {
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
	out := make([]*piecePlacement, len(openedPieces))
	for i, p := range openedPieces {
		pos := placements[i]
		out[i] = &piecePlacement{
			Piece: p,
			Position: pos,
		}
	}

	return out
}

func (m *Marker) drawPieces(d drawing.Drawing, placements []*piecePlacement) {
	// Draw each piece
	for _, placement := range placements {
		pos := placement.Position
		drawPiece(d, placement.Piece, pos.X, -pos.Y)
	}
}

func drawPiece(d drawing.Drawing, p pieces.Piece, cornerX, cornerY float64) error {
	bbox := pieces.BoundingBox(p)

	pieceOffset := &geometry.Point{
		X: cornerX - bbox.Left,
		Y: cornerY - bbox.Top,
	}

	err := DrawBlock(d, p.CutLayer(), pieceOffset)
	if err != nil {
		return err
	}

	return err
}

func DrawBlock(d drawing.Drawing, b *geometry.Block, offset *geometry.Point) error {
	movedBlk := b.Move(offset.X, offset.Y)

	var err error
	for _, l := range movedBlk.StraightLines {
		err = drawStraightLine(d, l)
		if err != nil {
			return err
		}
	}

	for _, p := range movedBlk.Points {
		err = drawPoint(d, p)
		if err != nil {
			return err
		}
	}

	for _, t := range movedBlk.Text {
		err = drawText(d, t)
		if err != nil {
			return err
		}
	}

	for _, block := range movedBlk.Blocks {
		err = DrawBlock(d, block, &geometry.Point{X: 0.0, Y: 0.0})
		if err != nil {
			return err
		}
	}

	return nil
}

func drawStraightLine(d drawing.Drawing, l *geometry.StraightLine) error {
	return d.StraightLine(l.Start, l.End)
}

func drawPoint(pdf drawing.Drawing, p *geometry.Point) error {
	err := pdf.StraightLine(
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

	err = pdf.StraightLine(
		&geometry.Point{
			X: p.X,
			Y: p.Y-0.5,
		},
		&geometry.Point{
			X: p.X,
			Y: p.Y+0.5,
		},
	)

	return err
}

func drawText(d drawing.Drawing, t *geometry.Text) error {
	return d.Text(t.Position, t.Content, t.Rotation)
}
