package csvutil

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ValidateCsvHeader(actualHeaders []string, expectedHeaders []string) error {
	if len(actualHeaders) != len(expectedHeaders) {
		return fmt.Errorf("Invalid CSV header length. Expected %d columns, but found %d", len(expectedHeaders), len(actualHeaders))
	}
	for i, actualHeader := range actualHeaders {
		if actualHeader != expectedHeaders[i] {
			return fmt.Errorf("Invalid header at index %d. Expected '%s', but found '%s'", i, expectedHeaders[i], actualHeader)
		}
	}
	return nil
}

func MakeNegative(s string) (string, error) {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f", -num), nil
}

func DotToComma(s string) (string, error) {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(s, ".", ","), nil
}

func ConvertDate(s string) (string, error) {
	var t time.Time
	var err error
	if strings.Contains(s, ":") {
		t, err = time.Parse("2006-01-02 15:04:05 -0700", s)
	} else {
		t, err = time.Parse("2006-01-02", s)
	}
	if err != nil {
		return "", err
	}
	return t.Format("02.01.2006"), nil
}
