package gohelp

import (
	"fmt"
	"testing"
)

func TestIndex(t *testing.T) {
	t.Run("index_1", func(t *testing.T) {
		var slice = []int32{10, 20, 20, 30, 40, 40, 70}
		if Index(20, slice) != 1 {
			t.Fatal("index_1 index must be 1")
		}
	})

	t.Run("index_m1", func(t *testing.T) {
		var slice = []int32{10, 20, 20, 30, 40, 40, 70}
		if Index(21, slice) != -1 {
			t.Fatal("index_m1 index must be -1")
		}
	})

	t.Run("index_6", func(t *testing.T) {
		var slice = []int32{10, 20, 20, 30, 40, 40, 70}
		if Index(70, slice) != 6 {
			t.Fatal("index_6 index must be 6")
		}
	})

}

func TestExistsInArray(t *testing.T) {
	t.Run("int32", func(t *testing.T) {
		var slice = []int32{10, 20, 20, 30, 40, 40, 70}
		if !ExistsInArray(int32(10), slice) {
			t.Fatal("Must exists")
		}
	})

	t.Run("uint32", func(t *testing.T) {
		var slice = []uint32{10, 20, 20, 30, 40, 40, 70}
		if !ExistsInArray(uint32(20), slice) {
			t.Fatal("Must exists")
		}
	})

	t.Run("string", func(t *testing.T) {
		var slice = []string{"any", "foo", "bar", "baz"}
		if !ExistsInArray("baz", slice) {
			t.Fatal("Must exists")
		}
	})
}

func BenchmarkImplode(b *testing.B) {
	data := []int{10, 20, 30, 40, 50}
	for i := 0; i < b.N; i++ {
		Implode(data, ",")
	}
	b.ReportAllocs()
}

func TestImplode(t *testing.T) {
	data := []string{"10", "20", "30", "40", "50"}
	if Implode(data, ",") != "10,20,30,40,50" {
		t.Fatal("Wrong logic")
	}
}

func TestArrayIntersect(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		data := []string{"10", "20", "30", "40", "50"}
		dataTarget := []string{"30", "50"}

		var intersect, removed = ArrayIntersect(data, dataTarget)
		if removed != 3 {
			t.Fatal("count must be 2")
		}
		if len(intersect) != 2 {
			t.Fatal("len must be 2")
		}
		if intersect[0] != "30" {
			t.Fatal("must be a 30")
		}
	})

	t.Run("int64", func(t *testing.T) {
		data := []int64{10, 20, 30, 40, 50}
		dataTarget := []int64{30, 50}

		var intersect, removed = ArrayIntersect[int64](data, dataTarget)
		if removed != 3 {
			t.Fatal("count must be 2")
		}
		if len(intersect) != 2 {
			t.Fatal("len must be 2")
		}
		if intersect[0] != 30 {
			t.Fatal("must be a 30")
		}
	})

	t.Run("uint32", func(t *testing.T) {
		data := []uint32{10, 20, 30, 40, 50}
		dataTarget := []uint32{30, 50}

		var intersect, removed = ArrayIntersect[uint32](data, dataTarget)
		if removed != 3 {
			t.Fatal("count must be 2")
		}
		if len(intersect) != 2 {
			t.Fatal("len must be 2")
		}
		if intersect[0] != 30 {
			t.Fatal("must be a 30")
		}
	})
}

