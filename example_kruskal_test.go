// This example demonstrates the usage of UnionFind data structure
// to implement the Kruskal's Minimum Spanning Tree (MST) algorithm.
package unionfind_test

import (
	"container/heap"
	"fmt"

	"github.com/pietv/unionfind"
)

type Vertex interface{}
type Edge struct {
	v1, v2 Vertex
	w      int
}

func (e Edge) String() string {
	return fmt.Sprintf("(%v, %v)", e.v1, e.v2)
}

type PriorityQueue []Edge

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].w < pq[j].w }
func (pq *PriorityQueue) Push(x interface{}) {
	edge := x.(Edge)
	*pq = append(*pq, edge)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	edge := old[n-1]
	*pq = old[0 : n-1]
	return edge

}

func ExampleUnionFind() {
	//
	//                                MST
	//
	//  (1)----4----(2)         (1)---------(2)
	//   | \         |           |
	//   |   \       |           |
	//   1     7     8           |
	//   |       \   |           |
	//   |         \ |           |
	//  (3)----5----(4)         (3)---------(4)
	//   | \         |             \
	//   |   \       |               \
	//   4     3     6                 \
	//   |       \   |                   \
	//   |         \ |                     \
	//  (5)----2----(6)         (5)---------(6)
	//
	edges := []Edge{
		{1, 2, 4}, {1, 3, 1}, {2, 4, 8},
		{1, 4, 7}, {3, 5, 4}, {3, 4, 5},
		{4, 6, 6}, {3, 6, 3}, {5, 6, 2},
	}

	pq := PriorityQueue{}
	heap.Init(&pq)
	u := unionfind.New()
	for _, e := range edges {
		heap.Push(&pq, e)
		u.MakeSet(e.v1, e.v2)
	}

	mst := []Edge{}
	for {
		if pq.Len() == 0 {
			break
		}
		e := heap.Pop(&pq)
		v1, v2 := e.(Edge).v1, e.(Edge).v2
		if !u.Connected(v1, v2) {
			u.Union(v1, v2)
			mst = append(mst, e.(Edge))
		}
	}

	fmt.Println(mst)
	// Output: [(1, 3) (5, 6) (3, 6) (1, 2) (3, 4)]
}
