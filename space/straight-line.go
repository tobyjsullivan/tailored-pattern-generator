package space

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"math"
)

type StraightLine struct {
	geometry.StraightLine
}

func (l *StraightLine) GetAnchor(dist float64) Anchor {
	return &straightLineAnchor{
		line: l,
		dist: dist,
	}
}

type straightLineAnchor struct {
	line *StraightLine
	dist float64
}

func (a *straightLineAnchor) GetPosition() *geometry.Point {
	line := a.line
	lrise := line.End.Y - line.Start.Y
	lrun := line.End.X - line.Start.X

	ang := math.Atan(lrise / lrun)
	if lrun < 0.0 {
		ang = math.Pi / 2.0 - ang
	}

	ax := a.dist * math.Cos(ang)
	ay := a.dist * math.Sin(ang)

	return &geometry.Point{
		X: ax,
		Y: ay,
	}
}
