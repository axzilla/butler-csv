package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Payout struct {
	PayoutDate    string
	Status        string
	Charges       string
	Refunds       string
	Adjustments   string
	ReservedFunds string
	Fees          string
	RetriedAmount string
	Total         string
	Currency      string
}

type Payouts []Payout

func (p *Payouts) filterPaid() {
	var filteredPayouts Payouts
	for _, payout := range *p {
		if payout.Status == "paid" {
			filteredPayouts = append(filteredPayouts, payout)
		}
	}
	*p = filteredPayouts
}

func printRecordLines(records []Payout) {
	for _, record := range records {
		fmt.Println(record)
	}
}

func getPayouts() (Payouts, error) {
	file, err := os.Open("payouts.csv")

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var payouts Payouts

	for _, record := range records {
		payout := Payout{
			PayoutDate:    record[0],
			Status:        record[1],
			Charges:       record[2],
			Refunds:       record[3],
			Adjustments:   record[4],
			ReservedFunds: record[5],
			Fees:          record[6],
			RetriedAmount: record[7],
			Total:         record[8],
			Currency:      record[9],
		}

		payouts = append(payouts, payout)
	}

	return payouts, nil

}

func main() {
	payouts, err := getPayouts()
	if err != nil {
		log.Fatalf("Could not get payouts: %v", err)
	}
	payouts.filterPaid()
	printRecordLines(payouts)
}
