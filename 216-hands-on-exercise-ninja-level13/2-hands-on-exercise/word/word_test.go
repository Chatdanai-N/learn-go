package word

import (
	"fmt"
	"mymodule/216-hands-on-exercise-ninja-level13/2-hands-on-exercise/quote"
	"testing"
)

func TestUseCount(t *testing.T) {

	m := UseCount("one two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "three":
			if v != 3 {
				t.Error("got", v, "want", 3)
			}
		}
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	//Output:
	//3
}

func TestCount(t *testing.T) {

	type TestResult struct {
		data   string
		answer int
	}

	tests := []TestResult{
		TestResult{data: "one two three", answer: 3},
		TestResult{data: "Hello World", answer: 2},
		TestResult{data: "Spain is winner of world's cup", answer: 6},
	}

	for _, v := range tests {
		n := Count(v.data)
		if n != v.answer {
			t.Error("expected", v.answer, "got", n)
		}
	}
}

func ExampleUseCount() {
	fmt.Println(UseCount("one two three"))
	//Output:
	//map[one:1 three:1 two:1]
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
