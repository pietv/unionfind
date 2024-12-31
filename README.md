UnionFind [![Go Reference](https://pkg.go.dev/badge/github.com/pietv/unionfind.svg)](https://pkg.go.dev/github.com/pietv/unionfind)
=========

This is an implementation of the UnionFind (disjoint-set) data structure, as described, for example,
here: http://algs4.cs.princeton.edu/15uf .

The `Union()` and `Connected()` operations take $`O(log^* N)`$ “log-star” time, which is close to $`O(1)`$.

Install
=======

```shell
$ go get github.com/pietv/unionfind@latest
```

Example
=======

Using the `UnionFind` data structure to find the number of “islands” (independent,
disconnected areas of dry land) on a chart.

* The space ` ` character defines the sea.
* The `.` character defines the dry land.

```go
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

				// Union with a piece of land above.
				if y > 0 && len(lines[y-1]) > x && lines[y-1][x] == '.' {
					sets.Union(ξ{x, y}, ξ{x, y - 1})
				}
			}
		}
	}
	return sets.Count()
}
```

For the chart above, the code returns **4** (four islands).
