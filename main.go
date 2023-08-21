package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	transactions, err := os.Open("transactions.csv")
	if err != nil {
		log.Fatalf("Error on opening file: %v", err)
	}

	reader := csv.NewReader(transactions)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error on reading file: %v", err)
	}

	for _, record := range records {
		fmt.Println(record)
	}
}
