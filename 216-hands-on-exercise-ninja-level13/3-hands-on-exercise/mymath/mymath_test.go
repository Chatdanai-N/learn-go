package mymath

import (
	"fmt"
	"testing"
)

func ExampleCenteredAvg() {
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(CenteredAvg(xi))
	//Output:
	//3
}

func TestCenteredAvg(t *testing.T) {

	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{data: []int{1, 2, 3, 4, 5}, answer: 3},
		test{data: []int{10, 20, 30, 40}, answer: 25},
		test{data: []int{10, 20, 40, 60, 80}, answer: 40},
	}

	for _, v := range tests {
		got := CenteredAvg(v.data)
		if got != v.answer {
			t.Error("Expected", v.answer, "got", got)
		}
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}
