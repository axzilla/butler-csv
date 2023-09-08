package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/DerbeDotDev/butler-csv/pkg/shopify"
)

func main() {
	var input string
	var output string

	flag.StringVar(&input, "input", "", "Path to the input CSV")
	flag.StringVar(&output, "output", "", "Path to save the new CSV")

	flag.Parse()

	if input == "" || output == "" {
		fmt.Println("Both input and output flags are required.")
		os.Exit(1)
	}

	payoutInput := filepath.Join(input, "payouts.csv")
	payoutOutput := filepath.Join(output, "new_payouts.csv")

	payouts, err := shopify.ReadPayouts(payoutInput)
	if err != nil {
		fmt.Println("Error reading payouts:", err)
		return
	}

	err = shopify.WriteCsv(payouts, payoutOutput)
	if err != nil {
		fmt.Println("Error writing new payouts:", err)
		return
	}

	transactionInput := filepath.Join(input, "transactions.csv")
	transactionOutput := filepath.Join(output, "new_transactions.csv")

	transactions, err := shopify.ReadTransactions(transactionInput)
	if err != nil {
		fmt.Println("Error reading transactions:", err)
		return
	}

	err = shopify.WriteTransactions(transactions, transactionOutput)
	if err != nil {
		fmt.Println("Error writing new transactions:", err)
		return
	}
}
