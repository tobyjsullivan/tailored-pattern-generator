package geometry

import (
	"math"
	"fmt"
)

type Ellipse struct {
	XRadius float64
	YRadius float64
	Origin  *Point
}

func (e *Ellipse) f(t float64) float64 {
	return e.XRadius * math.Cos(t)
}

func (e *Ellipse) g(t float64) float64 {
	return e.YRadius * math.Sin(t)
}

func (e *Ellipse) PointAt(t float64) *Point {
	return &Point{
		X: e.Origin.X + e.f(t),
		Y: e.Origin.Y + e.g(t),
	}
}

func (e *Ellipse) Slope(t float64) float64 {
	x := e.f(t)
	y := e.g(t)

	if y == 0.0 {
		return math.Inf(int(-x))
	}

	m := -(math.Pow(e.YRadius, 2.0)*x)/(math.Pow(e.XRadius, 2.0)*y)

	return m
}

func (e *Ellipse) Arc(t0, t1 float64) []*StraightLine {
	numPieces := 200
	pieceArc := (t1 - t0) / float64(numPieces)

	out := make([]*StraightLine, numPieces)
	for i := 0; i < numPieces; i++ {
		start := e.PointAt(t0 + pieceArc * float64(i))
		end := e.PointAt(t0 + pieceArc * float64(i + 1))

		out[i] = &StraightLine{
			Start: start,
			End: end,
		}
	}

	return out
}

func (e *Ellipse) recFindT (slope float64, st float64, incr float64, dep int) float64 {
	lastM := math.Inf(-1)
	for t := st; t < 2.0 * math.Pi; t += incr {
		m := e.Slope(t)
		//fmt.Printf("Slope at t=%.4f is %.4f\n", t, m)
		if math.IsInf(m, 0) {
			continue
		}

		if m >= slope {
			fmt.Printf("T value of %.12f is looking good\n", t)
			if dep == 0 {
				if math.Abs(m) - math.Abs(slope) < math.Abs(lastM) - math.Abs(slope) {
					return t
				} else {
					return t - incr
				}
			}

			return e.recFindT(slope, t - incr, incr / 2, dep - 1)
		}

		lastM = m
	}

	if dep == 0 {
		fmt.Printf("Oops. Could not find T for slope %.2f\n", slope)
		return 0.0
	}

	return e.recFindT(slope, math.Pi - incr, incr / 128.0, dep - 1)
}

func (e *Ellipse) FindT(slope float64) float64 {
	fmt.Printf("Finding T for slope %.2f in %v\n", slope, e)

	if math.IsInf(slope, 0) {
		return math.Pi
	}

	return e.recFindT(slope, math.Pi + 0.0000001, math.Pi / 128.0, 15)
}

func FitEllipseToPoints(p0, p1 *Point, a0, a1 float64) (e *Ellipse, t0, t1 float64) {
	m0 := math.Tan(a0)
	m1 := math.Tan(a1)
	fmt.Printf("Fitting ellipse for a0 = %.2f, a1 = %.2f, m0 = %.2f, m1 = %.2f\n", a0, a1, m0, m1)

	distEpsilonX := 2.0
	distEpsilonY := 0.01
	slopeEpsilon := 0.01

	e = &Ellipse{
		XRadius: 1.0,
		YRadius: 1.0,
		Origin: &Point{X: 0.0, Y: 0.0},
	}

	fit := func(e *Ellipse, t0, t1 float64, p0 , p1 *Point, a0, a1 float64) bool {
		widthMiss := (p1.X - p0.X) - (e.PointAt(t1).X - e.PointAt(t0).X)
		heightMiss := (p1.Y - p0.Y) - (e.PointAt(t1).Y - e.PointAt(t0).Y)
		a0Miss := math.Abs(e.Slope(t0) - m0)
		a1Miss := math.Abs(e.Slope(t1) - m1)

		fmt.Printf("T0: %.6f, T1: %.6F; Ellipse: %v\n", t0, t1, e)
		fmt.Printf("Width is off by %.3f; Height is off by %.3f; A0 is off by %.3f; A1 is off by %.3f.\n", widthMiss, heightMiss, a0Miss, a1Miss)

		return widthMiss >= 0.00 &&
				math.Abs(widthMiss) < distEpsilonX &&
				math.Abs(heightMiss) < distEpsilonY &&
				a0Miss < slopeEpsilon &&
				a1Miss < slopeEpsilon
	}

	// Find a point, t0, where the slope matches a0
	t0 = e.FindT(m0)

	// Find a point, t1, where the slope matches a1
	t1 = e.FindT(m1)

	targetWidth := p1.X - p0.X
	targetHeight := p1.Y - p0.Y

	for i := 0; !fit(e, t0, t1, p0, p1, a0, a1) && i < 10000000; i++ {
		q0 := e.PointAt(t0)
		q1 := e.PointAt(t1)

		for math.Abs((q1.Y - q0.Y) - targetHeight) > distEpsilonY {
			arcHeight := math.Abs(q1.Y - q0.Y)
			fmt.Printf("We need height to be %.3f but it is actually %.3f\n", targetHeight, arcHeight)

			scale := math.Abs(targetHeight / arcHeight)
			fmt.Printf("Scaling x- and y-radius by %.4f\n", scale)
			e.XRadius *= scale
			e.YRadius *= scale

			// Refind t0 and t1
			t0 = e.FindT(m0)
			t1 = e.FindT(m1)

			q0 = e.PointAt(t0)
			q1 = e.PointAt(t1)

			arcHeight = math.Abs(q1.Y - q0.Y)
			fmt.Printf("Now height ought to be %.3f but it is actually %.3f\n", targetHeight, arcHeight)

		}

		if math.Abs((q1.X - q0.X) - targetWidth) > distEpsilonX || targetWidth - (q1.X - q0.X) < 0.00 {
			arcWidth := math.Abs(q1.X - q0.Y)
			fmt.Printf("We need width to be %.3f but it is actually %.3f\n", targetWidth, arcWidth)

			scale := math.Abs(targetWidth / arcWidth)
			fmt.Printf("Scaling x-radius by %.4f\n", scale)
			e.XRadius *= scale

			// Refind t0 and t1
			t0 = e.FindT(m0)
			t1 = e.FindT(m1)

			q0 = e.PointAt(t0)
			q1 = e.PointAt(t1)

			arcWidth = math.Abs(q1.X - q0.Y)
			fmt.Printf("Now width ought to be %.3f but it is actually %.3f\n", targetWidth, arcWidth)
		}
	}

	// Now move the ellipse so start of arc is at p0
	start := e.PointAt(t0)
	offsetX := p0.X - start.X
	offsetY := p0.Y - start.Y

	e.Origin = e.Origin.Move(offsetX, offsetY)

	fmt.Println("Fin.")
	return
}

func (e *Ellipse) String() string {
	return fmt.Sprintf("(%.2fX%.2f; %v)", e.XRadius, e.YRadius, e.Origin)
}
