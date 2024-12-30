package unionfind_test

import (
	"math/rand/v2"
	"testing"

	. "github.com/pietv/unionfind"
)

type union [2]interface{}
type sets []interface{}
type conn struct {
	a, b interface{}
	want bool
}

var BasicTests = []struct {
	name   string
	in     sets
	unions []union
	count  int
}{
	{"1", sets{}, nil, 0},
	{"2", sets{1}, nil, 1},
	{"3", sets{1, 1}, nil, 1},
	{"4", sets{1, 2}, nil, 2},
	{"5", sets{"one"}, nil, 1},
	{"6", sets{"one", "two"}, nil, 2},

	{"7", sets{1}, []union{{1, 1}}, 1},
	{"8", sets{1, 2}, []union{{1, 2}}, 1},
	{"9", sets{1, 2}, []union{{1, 1}, {2, 2}}, 2},
	{"10", sets{1, 2}, []union{{1, 2}}, 1},
	{"11", sets{1, 2}, []union{{2, 1}}, 1},
	{"12", sets{"one", "two"}, []union{{"one", "two"}}, 1},
	{"13", sets{"one", "two"}, []union{{"two", "one"}}, 1},

	{"14", sets{1, 2, 3, 4}, []union{{1, 2}, {3, 4}}, 2},
	{"15", sets{1, 2, 3, 4}, []union{{2, 3}, {3, 4}}, 2},
	{"16", sets{1, 2, 3, 4}, []union{{1, 3}, {2, 4}, {2, 3}}, 1},
	{"17", sets{1, 2, 3, 4}, []union{{2, 4}}, 3},
	{"18", sets{1, 2, 3, 4}, []union{{2, 3}, {3, 2}, {2, 3}}, 3},
}

func setup(in sets, unions []union) *UnionFind {
	u := New()

	for _, elem := range in {
		u.MakeSet(elem)
	}

	for _, elem := range unions {
		u.Union(elem[0], elem[1])
	}

	return u
}

func TestBasic(t *testing.T) {
	for _, test := range BasicTests {
		if actual := setup(test.in, test.unions).Count(); actual != test.count {
			t.Errorf("%q: got %v, want %v", test.name, actual, test.count)
		}
	}
}

var ConnectedTests = []struct {
	name   string
	in     sets
	unions []union
	conn   conn
}{
	{"1", sets{1}, nil, conn{1, 1, true}},
	{"2", sets{1, 2}, nil, conn{2, 2, true}},
	{"3", sets{1, 2}, nil, conn{1, 2, false}},
	{"4", sets{1, 2}, []union{{1, 2}}, conn{1, 2, true}},

	{"5", sets{1, 2, 3, 4, 5}, []union{{1, 2}, {3, 5}, {2, 3}}, conn{1, 5, true}},
	{"6", sets{1, 2, 3, 4, 5}, []union{{1, 2}, {2, 3}, {5, 3}}, conn{1, 4, false}},
	{"7", sets{1, 2, 3, 4, 5}, nil, conn{1, 5, false}},
	{"8", sets{1, 2, 3, 4, 5}, []union{{5, 4}, {4, 3}}, conn{5, 3, true}},
}

func TestConnected(t *testing.T) {
	for _, test := range ConnectedTests {
		if actual := setup(test.in, test.unions).Connected(test.conn.a, test.conn.b); actual != test.conn.want {
			t.Errorf("%q: got %v, want %v", test.name, actual, test.conn.want)
		}
	}
}

func TestMakeSet(t *testing.T) {
	// Ignore rather than spread panic.
	u := New()
	u.MakeSet()
	u.MakeSet(nil)
	if actual := u.Count(); actual != 0 {
		t.Errorf("empty: got %v, want 0", actual)

	}

	// Multiple sets.
	u = New()
	u.MakeSet(1, 2, 3)
	if actual := u.Count(); actual != 3 {
		t.Errorf("multiple: got %v, want 3", actual)

	}

	// Repeated sets.
	u = New()
	u.MakeSet(1, 1, 2, 2, 1, 2)
	u.MakeSet(2, 2, 4, 2, 4, 3)
	if actual := u.Count(); actual != 4 {
		t.Errorf("repeated: got %v, want 4", actual)
	}
}

