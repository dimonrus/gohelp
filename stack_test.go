package gohelp

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		fst := "first"
		scd := "second"
		trd := "third"

		var s = Stack[string]{}.Push(fst).Push(scd).Push(trd)
		v, ok, s := s.Pop()
		if !ok || v != "third" {
			t.Fatal("must be third")
		}
		v, ok, s = s.Pop()
		if !ok || v != "second" {
			t.Fatal("must be second")
		}
		v, ok, s = s.Pop()
		if !ok || v != "first" {
			t.Fatal("must be first")
		}
		v, ok, s = s.Pop()
		if ok {
			t.Fatal("must be not ok")
		}
		if v != "" {
			t.Fatal("wrong value")
		}
	})

	t.Run("custom", func(t *testing.T) {
		type Custom struct {
			Foo int
			Bar string
		}
		c1 := Custom{
			Foo: 1,
			Bar: "bar",
		}
		c2 := Custom{
			Foo: 2,
			Bar: "baz",
		}
		c3 := Custom{
			Foo: 3,
			Bar: "fuz",
		}
		var s = Stack[Custom]{}.Push(c1).Push(c2).Push(c3)
		v, ok, s := s.Pop()
		if !ok || v.Bar != "fuz" {
			t.Fatal("must be fuz")
		}
	})

}

func BenchmarkStackByte(b *testing.B) {
	s := Stack[[]byte]{}
	var v = []byte("first")
	for i := 0; i < b.N; i++ {
		s = s.Push(v)
		_, _, s = s.Pop()
	}
	b.ReportAllocs()
}
