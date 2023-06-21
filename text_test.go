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
	t.Run("simple", func(t *testing.T) {
		s := "camelCaseString"
		if ToUnderscore(s) != "camel_case_string" {
			t.Fatal("to underscore is not works")
		}
	})
	t.Run("digits", func(t *testing.T) {
		s := "camel11Case22String"
		if ToUnderscore(s) != "camel11_case22_string" {
			t.Log(ToUnderscore(s))
			t.Fatal("to underscore is not works")
		}
	})
	t.Run("long_line", func(t *testing.T) {
		s := "someLongSuperLineprovidedRightNow"
		if ToUnderscore(s) != "some_long_super_lineprovided_right_now" {
			t.Log(ToUnderscore(s))
			t.Fatal("to underscore is not works")
		}
	})
}

func TestToUnderscoreToCamelCase(t *testing.T) {
	t.Run("forward", func(t *testing.T) {
		s := "camelCaseString"
		str := ToUnderscore(s)
		if ToCamelCase(str, false) != s {
			t.Fatal("forward")
		}
	})
	t.Run("backward", func(t *testing.T) {
		s := "came_case_string"
		str := ToCamelCase(s, false)
		if ToUnderscore(str) != s {
			t.Fatal("backward")
		}
	})
}

func BenchmarkToUnderscore(b *testing.B) {
	// goos: darwin
	// goarch: arm64
	// pkg: github.com/dimonrus/gohelp
	// BenchmarkToUnderscore
	// BenchmarkToUnderscore/simple
	// BenchmarkToUnderscore/simple-12         	36219216	        32.07 ns/op	      24 B/op	       1 allocs/op
	b.Run("simple", func(b *testing.B) {
		s := "camelCaseString"
		for i := 0; i < b.N; i++ {
			ToUnderscore(s)
		}
		b.ReportAllocs()
	})
	// goos: darwin
	// goarch: arm64
	// pkg: github.com/dimonrus/gohelp
	// BenchmarkToUnderscore
	// BenchmarkToUnderscore/long
	// BenchmarkToUnderscore/long-12         	21209469	        55.24 ns/op	      48 B/op	       1 allocs/op
	b.Run("long", func(b *testing.B) {
		s := "someLongSuperLineprovidedRightNow"
		for i := 0; i < b.N; i++ {
			ToUnderscore(s)
		}
		b.ReportAllocs()
	})

}

func TestToCamelCase(t *testing.T) {
	underscored := "some_underscore_name"
	str := ToCamelCase(underscored, true)
	if str != "SomeUnderscoreName" {
		t.Fatal("Incorrect convertation")
	}

	underscored = "SomeName"
	str = ToCamelCase(underscored, true)
	if str != "SomeName" {
		t.Fatal("Incorrect convertation")
	}

	underscored = "some_Name1"
	str = ToCamelCase(underscored, true)
	if str != "SomeName1" {
		t.Fatal("Incorrect convertation")
	}

	underscored = "so23me_44_name1"
	str = ToCamelCase(underscored, true)
	if str != "So23me44Name1" {
		t.Fatal("Incorrect convertation")
	}

	underscored = "SomeName"
	str = ToCamelCase(underscored, false)
	if str != "someName" {
		t.Fatal("Incorrect convertation")
	}

	underscored = "Some_Name"
	str = ToCamelCase(underscored, false)
	if str != "someName" {
		t.Fatal("Incorrect convertation")
	}

	underscored = "Some_Name"
	str = ToCamelCase(underscored, true)
	if str != "SomeName" {
		t.Fatal("Incorrect convertation")
	}

	underscored = "some_underscore_name"
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

func TestCheckBracers(t *testing.T) {
	t.Run("empty_ok", func(t *testing.T) {
		var str = ``
		err := CheckBracers(str, Stack[byte]{})
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("simple_ok", func(t *testing.T) {
		var str = `{}`
		err := CheckBracers(str, Stack[byte]{})
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("simple_circle_ok", func(t *testing.T) {
		var str = `(())`
		err := CheckBracers(str, Stack[byte]{})
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("simple_circle_nok", func(t *testing.T) {
		var str = `())`
		err := CheckBracers(str, Stack[byte]{})
		if err == nil {
			t.Fatal("must be an error of circle bracers count")
		}
	})
	t.Run("simple_nok", func(t *testing.T) {
		var str = `{`
		err := CheckBracers(str, Stack[byte]{})
		if err == nil {
			t.Fatal("must be an error simple_nok")
		}
	})
	t.Run("three_pair_ok", func(t *testing.T) {
		var str = `{{{}}}`
		err := CheckBracers(str, Stack[byte]{})
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("three_pair_nok", func(t *testing.T) {
		var str = `{{{}}}}`
		err := CheckBracers(str, Stack[byte]{})
		if err == nil {
			t.Fatal("must be an error simple_nok")
		}
	})
	t.Run("three_pair_square_ok", func(t *testing.T) {
		var str = `[[[]]]`
		err := CheckBracers(str, Stack[byte]{})
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("three_pair_square_nok", func(t *testing.T) {
		var str = `[[[[]]]`
		err := CheckBracers(str, Stack[byte]{})
		if err == nil {
			t.Fatal("must be an error three_pair_square_nok")
		}
	})
	t.Run("all_bracers_order_ok", func(t *testing.T) {
		var str = `{[{}]}`
		err := CheckBracers(str, Stack[byte]{})
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("all_bracers_wrong_order", func(t *testing.T) {
		var str = `{[{]}}`
		err := CheckBracers(str, Stack[byte]{})
		if err == nil {
			t.Fatal("must be an order error all_bracers_wrong_order")
		}
	})
	t.Run("all_bracers_with_data_order_ok", func(t *testing.T) {
		var str = `{x:y[1,2,3]{z}:foo}`
		err := CheckBracers(str, Stack[byte]{})
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("all_type_of_bracers_with_data_ok", func(t *testing.T) {
		var str = `{[1]{(2)[3]}}`
		err := CheckBracers(str, Stack[byte]{})
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("all_type_of_bracers_with_data_nok", func(t *testing.T) {
		var str = `{[1]]{(2)[3]}[}`
		err := CheckBracers(str, Stack[byte]{})
		if err == nil {
			t.Fatal("must be an error all_type_of_bracers_with_data_nok")
		}
	})
}

// goos: darwin
// goarch: arm64
// pkg: github.com/dimonrus/gohelp
// BenchmarkCheckBracers
// BenchmarkCheckBracers-12    	45314067	        25.92 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCheckBracers(b *testing.B) {
	var str = `{[1]{(2)[3]}}`
	stack := make(Stack[byte], 0, 32)
	for i := 0; i < b.N; i++ {
		err := CheckBracers(str, stack)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ReportAllocs()
}

func TestRandString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "length10", args: args{length: 10}, want: 10},
		{name: "length5", args: args{length: 5}, want: 5},
		{name: "length1", args: args{length: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandString(tt.args.length); len(got) != tt.want {
				t.Errorf("RandString() = %v, want %v", got, tt.want)
			}
		})
	}
}
