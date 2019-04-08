package gohelp

import "testing"

func TestAppendUniqueInt(t *testing.T) {
	var ints []int
	ints = AppendUniqueInt(ints, 35)
	ints = AppendUniqueInt(ints, 35)
	if len(ints) > 1 {
		t.Fatal("AppendUniqueInt works incorrect")
	}
}

func TestAppendUniqueInt64(t *testing.T) {
	var ints []int64
	ints = AppendUniqueInt64(ints, 35)
	ints = AppendUniqueInt64(ints, 35)
	if len(ints) > 1 {
		t.Fatal("AppendUniqueInt64 works incorrect")
	}
}

func TestAppendUniqueString(t *testing.T) {
	var strs []string
	strs = AppendUniqueString(strs, "value")
	strs = AppendUniqueString(strs, "value")
	if len(strs) > 1 {
		t.Fatal("AppendUniqueString works incorrect")
	}
}

func TestAppendUniqueUint64(t *testing.T) {
	var uints []uint64
	uints = AppendUniqueUint64(uints, 313)
	uints = AppendUniqueUint64(uints, 313)
	if len(uints) > 1 {
		t.Fatal("AppendUniqueUint64 works incorrect")
	}
}

func TestExistsInArrayInt(t *testing.T) {
	ints := []int{10, 20}
	exists := ExistsInArrayInt(10, ints)
	if exists != true {
		t.Fatal("ExistsInArrayInt works incorrect")
	}
}

func TestExistsInArrayString(t *testing.T) {
	stringsArray := []string{"10", "20"}
	exists := ExistsInArrayString("10", stringsArray)
	if exists != true {
		t.Fatal("ExistsInArrayString works incorrect")
	}
}

func TestExistsInArrayInt64(t *testing.T) {
	ints64 := []int64{10, 20}
	exists := ExistsInArrayInt64(10, ints64)
	if exists != true {
		t.Fatal("ExistsInArrayInt64 works incorrect")
	}
}
