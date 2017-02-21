package drawing

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tailored-style/pattern-generator/geometry"
)

const (
	PDF_PAGE_MARGIN = 1.0
	PDF_PAGE_UNIT = "cm"
	PDF_PAGE_HEIGHT = 182.9 // 72 inches
)

type pdfDrawing struct {
	pdf         *gofpdf.Fpdf
	drawableWd  float64
	layerIds    map[string]int
	layerActive bool
}

func NewPDF(widthCm float64) Drawing {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr: PDF_PAGE_UNIT,
		Size: gofpdf.SizeType{Wd: widthCm, Ht: PDF_PAGE_HEIGHT},
	})
	pdf.SetMargins(PDF_PAGE_MARGIN, PDF_PAGE_MARGIN, PDF_PAGE_MARGIN)
	pdf.SetFont("Courier", "", 48.0)
	pdf.AddPage()

	return &pdfDrawing{
		pdf: pdf,
		drawableWd: widthCm - (2.0 * PDF_PAGE_MARGIN),
		layerIds: make(map[string]int),
	}
}

func (d *pdfDrawing) DrawableWidth() float64 {
	return d.drawableWd
}

func (d *pdfDrawing) StraightLine(start *geometry.Point, end *geometry.Point) error {
	d.pdf.Line(x(start.X), y(start.Y), x(end.X), y(end.Y))

	if d.pdf.Err() {
		return d.pdf.Error()
	}

	return nil
}

func (d *pdfDrawing) Text(position *geometry.Point, content string, rotation *geometry.Angle) error {
	if rotation != nil {
		d.pdf.TransformBegin()
		d.pdf.TransformRotate(rotation.Degrees(), x(position.X), y(position.Y))
	}

	d.pdf.Text(x(position.X), y(position.Y), content)

	if rotation != nil {
		d.pdf.TransformEnd()
	}

	if d.pdf.Err() {
		return d.pdf.Error()
	}

	return nil
}

func (d *pdfDrawing) SetLayer(layer string) error {
	switch layer {
	case LAYER_NORMAL:
		return d.findOrCreateLayer(LAYER_NORMAL, 0, 0, 0)
	case LAYER_ALT1:
		return d.findOrCreateLayer(LAYER_ALT1, 128, 128, 128)
	case LAYER_ALT2:
		return d.findOrCreateLayer(LAYER_ALT2, 0, 0, 187)
	default:
		panic("Unknown layer")
	}
}

func (d *pdfDrawing) findOrCreateLayer(name string, r, g, b int) error {
	layerId := d.layerIds[name]
	if layerId == 0 {
		layerId = d.pdf.AddLayer(name, true)
		d.layerIds[name] = layerId
	}

	d.closeOpenLayers()

	d.pdf.BeginLayer(layerId)
	d.pdf.SetDrawColor(r, g, b)
	d.pdf.SetTextColor(r, g, b)

	if err := d.pdf.Error(); err != nil {
		return err
	}

	return nil
}

func (d *pdfDrawing) closeOpenLayers() {
	if d.layerActive {
		d.pdf.EndLayer()
	}
}

func (d *pdfDrawing) SaveAs(filepath string) error {
	d.closeOpenLayers()

	return d.pdf.OutputFileAndClose(filepath)
}

func x(x float64) float64 {
	return PDF_PAGE_MARGIN + x
}

func y(y float64) float64 {
	return PDF_PAGE_MARGIN + -y
}