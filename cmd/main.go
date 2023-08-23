package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/DerbeDotDev/butler-csv/pkg/shopify"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	payoutCsvPath := filepath.Join(wd, "payouts.csv")
	newPayoutCsvPath := filepath.Join(wd, "new_payouts.csv")

	payouts, err := shopify.ReadPayouts(payoutCsvPath)
	if err != nil {
		fmt.Println("Error reading payouts:", err)
		return
	}

	err = shopify.WriteCsv(payouts, newPayoutCsvPath)
	if err != nil {
		fmt.Println("Error writing new payouts:", err)
		return
	}
}
