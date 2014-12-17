// Package unionfind implements a UnionFind (disjoint-set) data structure,
// as described, for example, here: http://algs4.cs.princeton.edu/15uf .
//
// The Union() and Connected() operations take O(log* N) “log-star” time,
// which is close to O(1).
//
// The MakeSet() operation when used with multiple arguments, unions elements
// in one set.
//
// Basic usage:
//
//   u := unionfind.New()
//
//   // Create sets.
//   u.MakeSet(1, 2, 3, 4)
//
//   // Join them together.
//   u.Union(1, 2)
//   u.Union(3, 4)
//   u.Union(2, 3)
//
//   // Check if they're connected.
//   fmt.Println(u.Connected(1, 4))
//
package unionfind

import (
	"fmt"
	"strings"
)

// Maintains sets and a number of connected elements.
type UnionFind struct {
	sets  map[interface{}]*set
	count int
}

type set struct {
	parent interface{}
	rank   int
}

// New return an initialized UnionFind data structure.
func New() *UnionFind {
	return &UnionFind{
		sets: make(map[interface{}]*set),
	}
}

// MakeSet makes an independent set of one element.  If called with multiple
// arguments, an independent set for every element is made.
func (u *UnionFind) MakeSet(a ...interface{}) {
	if len(a) == 0 {
		return
	}

	for _, elem := range a {
		if elem == nil {
			continue
		}

		// Skip already made sets.
		if _, ok := u.sets[elem]; ok {
			continue
		}

		u.sets[elem] = &set{parent: elem}
		u.count++
	}
}

// Union merges two independent sets as one. The number of sets is decreased by 1.
func (u *UnionFind) Union(x, y interface{}) {
	a := u.Find(x)
	b := u.Find(y)

	// Already connected.
	if a == b {
		return
	}

	// Weighting.
	switch {
	case u.sets[a].rank < u.sets[b].rank:
		u.sets[a].parent = b

	case u.sets[a].rank > u.sets[b].rank:
		u.sets[b].parent = a

	case u.sets[a].rank == u.sets[b].rank:
		u.sets[b].parent = a
		u.sets[a].rank++
	}

	u.count--
}

// Find returns the root element of the set. The root element is the same for
// all elements within the same set.
func (u UnionFind) Find(x interface{}) interface{} {
	if _, ok := u.sets[x]; !ok {
		panic(fmt.Sprintf("set %v hasn't been made yet with MakeSet", x))
	}

	// The root.
	if u.sets[x].parent == x {
		return x
	}

	// Path compression.
	u.sets[x].parent = u.Find(u.sets[x].parent)

	return u.sets[x].parent
}

// Exists returns true if the element belongs to any set, false otherwise.
func (u UnionFind) Exists(x interface{}) bool {
	if _, ok := u.sets[x]; ok {
		return true
	} else {
		return false
	}
}

// Connected returns true if the elements belong to the same set,
// false otherwise.
func (u UnionFind) Connected(x, y interface{}) bool {
	return u.Find(x) == u.Find(y)
}

// Count returns the number of independent sets.
func (u UnionFind) Count() int {
	return u.count
}

func (u UnionFind) prepareDump() []string {
	m := make(map[interface{}][]interface{})

	for k, v := range u.sets {
		parent := u.Find(v.parent)

		if _, ok := m[parent]; !ok {
			m[parent] = []interface{}{}
		}
		m[parent] = append(m[parent], k)
	}

	s := []string{}
	for _, e := range m {
		s = append(s, fmt.Sprintf("%v", e))
	}
	return s
}

// String returns a text representation of a UnionFind data structure.
func (u UnionFind) String() string {
	return strings.Join(u.prepareDump(), " ")
}
