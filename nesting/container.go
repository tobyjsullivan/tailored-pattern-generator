package nesting

import (
	"github.com/tailored-style/pattern-generator/geometry"
	"sort"
)

type Container struct {
	Width float64
}

type packedItem struct {
	label string
	rect *Rectangle
	pos *geometry.Point
}

type bySize []*packedItem
func (s bySize) Len() int {	return len(s) }
func (s bySize) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s bySize) Less(i, j int) bool { return s[i].rect.size() < s[j].rect.size() }

func (c *Container) Pack(packingList map[string]*Rectangle) map[string]*geometry.Point {
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


}
