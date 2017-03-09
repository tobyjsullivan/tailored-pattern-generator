package nesting

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"sort"
	"math"
	"fmt"
)

type Container struct {
	Width float64
	Height float64
}

type packedItem struct {
	label int
	rect *Rectangle
	pos *geometry.Point
}

type bySize []*packedItem
func (s bySize) Len() int {	return len(s) }
func (s bySize) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s bySize) Less(i, j int) bool { return s[j].rect.size() < s[i].rect.size() }

type grid [][]bool
func (g grid) Occupied(row, col int) bool { return g[row][col] }
func (g grid) Height() int { return len(g) }
func (g grid) Width() int {
	if g.Height() == 0 {
		return 0
	}

	return len(g[0])
}
func (g grid) Fits(width, height, row, col int) bool {
	if g.Width() <= col + width || g.Height() <= row + height {
		return false
	}

	for r := row; r < row + height; r++ {
		for c := col; c < col + width; c++ {
			if g.Occupied(r, c) {
				return false
			}
		}
	}

	return true
}
func (g grid) Occupy(width, height, row, col int) {
	for y := row; y < row + height; y++ {
		for x := col; x < col + width; x++ {
			g[y][x] = true
		}
	}
}
func (g grid) FindSpace(width, height int) (row, col int, found bool) {
	found = true

	for row = 0; row < g.Height(); row++ {
		for col = 0; col < g.Width(); col++ {
			if g.Fits(width, height, row, col) {
				return
			}
		}
	}

	found = false
	return
}

func (c *Container) Pack(packingList map[int]*Rectangle) map[int]*geometry.Point {
	items := make([]*packedItem, len(packingList))

	i := 0
	for k, v := range packingList {
		items[i] = &packedItem{
			label: k,
			rect: v,
		}
		i++
	}

	// Sort items by size
	sort.Sort(bySize(items))

	// Create a grid covering the container area
	squaresAcross := int(math.Floor(c.Width / 2.0))
	squareSize := c.Width / float64(squaresAcross)
	squaresHigh := int(math.Floor(c.Height / squareSize))

	fmt.Printf("Square size is %.2f\n", squareSize)
	fmt.Printf("Grid size is %dx%d\n", squaresAcross, squaresHigh)

	a := make([][]bool, squaresHigh)
	for i := range a {
		a[i] = make([]bool, squaresAcross)
	}

	grid := grid(a)

	// Fit items into grid
	for _, item := range items {
		itemWidth := int(math.Ceil(item.rect.Width / squareSize))
		itemHeight := int(math.Ceil(item.rect.Height / squareSize))

		// Scan full grid for a spot to fit item
		row, col, found := grid.FindSpace(itemWidth, itemHeight)

		if !found {
			fmt.Printf("Couldn't find any space for R: %v\n", item.rect)
			continue
		}

		fmt.Printf("An item fits at %dx%d\n", col, row)

		item.pos = &geometry.Point{X: float64(col) * squareSize, Y: float64(row) * squareSize}
		grid.Occupy(itemWidth, itemHeight, row, col)
	}

	result := make(map[int]*geometry.Point)
	for _, item := range items {
		result[item.label] = item.pos
	}

	return result
}
