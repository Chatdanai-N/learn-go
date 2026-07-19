package main

import (
	"log"
	"testing"
)

func TestParadise(t *testing.T) {

	got := paradise("Hawaii")
	want := "My idea of paradise is Thailand"

	if got != want {
		log.Fatalf("Error paradise - want %v and got %v", want, got)
	}
}