func TestAppendUnique(t *testing.T) {
	t.Run("uint32", func(t *testing.T) {
		data := []uint32{10, 20, 30, 40, 50}
		dataTarget := []uint32{30, 50}

		var unique = AppendUnique[uint32](dataTarget, data...)
		if len(unique) != 5 {
			t.Fatal("len must be 5")
		}
		if unique[0] != 30 {
			t.Fatal("must be a 10")
		}
		if unique[4] != 40 {
			t.Fatal("must be a 10")
		}
	})

	t.Run("string", func(t *testing.T) {
		data := []string{"10", "20", "30", "40", "50"}
		dataTarget := []string{"30", "50"}

		var unique = AppendUnique[string](dataTarget, data...)
		if len(unique) != 5 {
			t.Fatal("len must be 5")
		}
		if unique[0] != "30" {
			t.Fatal("must be a 10")
		}
		if unique[4] != "40" {
			t.Fatal("must be a 10")
		}
	})

	t.Run("int64", func(t *testing.T) {
		data := []int64{10, 20, 30, 40, 50}
		dataTarget := []int64{30, 50}

		var unique = AppendUnique[int64](dataTarget, data...)
		if len(unique) != 5 {
			t.Fatal("len must be 5")
		}
		if unique[0] != 30 {
			t.Fatal("must be a 10")
		}
		if unique[4] != 40 {
			t.Fatal("must be a 10")
		}
	})

	t.Run("uintptr", func(t *testing.T) {
		data := []uintptr{10, 20, 30, 40, 50}
		dataTarget := []uintptr{30, 50}

		var unique = AppendUnique[uintptr](dataTarget, data...)
		if len(unique) != 5 {
			t.Fatal("len must be 5")
		}
		if unique[0] != 30 {
			t.Fatal("must be a 10")
		}
		if unique[4] != 40 {
			t.Fatal("must be a 10")
		}
	})

	t.Run("byte", func(t *testing.T) {
		data := []byte{10, 20, 30, 40, 50}
		dataTarget := []byte{30, 50}

		var unique = AppendUnique[byte](dataTarget, data...)
		if len(unique) != 5 {
			t.Fatal("len must be 5")
		}
		if unique[0] != 30 {
			t.Fatal("must be a 10")
		}
		if unique[4] != 40 {
			t.Fatal("must be a 10")
		}
	})

	t.Run("rune", func(t *testing.T) {
		data := []rune{10, 20, 30, 40, 50}
		dataTarget := []rune{30, 50}

		var unique = AppendUnique[rune](dataTarget, data...)
		if len(unique) != 5 {
			t.Fatal("len must be 5")
		}
		if unique[0] != 30 {
			t.Fatal("must be a 10")
		}
		if unique[4] != 40 {
			t.Fatal("must be a 10")
		}
	})
}

func TestUniqueLeft(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		data := []int{10, 10, 30, 30}
		dataTarget := []int{30, 30, 50}

		var unique = UniqueLeft[int](dataTarget, data)
		if len(unique) != 1 {
			t.Fatal("len must be 1")
		}
		if unique[0] != 50 {
			t.Fatal("must be a 50")
		}
	})
}

func TestUnique(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		data := []int{10, 10, 30, 30}
		dataTarget := []int{30, 30, 50}

		var unique = Unique[int](dataTarget, data)
		if len(unique) != 3 {
			t.Fatal("len must be 3")
		}
		if unique[0] != 30 {
			t.Fatal("must be a 30")
		}
	})
}

func BenchmarkUnique(b *testing.B) {
	data := []int{10, 10, 30, 30}
	dataTarget := []int{30, 30, 50}
	for i := 0; i < b.N; i++ {
		Unique[int](dataTarget, data)
	}
	b.ReportAllocs()
}

func TestFilter(t *testing.T) {
	type My struct {
		One int
		Two bool
	}
	var bb = []My{{10, true}, {20, false}, {30, false}, {40, true}}
	result := Filter(bb, func(s My) bool {
		return s.Two
	})
	if len(result) != 2 {
		t.Fatal("logic is incorrect")
	}
}

func TestReduce(t *testing.T) {
	type My struct {
		One int
		Two bool
	}
	var bb = []My{{10, true}, {20, false}, {30, false}, {40, true}}
	result := Reduce(1, bb, func(d int, s My) int {
		return d * s.One
	})
	if result != 240000 {
		t.Fatal("reduce is not properly works")
	}
}

func TestMap(t *testing.T) {
	type My struct {
		One int
		Two bool
	}
	var bb = []My{{10, true}, {20, false}, {30, false}, {40, true}}
	result := Map(bb, func(s My) int {
		return s.One + 1
	})

	if len(result) != 4 {
		t.Fatal("result is wrong")
	}
	if result[0] != 11 {
		t.Fatal("result is wrong")
	}
}

func BenchmarkMap(b *testing.B) {
	type My struct {
		One int
		Two bool
	}
	var bb = []My{{10, true}, {20, false}, {30, false}, {40, true}}
	for i := 0; i < b.N; i++ {
		Map(bb, func(s My) int {
			return s.One + 1
		})
	}
	b.ReportAllocs()
}

func TestEach(t *testing.T) {
	type My struct {
		One int
		Two bool
	}
	var bb = []My{{10, true}, {20, false}, {30, false}, {40, true}}
	Each(bb, func(s My) {
		fmt.Println(s.One, s.Two)
	})
}

