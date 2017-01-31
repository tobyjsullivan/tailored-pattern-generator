package geometry

import (
	"github.com/yofu/dxf/drawing"
	"math"
)

type SimpleCurve struct {
	Start   *Point
	End     *Point
}

func (l *SimpleCurve) DrawDXF(d *drawing.Drawing) error {
	for _, line := range l.subLines() {
		line.DrawDXF(d)
	}

	return nil
}

func (l *SimpleCurve) subLines() []*StraightLine {
	out := []*StraightLine{}

	startAngle := math.Pi * (3.0 / 2.0)
	arcAngle := math.Pi / 2.0

	rise := l.End.Y - l.Start.Y
	run := l.End.X - l.Start.X

	numPieces := 50

	chunkRotation := (arcAngle / float64(numPieces))
	//fmt.Printf("We are drawing a line from %v to %v.\n", l.Start, l.End)
	//fmt.Printf("There are %d chunks of %.2f rads each.\n", numPieces, chunkRotation)
	for i := 0; i < numPieces; i++ {
		t1 := startAngle + (chunkRotation * float64(i))
		t2 := startAngle + (chunkRotation * float64(i+1))
		//fmt.Printf("This chunk uses t1 of %.2f rads and t2 of %.2f rads.\n", t1, t2)

		startX := l.Start.X + run*(math.Cos(t1))
		startY := l.Start.Y + rise*(math.Sin(t1)+1.0)
		start := &Point{X: startX, Y: startY}
		//fmt.Printf("The starting point of this chunk is %v.\n", start)

		endX := l.Start.X + run*(math.Cos(t2))
		endY := l.Start.Y + rise*(math.Sin(t2)+1.0)
		end := &Point{X: endX, Y: endY}
		//fmt.Printf("The ending point of this chunk is %v.\n", end)

		out = append(out, &StraightLine{Start: start, End: end})
	}

	return out
}
