package csvutil

import "testing"

func TestMakeNegativ(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		got, err := MakeNegative("4.50")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "-4.50"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Error Case", func(t *testing.T) {
		got, err := MakeNegative("abc")
		if err == nil {
			t.Fatalf("expected an error but got none")
		}
		if got != "" {
			t.Errorf("expected an empty string, got %q", got)
		}
	})
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
