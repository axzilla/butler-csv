package main

import (
	"fmt"
	"github.com/DerbeDotDev/butler-csv/pkg/shopify/payout"
)

func main() {
	payouts, err := payout.ReadPayouts()
	if err != nil {
		fmt.Println(err)
	}

	err = payout.WriteCsv(payouts)
	if err != nil {
		fmt.Println(err)
	}
}
