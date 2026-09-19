package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	type testResult struct {
		data   int
		answer int
	}

	testResults := []testResult{
		testResult{10, 70},
		testResult{2, 14},
		testResult{3, 21},
	}

	for _, v := range testResults {
		got := Years(v.data)
		if got != v.answer {
			t.Error("expected ", v.answer, "got", got)
		}
	}

}

func TestTwoYears(t *testing.T) {
	type testResult struct {
		data   int
		answer int
	}

	testResults := []testResult{
		testResult{10, 70},
		testResult{2, 14},
		testResult{3, 21},
	}

	for _, v := range testResults {
		got := YearsTwo(v.data)
		if got != v.answer {
			t.Error("expected ", v.answer, "got", got)
		}
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	//Output:
	//70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(3))
	//Output:
	//21
}

func BenchmarkYears(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
