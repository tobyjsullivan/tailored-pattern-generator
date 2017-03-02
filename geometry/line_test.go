package geometry

import (
	"testing"
	"math"
)

func TestTangentAtLineStart_StraightLine(t *testing.T) {
	sl := &StraightLine{
		Start: &Point{X: 0.0, Y: 0.0},
		End: &Point{X: 5.4, Y: 2.3},
	}

	expectedOrigin := &Point{X: 0.0, Y: 0.0}
	expectedDirection := &Angle{Rads: math.Atan2(2.3, 5.4)}

	tan := TangentAtLineStart(sl)

	if !tan.Origin.Equals(expectedOrigin) {
		t.Errorf("Wrong origin. Expected: %v; Got: %v.", expectedOrigin, tan.Origin)
	}

	if !tan.Direction.Equivalent(expectedDirection) {
		t.Errorf("Wrong direction. Expected: %v; Got: %v.", expectedDirection, tan.Direction)
	}
}

func TestTangentAtLineStart_Polyline(t *testing.T) {
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
	expectedOrigin := &Point{X: 0.0, Y: 0.0}
	expectedDirection := &Angle{Rads: math.Atan2(2.3, 5.4)}

	tan := TangentAtLineStart(poly)

	if !tan.Origin.Equals(expectedOrigin) {
		t.Errorf("Wrong origin. Expected: %v; Got: %v.", expectedOrigin, tan.Origin)
	}

	if !tan.Direction.Equivalent(expectedDirection) {
		t.Errorf("Wrong direction. Expected: %v; Got: %v.", expectedDirection, tan.Direction)
	}
}


func TestTangentAtLineStart_ReverseLine(t *testing.T) {
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

	expectedOrigin := &Point{X: 1.3, Y: 0.3}
	expectedDirection := &Angle{Rads: math.Atan2(-2.0 - 0.3, 4.3 - 1.3)}

	tan := TangentAtLineStart(reverse)

	if !tan.Origin.Equals(expectedOrigin) {
		t.Errorf("Wrong origin. Expected: %v; Got: %v.", expectedOrigin, tan.Origin)
	}

	if !tan.Direction.Equivalent(expectedDirection) {
		t.Errorf("Wrong direction. Expected: %v; Got: %v.", expectedDirection, tan.Direction)
	}
}

func TestTangentAtLineEnd_StraightLine(t *testing.T) {
	sl := &StraightLine{
		Start: &Point{X: 0.0, Y: 0.0},
		End: &Point{X: 5.4, Y: 2.3},
	}

	expectedOrigin := &Point{X: 5.4, Y: 2.3}
	expectedDirection := &Angle{Rads: math.Atan2(2.3, 5.4)}

	tan := TangentAtLineEnd(sl)

	if !tan.Origin.Equals(expectedOrigin) {
		t.Errorf("Wrong origin. Expected: %v; Got: %v.", expectedOrigin, tan.Origin)
	}

	if !tan.Direction.Equivalent(expectedDirection) {
		t.Errorf("Wrong direction. Expected: %v; Got: %v.", expectedDirection, tan.Direction)
	}
}

func TestTangentAtLineEnd_Polyline(t *testing.T) {
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
	expectedOrigin := &Point{X: 1.3, Y: 0.3}
	expectedDirection := &Angle{Rads: math.Atan2(0.3 -  -2.0, 1.3 - 4.3)}

	tan := TangentAtLineEnd(poly)

	if !tan.Origin.Equals(expectedOrigin) {
		t.Errorf("Wrong origin. Expected: %v; Got: %v.", expectedOrigin, tan.Origin)
	}

	if !tan.Direction.Equivalent(expectedDirection) {
		t.Errorf("Wrong direction. Expected: %v; Got: %v.", expectedDirection, tan.Direction)
	}
}

func TestTangentAtLineEnd_ReverseLine(t *testing.T) {
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

	expectedOrigin := &Point{X: 0.0, Y: 0.0}
	expectedDirection := &Angle{Rads: math.Atan2(0.0 - 2.3, 0.0 - 5.4)}

	tan := TangentAtLineEnd(reverse)

	if o := tan.Origin; !o.Equals(expectedOrigin) {
		t.Errorf("Wrong origin. Expected: %v; Got: %v.", expectedOrigin, o)
	}

	if d := tan.Direction; !d.Equivalent(expectedDirection) {
		t.Errorf("Wrong direction. Expected: %v; Got: %v.", expectedDirection, d)
	}
}
