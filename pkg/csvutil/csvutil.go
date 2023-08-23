package csvutil

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func MakeNegative(s string) (string, error) {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return "", err
	}
	negativeNum := -num
	negativeString := fmt.Sprintf("%.2f", negativeNum)
	return negativeString, nil
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
