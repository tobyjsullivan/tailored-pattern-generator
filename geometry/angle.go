package geometry

import (
	"math"
)

type Angle struct {
	Rads float64
}

// Returns the equivalent value between -math.Pi and +math.Pi
func normalizeRads(rads float64) float64 {
	r := math.Mod(rads, math.Pi * 2.0)

	if r > math.Pi {
		r -= math.Pi * 2.0
	}

	return r
}

func (a *Angle) Radians() float64 {
	return normalizeRads(a.Rads)
}

func (a *Angle) Degrees() float64 {
	return 360.0 * (a.Radians() / (2.0 * math.Pi))
}

func (a *Angle) Perpendicular() *Angle {
	return &Angle{
		Rads: normalizeRads(a.Rads + (math.Pi / 2.0)),
	}
}

func (a *Angle) Opposite() *Angle {
	return &Angle{
		Rads: normalizeRads(a.Rads + math.Pi),
	}
}

func (a *Angle) Add(other *Angle) *Angle {
	return &Angle{
		Rads: a.Rads + other.Rads,
	}
}

func (a *Angle) Subtract(other *Angle) *Angle {
	return &Angle{
		Rads: a.Rads - other.Rads,
	}
}

func (a *Angle) Divide(x float64) *Angle {
	return &Angle{
		Rads: a.Rads / x,
	}
}

func (a *Angle) Multiply(x float64) *Angle {
	return &Angle{
		Rads: a.Rads * x,
	}
}

func (a *Angle) Sin() float64 {
	return math.Sin(a.Radians())
}

func (a *Angle) Cos() float64 {
	return math.Cos(a.Radians())
}