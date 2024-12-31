// This example demonstrates the usage of the UnionFind data structure
// to count the number of “islands” (independent, disconnected areas
// of “land”) on a chart.
//
//   - The space ‘ ’ character defines the sea.
//   - The ‘.’ character defines the land.
package unionfind_test

import (
	"fmt"
	"strings"

	"github.com/pietv/unionfind"
)

const chart = `
  ......   .
  .    ..
       ..
       .....
  ..       .
  .  ..... .
         . .
`

// ξ (for ξηρɑ, dry land) is used as a symbol for point.
type ξ struct{ x, y int }

func scan(chart string) int {
	var (
		lines = strings.Split(chart, "\n")
		sets  = unionfind.New()
	)
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if line[x] == '.' {
				// Add a standalone piece of land as a disjoint set.
				sets.MakeSet(ξ{x, y})

				// Union it with a piece of land to the left.
				if x > 0 && line[x-1] == '.' {
					sets.Union(ξ{x, y}, ξ{x - 1, y})
				}

				// Union it with a piece of land above.
				if y > 0 && len(lines[y-1]) > x && lines[y-1][x] == '.' {
					sets.Union(ξ{x, y}, ξ{x, y - 1})
				}
			}
		}
	}
	return sets.Count()
}

func ExampleCountIslands() {
	fmt.Printf("number of islands = %d\n", scan(chart))
	// Output: number of islands = 4
}
