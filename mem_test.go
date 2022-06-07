package gohelp

import (
	"encoding/json"
	"testing"
)

type ConvertStruct struct {
	Foo int
	Bar *string
}

type DestStruct struct {
	Foo int `json:"foo"`
}

func TestConvert(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		tt := ConvertStruct{
			Foo: 1,
			Bar: Ptr("bar"),
		}
		result := Convert[ConvertStruct, DestStruct](tt)
		if result.Foo != 1 {
			t.Fatal("wrong convert")
		}
		data, err := json.Marshal(result)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(data))
		if string(data) != `{"foo":1}` {
			t.Fatal("wrong marshal")
		}
	})
	t.Run("data", func(t *testing.T) {
		data := "foobarbaz"
		bytes := Convert[string, []byte](data)
		if string(bytes) != "foobarbaz" {
			t.Fatal("wrong convert")
		}
	})
}
