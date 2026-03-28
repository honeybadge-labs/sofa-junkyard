package main

import "testing"

func TestFarewell(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"World", "Goodbye, World!"},
		{"Alice", "Goodbye, Alice!"},
	}
	for _, tt := range tests {
		got := Farewell(tt.name)
		if got != tt.expected {
			t.Errorf("Farewell(%q) = %q, want %q", tt.name, got, tt.expected)
		}
	}
}
