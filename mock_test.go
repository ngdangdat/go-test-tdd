package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

var _ Sleeper = &SpySleeper{}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	Countdown(buffer, &SpySleeper{})
	want := `3
2
1
Go!`
	got := buffer.String()
	if want != got {
		t.Errorf("want %q got %q", want, got)
	}
}
