package geometry

import (
	"bytes"
	"fmt"
	"github.com/yofu/dxf/drawing"
	"io"
	"math"
)

type CurvedLine struct {
	Start *Point
	End   *Point
}

func (l *CurvedLine) GetStart() *Point {
	return l.Start
}

func (l *CurvedLine) GetEnd() *Point {
	return l.End
}

func (l *CurvedLine) ToEnglish() string {
	return fmt.Sprintf("Curved line from %v to %v\n", l.Start, l.End)
}

func (l *CurvedLine) ToAutoCAD() string {

	rise := l.End.Y - l.Start.Y
	run := l.End.X - l.Start.X

	tangentLength := run - rise
	//radius := rise
	arcStart := &Point{X: l.Start.X + tangentLength, Y: l.Start.Y}
	if math.Abs(rise) > math.Abs(run) {
		tangentLength = rise - run
		arcStart = &Point{X: l.Start.X, Y: l.Start.Y + tangentLength}
		//radius = run
	}

	arcMidpoint := &Point{
		X: arcStart.X,
		Y: l.End.Y,
	}

	tangent := &StraightLine{
		Start: l.Start,
		End:   arcStart,
	}

	w := new(bytes.Buffer)

	// Draw tangent
	fmt.Fprint(w, tangent.ToAutoCAD())

	// Draw arc
	fmt.Fprintf(
		w,
		"(command \"ARC\" \"%.1f,%.1f\" \"C\" \"%.1f,%.1f\" \"%.1f,%.1f\")\n",
		arcStart.X,
		arcStart.Y,
		arcMidpoint.X,
		arcMidpoint.Y,
		l.End.X,
		l.End.Y)

	//drawArc(w, radius, arcStart.X, arcStart.Y, 1.57)

	return w.String()
}

func (l *CurvedLine) DrawDXF(d *drawing.Drawing) error {
	//_, err := d.Line(l.Start.X, l.Start.Y, 0.0, l.End.X, l.End.Y, 0.0)
	for _, line := range l.subLines() {
		line.DrawDXF(d)
	}

	return nil
}

func (l *CurvedLine) subLines() []*StraightLine {
	out := []*StraightLine{}

	numPieces := 10
	rise := l.End.Y - l.Start.Y
	run := l.End.X - l.Start.X

	chunkRotation := ((math.Pi / 2) / float64(numPieces))
	//fmt.Printf("We are drawing a line from %v to %v.\n", l.Start, l.End)
	//fmt.Printf("There are %d chunks of %.2f rads each.\n", numPieces, chunkRotation)
	for i := 0; i < numPieces; i++ {
		t1 := (math.Pi * (3.0 / 2.0)) + (chunkRotation * float64(i))
		t2 := (math.Pi * (3.0 / 2.0)) + (chunkRotation * float64(i+1))
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

func drawArc(w io.Writer, r float64, x float64, y float64, radian float64) {
	fmt.Fprint(w, "(command \"LINE\"")
	for t := 0.0; t < radian; t += 0.1 {
		cx := x + (r * (math.Cos(float64(t - (math.Pi / 2)))))
		cy := y + r + (r * (math.Sin(float64(t - (math.Pi / 2)))))

		fmt.Fprintf(w, " \"%.4f,%.4f\"", cx, cy)
	}
	fmt.Fprint(w, " \"\")\n")
}
