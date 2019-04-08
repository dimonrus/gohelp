package gohelp

import (
	"fmt"
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
