package main

import "testing"

func TestSimple(t *testing.T) {
	if 1+1 != 2 {
		t.Error("math is broken")
	}
}
