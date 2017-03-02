package geometry

import (
	"testing"
	"math"
)

func TestStraightLine_StraightLines(t *testing.T) {
	sl := &StraightLine{
		Start: &Point{X: 0.3, Y: 0.6},
		End: &Point{X: -4.5, Y: 5.3},
	}

	lines := sl.StraightLines()

	if len(lines) != 1 {
		t.Errorf("Wrong number of lines: %d", len(lines))
	}

	if !lines[0].Equals(sl) {
		t.Errorf("Incorrect straight line found. Expected: %v; Got: %v", sl, lines[0])
	}
}

func TestStraightLine_AngleAt(t *testing.T) {
	sl := &StraightLine{
		Start: &Point{X: 0.3, Y: 0.6},
		End: &Point{X: -4.5, Y: 5.3},
	}

	expected := &Angle{Rads: math.Atan2(5.3 - 0.6, -4.5 - 0.3)}

	a0 := sl.AngleAt(0.0)
	if !a0.Equivalent(expected) {
		t.Errorf("Wrong angle detected. Expected: %.3f; Got: %.3f", expected, a0)
	}
}

func TestStraightLine_Equals(t *testing.T) {
	sl0 := &StraightLine{
		Start: &Point{X: 0.3, Y: 0.6},
		End: &Point{X: -4.5, Y: 5.3},
	}

	sl1 := &StraightLine{
		Start: &Point{X: 0.3, Y: 0.6},
		End: &Point{X: -4.5, Y: 5.3},
	}

	sl2 := &StraightLine{
		Start: &Point{X: 0.3, Y: 0.6},
		End: &Point{X: -5.5, Y: 5.3},
	}

	if !sl0.Equals(sl1) {
		t.Error("Expected straight lines to be equal but they were not.")
	}

	if sl0.Equals(sl2) {
		t.Error("Expected straight lines to not be equal but they were.")
	}
}

func TestStraightLine_Reverse(t *testing.T) {
	sl := &StraightLine{
		Start: &Point{X: 0.3, Y: 0.6},
		End: &Point{X: -4.5, Y: 5.3},
	}

	reverse := sl.Reverse()

	expectedStart := &Point{X: -4.5, Y: 5.3}
	expectedEnd := &Point{X: 0.3, Y: 0.6}

	if s := reverse.Start; !s.Equals(expectedStart) {
		t.Errorf("Wrong start. Expected: %v; Got: %v.", expectedStart, s)
	}

	if e := reverse.End; !e.Equals(expectedEnd) {
		t.Errorf("Wrong end. Expected: %v; Got: %v.", expectedEnd, e)
	}

	expectedTangentOrigin := expectedStart
	expectedTangentDirection := &Angle{Rads: math.Atan2(0.6 - 5.3, 0.3 - -4.5)}

	tangent := reverse.TangentAt(0.0)
	if o := tangent.Origin; !o.Equals(expectedTangentOrigin) {
		t.Errorf("Wrong tangent origin. Expected: %v; Got: %v", expectedTangentOrigin, o)
	}

	if d := tangent.Direction; !d.Equivalent(expectedTangentDirection) {
		t.Errorf("Wrong tangent direction. Expected: %v; Got: %v", expectedTangentDirection, d)
	}

}