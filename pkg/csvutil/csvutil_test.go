package csvutil

import "testing"

func TestValidateCsvHeader(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		actualHeaders := []string{"Payout Date", "Status", "Charges", "Refunds", "Adjustments", "Reserved Funds", "Fees", "Retried Amount", "Total", "Currency"}
		expectedHeaders := []string{"Payout Date", "Status", "Charges", "Refunds", "Adjustments", "Reserved Funds", "Fees", "Retried Amount", "Total", "Currency"}
		err := ValidateCsvHeader(actualHeaders, expectedHeaders)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
	})

	t.Run("Error Case: Missing Header", func(t *testing.T) {
		actualHeaders := []string{"Payout Date", "Status", "Charges", "Refunds", "Adjustments", "Reserved Funds", "Fees", "Retried Amount", "Total"}
		expectedHeaders := []string{"Payout Date", "Status", "Charges", "Refunds", "Adjustments", "Reserved Funds", "Fees", "Retried Amount", "Total", "Currency"}
		err := ValidateCsvHeader(actualHeaders, expectedHeaders)
		if err == nil {
			t.Errorf("Expected an error, but got none")
		}
	})

	t.Run("Error Case: Wrong Header", func(t *testing.T) {
		actualHeaders := []string{"abc", "Status", "Charges", "Refunds", "Adjustments", "Reserved Funds", "Fees", "Retried Amount", "Total", "Currency"}
		expectedHeaders := []string{"Payout Date", "Status", "Charges", "Refunds", "Adjustments", "Reserved Funds", "Fees", "Retried Amount", "Total", "Currency"}
		err := ValidateCsvHeader(actualHeaders, expectedHeaders)
		if err == nil {
			t.Errorf("Expected an error, but got none")
		}
	})
}

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
	t.Run("Happy Path", func(t *testing.T) {
		got, err := DotToComma("4.50")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "4,50"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Error Case", func(t *testing.T) {
		got, err := DotToComma("abc")
		if err == nil {
			t.Fatalf("expected an error but got none")
		}
		if got != "" {
			t.Errorf("expected an empty string, got %q", got)
		}
	})
}

func TestConvertDate(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		got, err := ConvertDate("2006-01-02")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "02.01.2006"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Error Case", func(t *testing.T) {
		got, err := ConvertDate("abc")
		if err == nil {
			t.Fatalf("expected an error but got none")
		}
		if got != "" {
			t.Errorf("expected an empty string, got %q", got)
		}
	})
}
