package main

import (
	"strings"
	"testing"
)

func TestMooDefault(t *testing.T) {
	out := moo("Cow")
	if !strings.Contains(out, "Cow says moo!") {
		t.Errorf("expected 'Cow says moo!' in output, got:\n%s", out)
	}
}

func TestMooCustomAnimal(t *testing.T) {
	out := moo("Goat")
	if !strings.Contains(out, "Goat says moo!") {
		t.Errorf("expected 'Goat says moo!' in output, got:\n%s", out)
	}
}

func TestMooContainsArt(t *testing.T) {
	out := moo("Cow")
	if !strings.Contains(out, "(oo)") {
		t.Errorf("expected cow art in output, got:\n%s", out)
	}
}
