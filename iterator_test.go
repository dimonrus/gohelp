package gohelp

import "testing"

func BenchmarkNewIterator(b *testing.B) {
	var it = NewIterator(100)
	for i := 0; i < b.N; i++ {
		for it.Next() {
		}
		it.Reset()
	}
	b.ReportAllocs()
}
