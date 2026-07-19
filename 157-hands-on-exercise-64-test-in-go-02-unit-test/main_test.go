package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := add(2, 4)
	want := 5
	if got != want {
		//log.Fatalf("Error add - want %v and got %v", want, got)
		t.Errorf("add values was incorrent, got : %d, want : %d ", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := subtract(4, 2)
	want := 1
	if got != want {
		//log.Fatalf("Error subtract - want %v and got %v", want, got)
		t.Errorf("subtract values was incorrent, got : %d, value : %d", got, want)
	}
}

func TestDoMath(t *testing.T) {
	gotAdd := doMath(3, 4, add)
	wantAdd := 6
	if gotAdd != wantAdd {
		//log.Fatalf("Error doMath add - want %v and got %v", wantAdd, gotAdd)
		t.Errorf("add values was incorrent, got : %d, value : %d ", gotAdd, wantAdd)
	}

	gotSubtract := doMath(5, 3, subtract)
	wantSubtract := 1
	if gotSubtract != wantSubtract {
		//log.Fatalf("Error doMath subtract - want %v and got %v", wantSubtract, gotSubtract)
		t.Errorf("subtract values was incorrent, got : %d, value : %d", gotSubtract, wantSubtract)
	}

}
