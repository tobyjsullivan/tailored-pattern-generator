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

	collectiveBox := cs[0].BoundingBox()

	for i := 1; i < len(cs); i++ {
		box := cs[i].BoundingBox()

		// Empty blocks can have no bounding box
		if box == nil {
			continue
		}

		if collectiveBox == nil {
			collectiveBox = box
		}

		if box.Top > collectiveBox.Top {
			collectiveBox.Top = box.Top
		}

		if box.Left < collectiveBox.Left {
			collectiveBox.Left = box.Left
		}

		if box.Right > collectiveBox.Right {
			collectiveBox.Right = box.Right
		}

		if box.Bottom < collectiveBox.Bottom {
			collectiveBox.Bottom = box.Bottom
		}
	}

	return collectiveBox
}
