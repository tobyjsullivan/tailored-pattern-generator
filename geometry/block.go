package geometry

type Block struct {
	StraightLines []*StraightLine
	Points        []*Point
	Text          []*Text
	Blocks        []*Block
}

func (b *Block) AddPoint(p ...*Point) {
	b.Points = append(b.Points, p...)
}

func (b *Block) AddLine(ls ...Line) {
	for _, l := range ls {
		b.StraightLines = append(b.StraightLines, l.StraightLines()...)
	}
}

func (b *Block) AddText(t ...*Text) {
	b.Text = append(b.Text, t...)
}

func (b *Block) AddBlock(blks ...*Block) {
	for _, blk := range blks {
		if b == blk {
			panic("Cannot add block instance to itself.")
		}
	}

	b.Blocks = append(b.Blocks, blks...)
}

func (b *Block) Move(x, y float64) *Block {
	straightLines := make([]*StraightLine, 0, len(b.StraightLines))
	for _, sl := range b.StraightLines {
		sl = sl.Move(x, y)
		straightLines = append(straightLines, sl)
	}

	points := make([]*Point, 0, len(b.Points))
	for _, p := range b.Points {
		p = p.Move(x, y)
		points = append(points, p)
	}

	text := make([]*Text, 0, len(b.Text))
	for _, t := range b.Text {
		t = t.Move(x, y)
		text = append(text, t)
	}

	blocks := make([]*Block, 0, len(b.Blocks))
	for _, blk := range b.Blocks {
		blk = blk.Move(x, y)
		blocks = append(blocks, blk)
	}

	return &Block{
		StraightLines: straightLines,
		Points:        points,
		Text:          text,
		Blocks:        blocks,
	}
}

func (b *Block) BoundingBox() *BoundingBox {
	children := make([]BoundedShape, 0, len(b.StraightLines) + len(b.Blocks) + len(b.Points))
	for _, c := range b.StraightLines {
		children = append(children, c)
	}
	for _, c := range b.Points {
		children = append(children, c)
	}
	for _, c := range b.Blocks {
		children = append(children, c)
	}
	return CollectiveBoundingBox(children...)
}
