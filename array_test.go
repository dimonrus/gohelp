package gohelp

import (
	"fmt"
	"testing"
)

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
