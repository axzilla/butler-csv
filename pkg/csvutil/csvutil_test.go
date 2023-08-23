package csvutil

import "testing"

func TestMakeNegativ(t *testing.T) {
	got := MakeNegative("4,5")
	want := "-4,5"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestDotToComma(t *testing.T) {
	got := DotToComma("34.99")
	want := "34,99"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestConvertDate(t *testing.T) {
	got, err := ConvertDate("2006-01-02")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	want := "02-01-2006"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
