package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 3)
	if total != 10 {
		t.Errorf("Sum was incorrent, got %d, want: %d", total, 10)
	}
}
