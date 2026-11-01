package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

// Keep the original scaffold assertion without retaining a placeholder product
// entry point. This helper is only used by the preserved pre-existing test.
func greeting(w io.Writer) error {
	_, err := fmt.Fprintln(w, "Hello, World!")
	return err
}

func TestGreeting(t *testing.T) {
	var buf bytes.Buffer
	err := greeting(&buf)
	if err != nil {
		t.Fatalf("greeting() error = %v, want nil", err)
	}

	got := buf.String()
	want := "Hello, World!\n"

	if got != want {
		t.Errorf("greeting() = %q, want %q", got, want)
	}
}
