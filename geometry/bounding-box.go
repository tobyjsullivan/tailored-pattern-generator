package geometry

type BoundingBox struct {
	Top    float64
	Left   float64
	Right  float64
	Bottom float64
}

func (bb *BoundingBox) Width() float64 {
	return bb.Right - bb.Left
}

func (bb *BoundingBox) Height() float64 {
	return bb.Top - bb.Bottom
}

type BoundedShape interface {
	BoundingBox() *BoundingBox
}

func CollectiveBoundingBox(cs ...BoundedShape) *BoundingBox {
	if len(cs) == 0 {
		return nil
	}

	box := cs[0].BoundingBox()
	top := box.Top
	left := box.Left
	right := box.Right
	bottom := box.Bottom

	for i := 1; i < len(cs); i++ {
		box = cs[i].BoundingBox()

		if box.Top > top {
			top = box.Top
		}

		if box.Left < left {
			left = box.Left
		}

		if box.Right > right {
			right = box.Right
		}

		if box.Bottom < bottom {
			bottom = box.Bottom
		}
	}

	return &BoundingBox{
		Top: top,
		Left: left,
		Right: right,
		Bottom: bottom,
	}
}