func TestFind(t *testing.T) {
	ds := New()

	// Non-existent sets.
	if ds.Find(1) != nil {
		t.Errorf("1: expected set not to exist")
	}
	if ds.Find(2) != nil {
		t.Errorf("2: expected set not to exist")
	}
}

func TestUnion(t *testing.T) {
	ds := New()

	ds.Union(3, 4)

	if ds.Find(3) != 3 {
		t.Errorf("3: expected set to exist")
	}
	if ds.Find(4) != 3 {
		t.Errorf("4: expected set to exist")
	}

	if ds.Connected(3, 4) != true {
		t.Errorf("expected sets 3 and 4 to be connected")
	}
}

func TestExists(t *testing.T) {
	u := New()
	u.MakeSet(nil)
	if actual := u.Exists(nil); actual != false {
		t.Errorf("nil: got %v, want false", actual)
	}

	u = New()
	if actual := u.Exists(1); actual != false {
		t.Errorf("doesn't exist: got %v, want false", actual)
	}
	u.MakeSet(1)
	if actual := u.Exists(1); actual != true {
		t.Errorf("exists: got %v, want true", actual)
	}
}

func TestCount(t *testing.T) {
	u := New()

	if actual := u.Count(); actual != 0 {
		t.Errorf("empty: got %v, want 0", actual)
	}

	u.MakeSet()
	if actual := u.Count(); actual != 0 {
		t.Errorf("nothing made: got %v, want 0", actual)
	}

	u.MakeSet(nil)
	if actual := u.Count(); actual != 0 {
		t.Errorf("nil made: got %v, want 0", actual)
	}
}

func TestString(t *testing.T) {
	u := New()

	if actual := u.String(); actual != "" {
		t.Errorf("empty: got %q, want %q", actual, "")
	}

	u.MakeSet(1, 2)
	if actual := u.String(); actual != "[1] [2]" && actual != "[2] [1]" {
		t.Errorf("got %q, want %q", actual, "[1] [2]")
	}
}

func BenchmarkManyUnions(b *testing.B) {
	b.StopTimer()
	u := New()

	a := make([]int, 0)
	for i := 0; i < b.N; i++ {
		n := rand.Int()
		u.MakeSet(n)
		a = append(a, n)

		n1 := a[rand.IntN(len(a))]
		n2 := a[rand.IntN(len(a))]
		b.StartTimer()
		u.Union(n1, n2)
		b.StopTimer()
	}
}

func BenchmarkManyCheckConnects(b *testing.B) {
	b.StopTimer()
	u := New()

	a := make([]int, 0)
	for i := 0; i < b.N; i++ {
		n := rand.Int()
		u.MakeSet(n)
		a = append(a, n)

		// Random union.
		n1 := a[rand.IntN(len(a))]
		n2 := a[rand.IntN(len(a))]
		u.Union(n1, n2)

		// Random check.
		n1 = a[rand.IntN(len(a))]
		n2 = a[rand.IntN(len(a))]
		b.StartTimer()
		_ = u.Connected(n1, n2)
		b.StopTimer()
	}
}

func TestRandomShuffle10(t *testing.T) {
	for k := 1; k < 10; k++ {
		// Cut a shuffled slice with distinct integers in two.
		x := rand.Perm(2000)
		a, b := x[:1000], x[1000:]

		u := New()
		for i := range a {
			u.MakeSet(a[i])
			u.MakeSet(b[i])
		}

		for i := range a[:len(a)-1] {
			// Connect pairwise.
			u.Union(a[i], a[i+1])
			u.Union(b[i], b[i+1])
		}

		// Number of connected components should be 2.
		if actual := u.Count(); actual != 2 {
			t.Errorf("random slices: got %v, want 2", actual)
		}
	}
}

func TestRandomConnect10(t *testing.T) {
	for k := 0; k < 10; k++ {
		u := New()
		x := rand.Perm(2000)

		for _, elem := range x {
			u.MakeSet(elem)
		}

		// Connect pairwise from the first to the last.
		for i := range x[:len(x)-1] {
			u.Union(x[i], x[i+1])
		}

		// Check that the first and the last are connected.
		if actual := u.Count(); actual != 1 {
			t.Errorf("random connect: got %v, want 1", actual)
		}
	}
}
