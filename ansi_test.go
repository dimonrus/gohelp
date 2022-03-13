package gohelp

import "testing"

func TestRainbow(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		str := "Каждый охотник желает знать где сидит фазан"
		//t.Log(Rainbow(str, " "))
		result := Rainbow(str, ' ')
		t.Log(result)
	})

	t.Run("one_word", func(t *testing.T) {
		str := "Каждый"
		result := Rainbow(str, ' ')
		if len(result) != 25 {
			t.Fatal("two_word is not work's properly")
		}
		t.Log(result)
	})

	t.Run("one_word_left_sep", func(t *testing.T) {
		str := "  Каждый"
		result := Rainbow(str, ' ')
		if len(result) != 27 {
			t.Fatal("one_word_left_sep is not work's properly")
		}
		t.Log(result)
	})

	t.Run("one_word_right_sep", func(t *testing.T) {
		str := "Каждый  "
		result := Rainbow(str, ' ')
		if len(result) != 27 {
			t.Fatal("one_word_right_sep is not work's properly")
		}
		t.Log(result)
	})

	t.Run("two_word", func(t *testing.T) {
		str := "Каждый охотник"
		result := Rainbow(str, ' ')
		if len(result) != 55 {
			t.Fatal("two_word is not work's properly")
		}
		t.Log(result)
	})

	t.Run("two_word_sep", func(t *testing.T) {
		str := "  Каждый  охотник  "
		result := Rainbow(str, ' ')
		if len(result) != 60 {
			t.Fatal("two_word_sep is not work's properly")
		}
		t.Log(result)
	})

	t.Run("double", func(t *testing.T) {
		str := "Каждый охотник желает знать где сидит фазан Каждый охотник желает знать где сидит фазан"
		result := Rainbow(str, ' ')
		if len(result) != 357 {
			t.Fatal("double is not work's properly")
		}
		t.Log(result)
	})
}

func BenchmarkRainbow(b *testing.B) {
	b.Run("classic", func(b *testing.B) {
		str := "Каждый охотник желает знать где сидит фазан"
		for i := 0; i < b.N; i++ {
			Rainbow(str, ' ')
		}
		b.ReportAllocs()
	})
}
