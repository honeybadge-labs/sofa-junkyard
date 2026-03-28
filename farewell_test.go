package main

import "testing"

func TestFarewellBasic(t *testing.T) {
	out := Farewell("Alice")
	expected := "Goodbye, Alice!"
	if out != expected {
		t.Errorf("expected %q, got %q", expected, out)
	}
}

func TestFarewellDifferentName(t *testing.T) {
	out := Farewell("Bob")
	expected := "Goodbye, Bob!"
	if out != expected {
		t.Errorf("expected %q, got %q", expected, out)
	}
}
