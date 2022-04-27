package gohelp

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	uuid := UUID{}

	if uuid.Generate().Get() == nil {
		t.Errorf("Cant create UUID")
	}

	if *uuid.Generate().Get() == *uuid.Generate().Get() {
		t.Errorf("Same uuid generated")
	}
}

func TestToUnderscore(t *testing.T) {
	s := "camelCaseString"
	if ToUnderscore(s) != "camel_case_string" {
		t.Fatal("to underscore is not works")
	}
}

func BenchmarkToUnderscore(b *testing.B) {
	s := "camelCaseString"
	for i := 0; i < b.N; i++ {
		ToUnderscore(s)
	}
	b.ReportAllocs()
}

func TestToCamelCase(t *testing.T) {
	underscored := "some_underscore_name"
	str := ToCamelCase(underscored, true)
	if str != "SomeUnderscoreName" {
		t.Fatal("Incorrect convertation")
	}

	str = ToCamelCase(underscored, false)
	if str != "someUnderscoreName" {
		t.Fatal("Incorrect convertation")
	}

	underscored = "__som_e_underscore_name_"
	str = ToCamelCase(underscored, false)
	if str != "somEUnderscoreName" {
		t.Fatal("Incorrect convertation")
	}
}

func BenchmarkToCamelCase(b *testing.B) {
	underscored := "some_underscore_name"
	for i := 0; i < b.N; i++ {
		ToCamelCase(underscored, false)
	}
	b.ReportAllocs()
}

func TestBeforeString(t *testing.T) {
	str := "user/local/go/golkp/app/web/api/system"
	root := BeforeString(str, "golkp")

	if root != "user/local/go/" {
		t.Fatal("incorrect before logic")
	}
}

func TestColour(t *testing.T) {
	fmt.Printf(AnsiBackgroundCustom+AnsiBlue+AnsiReversed+"%s"+AnsiReset, 90, "color blue")
}

func TestUUID_IsValid(t *testing.T) {
	uuid := (&UUID{}).Generate()

	fmt.Println(*uuid.Get())
	isOk := uuid.IsValid()
	if !isOk {
		t.Fatal("Uuid is invalid")
	}
}
