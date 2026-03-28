package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCopyrightHeader(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		data, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.HasPrefix(string(data), "// Copyright Honeybadge Labs") {
			t.Errorf("%s: first line must be '// Copyright Honeybadge Labs' (got %q)", f, strings.SplitN(string(data), "\n", 2)[0])
		}
	}
}
