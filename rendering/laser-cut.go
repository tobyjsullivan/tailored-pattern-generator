package rendering

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"github.com/tailored-style/pattern-generator/drawing"
	"github.com/tailored-style/pattern-generator/styles"
	"github.com/tailored-style/pattern-generator/nesting"
	"fmt"
)

const (
	LASER_CUT_WIDTH = 152.4 // 60"
	LASER_CUT_MAX_HEIGHT = 121.92 // 48"
)

type LaserCut struct {
	Blocks []*geometry.Block
}

func NewLaserCutFromStyle(style styles.Style) *LaserCut {
	lc := &LaserCut{}
	lc.Blocks = []*geometry.Block{}

	for _, p := range style.Pieces() {
		lc.Blocks = append(lc.Blocks, p.CutLayer())
	}

	return lc
}

func (l *LaserCut) SavePDF(filepath string) error {
	pdf := drawing.NewPDF(LASER_CUT_WIDTH)

	for pos, blk := range l.nestBlocks(pdf) {
		fmt.Printf("Drawing a block at %v\n", pos)

		bbox := blk.BoundingBox()
		offset := &geometry.Point{
			X: pos.X - bbox.Left,
			Y: pos.Y - bbox.Top,
		}

		DrawBlock(pdf, blk, offset)
	}

	return pdf.SaveAs(filepath)
}

func (l *LaserCut) nestBlocks(d drawing.Drawing) map[*geometry.Point]*geometry.Block {
	// Create a container to pack
	c := &nesting.Container{
		Width: d.DrawableWidth(),
		Height: LASER_CUT_MAX_HEIGHT,
	}

	// Initialize the packing list
	blocks := make(map[int]*geometry.Block)
	for i, blk := range l.Blocks {
		blocks[i] = blk
	}

	packingList := make(map[int]*nesting.Rectangle)
	for label, blk := range blocks {
		bbox := blk.BoundingBox()
		packingList[label] = &nesting.Rectangle{
			Width: bbox.Width(),
			Height: bbox.Height(),
		}
	}

	positions := c.Pack(packingList)
	out := make(map[*geometry.Point]*geometry.Block)
	for label, blk := range blocks {
		pos := positions[label]

		if pos == nil {
			panic("Couldn't fit a block")
		}

		pos = &geometry.Point{
			X: pos.X,
			Y: -pos.Y,
		}

		out[pos] = blk
	}

	return out
}