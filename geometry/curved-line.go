package geometry

import (
    "fmt"
    "bytes"
    "math"
    "io"
)

type CurvedLine struct {
    Start *Point
    End *Point
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
        End: arcStart,
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

func drawArc(w io.Writer, r float64, x float64, y float64, radian float64) {
    fmt.Fprint(w, "(command \"LINE\"")
    for t := 0.0; t < radian; t += 0.1 {
        cx := x + (r * (math.Cos(float64(t - (math.Pi/2)))))
        cy := y + r + (r * (math.Sin(float64(t - (math.Pi/2)))))

        fmt.Fprintf(w, " \"%.4f,%.4f\"", cx, cy)
    }
    fmt.Fprint(w, " \"\")\n")
}