func uniqueCompare(a, b []int32) (common []int32, removed []int32, added []int32) {
	if len(b) == 0 {
		return
	}
	for _, va := range a {
		var found bool
		for _, vb := range b {
			if va == vb {
				common = AppendUnique(common, va)
				found = true
				break
			}
		}
		if !found {
			removed = AppendUnique(removed, va)
		}
	}
	for _, vb := range b {
		var found bool
		for _, va := range a {
			if va == vb {
				found = true
				break
			}
		}
		if !found {
			added = AppendUnique(added, vb)
		}
	}
	return
}

func TestUniqueCompare(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		a := []int32{10, 2, 3, 4, 7, 8, 9, 2, 3}
		b := []int32{1, 3, 44, 11, 9, 11}
		a1, b1, c1 := uniqueCompare(a, b)
		t.Log(a1, b1, c1)
		a, b, c := UniqueCompare(a, b)
		t.Log(a, b, c)
		if len(a1) != len(a) {
			t.Fatal("wrong common")
		}
		if a1[0] != a[0] {
			t.Fatal("wrong common first")
		}
		if a1[1] != a[1] {
			t.Fatal("wrong common last")
		}

		if len(b1) != len(b) {
			t.Fatal("wrong removed")
		}
		if b1[0] != b[0] {
			t.Fatal("wrong removed first")
		}
		if b1[4] != b[4] {
			t.Fatal("wrong removed last")
		}

		if len(c1) != len(c) {
			t.Fatal("wrong added")
		}
		if c1[0] != c[0] {
			t.Fatal("wrong added first")
		}
		if c1[2] != c[2] {
			t.Fatal("wrong added last")
		}
	})
	t.Run("left", func(t *testing.T) {
		a := []int32{10, 2}
		b := []int32{}
		a1, b1, c1 := uniqueCompare(a, b)
		t.Log(a1, b1, c1)
		a, b, c := UniqueCompare(a, b)
		t.Log(a, b, c)
		if len(a1) != len(a) {
			t.Fatal("wrong common")
		}
		if len(a) != 0 {
			t.Fatal("wrong common")
		}
		if len(b) != 0 {
			t.Fatal("wrong removed")
		}
		if len(c) != 0 {
			t.Fatal("wrong added")
		}
	})
	t.Run("right", func(t *testing.T) {
		a := []int32{}
		b := []int32{10, 2}
		a1, b1, c1 := uniqueCompare(a, b)
		t.Log(a1, b1, c1)
		a, b, c := UniqueCompare(a, b)
		t.Log(a, b, c)
		if len(a1) != len(a) {
			t.Fatal("wrong common")
		}
		if len(a) != 0 {
			t.Fatal("wrong common")
		}
		if len(b) != 0 {
			t.Fatal("wrong removed")
		}
		if len(c) != 2 {
			t.Fatal("wrong added")
		}
	})
	t.Run("all unique", func(t *testing.T) {
		a := []int32{1}
		b := []int32{10, 2, 3}
		a1, b1, c1 := uniqueCompare(a, b)
		t.Log(a1, b1, c1)
		a, b, c := UniqueCompare(a, b)
		t.Log(a, b, c)
		if len(a1) != len(a) {
			t.Fatal("wrong common")
		}
		if len(a) != 0 {
			t.Fatal("wrong common")
		}
		if len(b) != 1 {
			t.Fatal("wrong removed")
		}
		if len(c) != 3 {
			t.Fatal("wrong added")
		}
	})
	t.Run("equal", func(t *testing.T) {
		a := []int32{10}
		b := []int32{10}
		a1, b1, c1 := uniqueCompare(a, b)
		t.Log(a1, b1, c1)
		a, b, c := UniqueCompare(a, b)
		t.Log(a, b, c)
		if len(a1) != len(a) {
			t.Fatal("wrong common")
		}
		if len(a) != 1 {
			t.Fatal("wrong common")
		}
		if len(b) != 0 {
			t.Fatal("wrong removed")
		}
		if len(c) != 0 {
			t.Fatal("wrong added")
		}
	})
}

// goos: darwin
// goarch: amd64
// pkg: github.com/dimonrus/gohelp
// cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
// BenchmarkUniqueCompare
// BenchmarkUniqueCompare-8   	 7682722	       166.6 ns/op	     192 B/op	       1 allocs/op
func BenchmarkUniqueCompare(b *testing.B) {
	a1 := []int{10, 2, 3, 4, 7, 8, 9, 2, 3}
	b1 := []int{1, 3, 44, 11, 9, 11}
	for i := 0; i < b.N; i++ {
		UniqueCompare(a1, b1)
	}
	b.ReportAllocs()
}
