package nesting

type Rectangle struct {
	Width float64
	Height float64
}

func (r *Rectangle) size() float64 {
	return r.Width * r.Height
}
