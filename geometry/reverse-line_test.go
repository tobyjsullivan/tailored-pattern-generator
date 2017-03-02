package geometry

import (
	"testing"
	"math"
)

func TestReverseLine_StraightLines(t *testing.T) {
	poly := &Polyline{}

	poly.AddLine(
		&StraightLine{
			Start: &Point{X: 0.0, Y: 0.0},
			End: &Point{X: 5.4, Y: 2.3},
		},
		&StraightLine{
			Start: &Point{X: 5.4, Y: 2.3},
			End: &Point{X: 4.3, Y: -2.0},
		},
		&StraightLine{
			Start: &Point{X: 4.3, Y: -2.0},
			End: &Point{X: 1.3, Y: 0.3},
		},
	)

	reverse := &ReverseLine{InnerLine: poly}

	expectedLineCount := 3
	expectedLine0 := &StraightLine{
		Start: &Point{X: 1.3, Y: 0.3},
		End: &Point{X: 4.3, Y: -2.0},
	}
	expectedLine1 := &StraightLine{
		Start: &Point{X: 4.3, Y: -2.0},
		End: &Point{X: 5.4, Y: 2.3},
	}
	expectedLine2 := &StraightLine{
		Start: &Point{X: 5.4, Y: 2.3},
		End: &Point{X: 0.0, Y: 0.0},
	}

	lines := reverse.StraightLines()

	if len(lines) != expectedLineCount {
		t.Errorf("Wrong number of lines. Expected: %d; Got: %d.", expectedLineCount, len(lines))
	}

	if !lines[0].Equals(expectedLine0) {
		t.Errorf("Line 0 didn't match expected. Expected: %v; Got: %v", expectedLine0, lines[0])
	}

	if !lines[1].Equals(expectedLine1) {
		t.Errorf("Line 1 didn't match expected. Expected: %v; Got: %v", expectedLine1, lines[1])
	}

	if !lines[2].Equals(expectedLine2) {
		t.Errorf("Line 2 didn't match expected. Expected: %v; Got: %v", expectedLine2, lines[2])
	}
}

func TestReverseLine_AngleAt(t *testing.T) {
	poly := &Polyline{}

	poly.AddLine(
		&StraightLine{
			Start: &Point{X: 0.0, Y: 0.0},
			End: &Point{X: 5.4, Y: 2.3},
		},
		&StraightLine{
			Start: &Point{X: 5.4, Y: 2.3},
			End: &Point{X: 4.3, Y: -2.0},
		},
		&StraightLine{
			Start: &Point{X: 4.3, Y: -2.0},
			End: &Point{X: 1.3, Y: 0.3},
		},
	)

	reverse := &ReverseLine{InnerLine: poly}

	expectedAngleStart := &Angle{Rads: math.Atan2(-2.0 - 0.3, 4.3 - 1.3)}
	expectedAngleEnd := &Angle{Rads: math.Atan2(0.0 - 2.3, 0.0 - 5.4)}

	if a := reverse.AngleAt(0.0); !a.Equivalent(expectedAngleStart) {
		t.Errorf("Wrong starting angle. Expected %v; Got %v.", expectedAngleStart, a)
	}

	if a := reverse.AngleAt(reverse.Length() - 0.001); !a.Equivalent(expectedAngleEnd) {
		t.Errorf("Wrong ending angle. Expected %v; Got %v.", expectedAngleEnd, a)
	}
}
