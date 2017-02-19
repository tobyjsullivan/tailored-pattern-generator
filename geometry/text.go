package geometry

type Text struct {
	Content  string
	Position *Point
	Rotation *Angle
}

const (
	TEXT_CHAR_WIDTH = 1.0
	TEXT_CHAR_HEIGHT = 1.0
)

func (t *Text) Move(x, y float64) *Text {
	var out Text = *t
	out.Position = t.Position.Move(x, y)
	return &out
}

func (t * Text) BoundingBox() *BoundingBox {
	return &BoundingBox{
		Top: t.Position.Y,
		Left: t.Position.X,
		Right: t.Position.X + float64(len(t.Content)) * TEXT_CHAR_WIDTH,
		Bottom: t.Position.Y + TEXT_CHAR_HEIGHT,
	}
}
