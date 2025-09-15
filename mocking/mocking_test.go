package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buf := &bytes.Buffer{}

	Countdown(buf)

	got := buf.String()
	want := "3"

	if got != want {
		t.Fatalf("want %q but got %q", want, got)
	}
}
