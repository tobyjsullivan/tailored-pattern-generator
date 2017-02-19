package marker

import (
	"github.com/tailored-style/pattern-generator/styles"
	"github.com/jung-kurt/gofpdf"
	"github.com/tailored-style/pattern-generator/pieces"
	"github.com/tailored-style/pattern-generator/geometry"
	"fmt"
)

const (
	PAGE_WIDTH = 152.4 // 60 inches
	PAGE_HEIGHT = 182.9 // 72 inches
	PAGE_MARGIN = 1.0
	PAGE_UNIT = "cm"
	PIECE_MARGIN = 4.0
)

type Marker struct {
	Style styles.Style
}

func (m *Marker) SavePDF(filepath string) error {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr: PAGE_UNIT,
		Size: gofpdf.SizeType{Wd: PAGE_WIDTH, Ht: PAGE_HEIGHT},
	})
	pdf.SetMargins(PAGE_MARGIN, PAGE_MARGIN, PAGE_MARGIN)
	pdf.SetFont("Courier", "", 48.0)
	pdf.AddPage()

	m.drawPieces(pdf)

	if err := pdf.Error(); err != nil {
		panic (err.Error())
	}

	return pdf.OutputFileAndClose(filepath)
}

func (m *Marker) drawPieces(pdf *gofpdf.Fpdf) {
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
	maxWidth := (PAGE_WIDTH - 2.0 * PAGE_MARGIN)
	cornerX, cornerY := 0.0, 0.0
	j := 0
	for i := 0; i < len(piecesOnFold); i++ {
		// Draw first piece on fold
		p := piecesOnFold[i]

		drawPiece(pdf, p, cornerX, cornerY)

		bbox := pieceBoundingBox(p)
		rowMaxHeight := bbox.Height()
		cornerX += bbox.Width() + PIECE_MARGIN

		for ; cornerX < maxWidth && j < len(piecesOffFold); j++ {
			p = piecesOffFold[j]
			bbox = pieceBoundingBox(p)

			if cornerX + bbox.Width() > maxWidth {
				break
			}

			drawPiece(pdf, p, cornerX, cornerY)

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

		drawPiece(pdf, p, cornerX, cornerY)

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

func drawPiece(pdf *gofpdf.Fpdf, p pieces.Piece, cornerX, cornerY float64) error {
	bbox := pieceBoundingBox(p)

	pieceOffset := &geometry.Point{
		X: cornerX - bbox.Left,
		Y: cornerY - bbox.Top,
	}

	err := DrawBlock(pdf, p.CutLayer(), pieceOffset)
	if err != nil {
		return err
	}

	err = DrawBlock(pdf, p.StitchLayer(), pieceOffset)
	if err != nil {
		return err
	}

	err = DrawBlock(pdf, p.NotationLayer(), pieceOffset)
	if err != nil {
		return err
	}

	return err
}

func DrawBlock(pdf *gofpdf.Fpdf, b *geometry.Block, offset *geometry.Point) error {
	movedBlk := b.Move(offset.X, offset.Y)

	var err error
	for _, l := range movedBlk.StraightLines {
		err = drawStraightLine(pdf, l)
		if err != nil {
			return err
		}
	}

	for _, p := range movedBlk.Points {
		err = drawPoint(pdf, p)
		if err != nil {
			return err
		}
	}

	for _, t := range movedBlk.Text {
		err = drawText(pdf, t)
		if err != nil {
			return err
		}
	}

	for _, block := range movedBlk.Blocks {
		err = DrawBlock(pdf, block, offset)
		if err != nil {
			return err
		}
	}

	return nil
}

func drawStraightLine(pdf *gofpdf.Fpdf, l *geometry.StraightLine) error {
	fmt.Printf("Drawing line %v\n", l)
	pdf.Line(x(l.Start.X), y(l.Start.Y), x(l.End.X), y(l.End.Y))

	if pdf.Err() {
		return pdf.Error()
	}

	return nil
}

func drawPoint(pdf *gofpdf.Fpdf, p *geometry.Point) error {
	pdf.Line(x(p.X-0.5), y(p.Y), x(p.X+0.5), y(p.Y))
	pdf.Line(x(p.X), y(p.Y-0.5), x(p.X), y(p.Y+0.5))

	if pdf.Err() {
		return pdf.Error()
	}

	return nil
}

func drawText(pdf *gofpdf.Fpdf, t *geometry.Text) error {
	rotated := t.Rotation != nil
	if rotated {
		pdf.TransformBegin()
		pdf.TransformRotate(t.Rotation.Degrees(), x(t.Position.X), y(t.Position.Y))
	}

	pdf.Text(x(t.Position.X), y(t.Position.Y), t.Content)

	if rotated {
		pdf.TransformEnd()
	}

	if pdf.Err() {
		return pdf.Error()
	}

	return nil
}


func pieceBoundingBox(p pieces.Piece) *geometry.BoundingBox {
	cl := p.CutLayer()
	sl := p.StitchLayer()
	nl := p.NotationLayer()

	return geometry.CollectiveBoundingBox(cl, sl, nl)
}

func x(x float64) float64 {
	return PAGE_MARGIN + x
}

func y(y float64) float64 {
	return PAGE_MARGIN + -y
}