package geometry

import (
	"github.com/yofu/dxf/drawing"
	"math"
	"fmt"
)

type EllipseCurve struct {
	Start             *Point
	End               *Point
	StartingAngleRads float64
	ArcAngle          float64
}

func (c *EllipseCurve) DrawDXF(d *drawing.Drawing) error {
	for _, line := range c.subLines() {
		line.DrawDXF(d)
	}

	return nil
}

func (c *EllipseCurve) subLines() []*StraightLine {
	out := []*StraightLine{}
	fmt.Printf("Drawing an ellipse curve from %v to %v.\n", c.Start, c.End)

	arcStartAngle := c.StartingAngleRads
	arcEndAngle := arcStartAngle + c.ArcAngle
	fmt.Printf("Drawing arc from %.2f to %.2f\n", arcStartAngle, arcEndAngle)

	// Simulate a simple circle until segment angle >= end angle
	sx := c.Start.X + math.Cos(arcStartAngle)
	sy := c.Start.Y + math.Sin(arcStartAngle)

	ex := c.Start.X + math.Cos(arcEndAngle)
	ey := c.Start.Y + math.Sin(arcEndAngle)

	fmt.Printf("Arc ends at (%.1f, %.1f)\n", ex, ey)

	// Shift is how much the arc start is offset from curved line start
	shiftX := c.Start.X - sx
	shiftY := c.Start.Y - sy

	// Compute the x and y scaling required to fit end point
	pointDistX := c.End.X - c.Start.X
	pointDistY := c.End.Y - c.Start.Y

	arcEndDistX := ex - sx
	arcEndDistY := ey - sy

	fmt.Printf("Distance between points is %.2f x %.2f\n", pointDistX, pointDistY)
	fmt.Printf("Distance between arc endpoints is %.2f x %.2f\n", arcEndDistX, arcEndDistY)

	scaleX :=  pointDistX / arcEndDistX
	scaleY := pointDistY / arcEndDistY
	//scaleX = 1.0
	//scaleY = 1.0

	fmt.Printf("Scaling at %.2fx%.2f\n", scaleX, scaleY)

	// Draw out the transform
	numPieces := 50
	chunkSize := c.ArcAngle / float64(numPieces)

	for i := 0; i < numPieces; i++ {
		sRad := arcStartAngle + chunkSize * float64(i)
		eRad := arcStartAngle + chunkSize * float64(i + 1)

		p1 := &Point{
			X: c.Start.X + scaleX * (math.Cos(sRad) + shiftX),
			Y: c.Start.Y + scaleY * (math.Sin(sRad) + shiftY),
		}

		p2 := &Point{
			X: c.Start.X + scaleX * (math.Cos(eRad) + shiftX),
			Y: c.Start.Y + scaleY * (math.Sin(eRad) + shiftY),
		}

		out = append(out, &StraightLine{Start: p1, End: p2})
	}

	return out
}