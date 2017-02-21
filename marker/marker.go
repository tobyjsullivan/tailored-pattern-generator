package marker

import (
	"github.com/tailored-style/pattern-generator/styles"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/drawing"
)

const (
	PAGE_WIDTH = 152.4 // 60 inches
	PIECE_MARGIN = 4.0
)

type Marker struct {
	Style styles.Style
}

func (m *Marker) SavePDF(filepath string) error {
	drawing := drawing.NewPDF(PAGE_WIDTH)

	m.drawPieces(drawing)

	return drawing.SaveAs(filepath)
}

func (m *Marker) drawPieces(d drawing.Drawing) {
	// Pieces on fold should be most left.
	piecesOnFold := []pieces.Piece{}
	piecesOffFold := []pieces.Piece{}
	for _, p := range m.Style.Pieces() {
		if p.OnFold() {
			piecesOnFold = append(piecesOnFold, p)
		} else {
			piecesOffFold = append(piecesOffFold, p)
		}
	}

	// Draw each piece
	maxWidth := d.DrawableWidth()
	cornerX, cornerY := 0.0, 0.0
	j := 0
	for i := 0; i < len(piecesOnFold); i++ {
		// Draw first piece on fold
		p := piecesOnFold[i]

		drawPiece(d, p, cornerX, cornerY)

		bbox := pieceBoundingBox(p)
		rowMaxHeight := bbox.Height()
		cornerX += bbox.Width() + PIECE_MARGIN

		for ; cornerX < maxWidth && j < len(piecesOffFold); j++ {
			p = piecesOffFold[j]
			bbox = pieceBoundingBox(p)

			if cornerX + bbox.Width() > maxWidth {
				break
			}

			drawPiece(d, p, cornerX, cornerY)

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
	for ; cornerX < maxWidth && j < len(piecesOffFold); j++ {
		p := piecesOffFold[j]

		drawPiece(d, p, cornerX, cornerY)

		bbox := pieceBoundingBox(p)
		cornerX += bbox.Width() + PIECE_MARGIN
		height := bbox.Height()
		if height > rowMaxHeight {
			rowMaxHeight = height
		}

		if cornerX > maxWidth {
			cornerX = 0.0
			cornerY -= bbox.Height() + PIECE_MARGIN
			rowMaxHeight = 0.0
		}
	}
}

func drawPiece(d drawing.Drawing, p pieces.Piece, cornerX, cornerY float64) error {
	bbox := pieceBoundingBox(p)

	pieceOffset := &geometry.Point{
		X: cornerX - bbox.Left,
		Y: cornerY - bbox.Top,
	}

	err := DrawBlock(d, p.CutLayer(), pieceOffset)
	if err != nil {
		return err
	}

	err = DrawBlock(d, p.StitchLayer(), pieceOffset)
	if err != nil {
		return err
	}

	err = DrawBlock(d, p.NotationLayer(), pieceOffset)
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
		err = DrawBlock(d, block, offset)
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


func pieceBoundingBox(p pieces.Piece) *geometry.BoundingBox {
	cl := p.CutLayer()
	sl := p.StitchLayer()
	nl := p.NotationLayer()

	return geometry.CollectiveBoundingBox(cl, sl, nl)
}
