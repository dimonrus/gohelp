package gohelp

import (
	"fmt"
	"testing"
)

func TestStackByte(t *testing.T) {
	s := ByteStack{}
	fst := []byte("first")
	scd := []byte("second")
	thr := []byte("third")
	s = s.Push(fst)
	s = s.Push(scd)

	var v []byte
	v, s = s.Pop()

	fmt.Printf("%s", v)

	for i := 0; i < 10; i++ {
		s = s.Push(thr)
		v, s = s.Pop()
		if v[0] != thr[0] {
			t.Fatal("wrong")
		}
	}
}

func BenchmarkStackByte(b *testing.B) {
	s := ByteStack{}
	var v = []byte("first")
	for i := 0; i < b.N; i++ {
		s = s.Push(v)
		v, s = s.Pop()
	}
	b.ReportAllocs()
	fmt.Printf("%s\n",v)
}