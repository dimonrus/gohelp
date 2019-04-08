package gohelp

import (
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

func TestToCamelCase(t *testing.T) {
	underscored := "some_underscore_name"
	str, err := ToCamelCase(underscored, true)
	if err != nil {
		t.Fatal(err)
	}

	if str != "SomeUnderscoreName" {
		t.Fatal("Incorrect convertation")
	}

	str, err = ToCamelCase(underscored, false)
	if err != nil {
		t.Fatal(err)
	}

	if str != "someUnderscoreName" {
		t.Fatal("Incorrect convertation")
	}
}

func TestBeforeString(t *testing.T) {
	str := "user/local/go/golkp/app/web/api/system"
	root := BeforeString(str, "golkp")

	if root != "user/local/go/" {
		t.Fatal("incorrect before logic")
	}

}
