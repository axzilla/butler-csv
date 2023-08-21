package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func printRecordLines(records [][]string) {
	for _, record := range records {
		fmt.Println(record)
	}
}

func getRecords(filename string) [][]string {
	payouts, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Error on opening file: %v", err)
	}

	reader := csv.NewReader(payouts)
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatalf("Error on reading file: %v", err)
	}

	return records
}

func getHeaderIndex(name string, header []string) int {
	index := -1

	for i, column := range header {
		if column == name {
			index = i
			break
		}
	}

	if index == -1 {
		log.Fatal("Could not find header index!")
	}

	return index
}

func removePendingPayouts(payouts [][]string) [][]string {
	var newPayouts [][]string
	header := payouts[0]
	newPayouts = append(newPayouts, header)
	status := getHeaderIndex("Status", header)

	for _, payout := range payouts {
		if payout[status] == "paid" {
			newPayouts = append(newPayouts, payout)
		}
	}

	return newPayouts
}

func main() {
	payouts := getRecords("payouts.csv")
	printRecordLines(payouts)

	payouts = removePendingPayouts(payouts)
	printRecordLines(payouts)
}
