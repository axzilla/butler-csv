package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
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

func (p *Payouts) convertNumbers() {
	var convertedPayouts Payouts
	for _, payout := range *p {
		payout.Fees = strings.ReplaceAll(payout.Fees, ".", ",")
		payout.Charges = strings.ReplaceAll(payout.Charges, ".", ",")
		payout.Refunds = strings.ReplaceAll(payout.Refunds, ".", ",")
		payout.Adjustments = strings.ReplaceAll(payout.Adjustments, ".", ",")
		payout.ReservedFunds = strings.ReplaceAll(payout.ReservedFunds, ".", ",")
		payout.Fees = strings.ReplaceAll(payout.Fees, ".", ",")
		payout.RetriedAmount = strings.ReplaceAll(payout.RetriedAmount, ".", ",")
		payout.Total = strings.ReplaceAll(payout.Total, ".", ",")
		convertedPayouts = append(convertedPayouts, payout)
	}
	*p = convertedPayouts
}

func validateHeader(header []string) error {
	expectedHeader := []string{"Payout Date", "Status", "Charges", "Refunds", "Adjustments", "Reserved Funds", "Fees", "Retried Amount", "Total", "Currency"}
	if !reflect.DeepEqual(header, expectedHeader) {
		return fmt.Errorf("unexpected header: got %v, want %v", header, expectedHeader)
	}
	return nil
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

	header := records[0]
	err = validateHeader(header)
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
		log.Fatalf("An error occurred: %v", err)
	}

	printRecordLines(payouts)
	payouts.filterPaid()
	payouts.convertNumbers()
	printRecordLines(payouts)
}
