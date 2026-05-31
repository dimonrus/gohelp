package gohelp

import (
	"fmt"
	"math/big"
	"testing"
)

func TestGetRndNumber(t *testing.T) {
	number := GetRndNumber(1000, 2000)
	if number == 0 {
		t.Fatal("Number cant be 0")
	}
}

func TestGetRndPhone(t *testing.T) {
	phone := GetRndPhone()
	if phone == 0 {
		t.Fatal("Phone cant be 0")
	}
}

func TestGetRndId(t *testing.T) {
	id := GetRndId()
	if id == 0 {
		t.Fatal("Id cant be 0")
	}
}

func TestGetRndDateTime(t *testing.T) {
	dateTime := GetRndDateTime()
	fmt.Print(dateTime)
}

func TestRandInt64(t *testing.T) {
	i, err := RandInt64()
	if err != nil {
		t.Fatal(err)
	}
	if i == 0 {
		t.Fatal("Int64 cant be 0")
	}
}

func TestRndNum(t *testing.T) {
	t.Log(int64(1<<63 - 1))
	t.Log(-int64(1<<63 - 1))
	bi := big.NewInt(1<<63 - 1)
	t.Log(bi.Int64())
}

func BenchmarkGetRndNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRndNumber(1000, 2000)
	}
	b.ReportAllocs()
}
