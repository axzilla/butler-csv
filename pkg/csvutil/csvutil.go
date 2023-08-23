package csvutil

import (
	"strings"
	"time"
)

func MakeNegative(s string) string {
	return "-" + s
}

func DotToComma(s string) string {
	return strings.ReplaceAll(s, ".", ",")
}

func ConvertDate(s string) (string, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return "", err
	}
	return t.Format("02-01-2006"), nil
}
