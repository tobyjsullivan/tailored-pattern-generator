package pieces

import "github.com/tailored-style/pattern-generator/geometry"

func SeamAllowance(line ...geometry.Line) geometry.Line {
	poly := &geometry.Polyline{}

	var prev geometry.Line = nil
	for _, l := range line {
		if prev != nil {
			poly.AddLine(geometry.Connect(prev, l))
		}

		poly.AddLine(l)

		prev = l
	}

	return poly
}
