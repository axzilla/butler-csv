package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime/debug"
	"strings"
)

type Records [][]string

type Payout struct {
	Date      string
	Currency  string
	Recipient string
	Total     string
}

type Payouts []Payout

func logError(err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	log.Output(2, trace)
}

func getHeaderIndex(name string, headers []string) (int, error) {
	headerIndex := -1
	for i, header := range headers {
		if header == name {
			headerIndex = i
			break
		}
	}
	if headerIndex == -1 {
		return -1, errors.New("Header not found")
	}
	return headerIndex, nil
}

func getRecords(filename string) (Records, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func dotToComma(field string) string {
	return strings.ReplaceAll(field, ".", ",")
}

func convertDate(date string) string {
	dateArr := strings.Split(date, "-")
	year := dateArr[0]
	month := dateArr[1]
	day := dateArr[2]
	return fmt.Sprintf("%v-%v-%v", day, month, year)
}

func validateHeader(header []string) error {
	expectedHeader := []string{"Payout Date", "Status", "Charges", "Refunds", "Adjustments", "Reserved Funds", "Fees", "Retried Amount", "Total", "Currency"}
	if !reflect.DeepEqual(header, expectedHeader) {
		return fmt.Errorf("unexpected header: got %v, want %v", header, expectedHeader)
	}
	return nil
}

func printRecordLines(records Records) {
	for _, record := range records {
		fmt.Println(record)
	}
}

func createPayout(payoutRecords Records) (Payouts, error) {
	var payouts Payouts

	header := payoutRecords[0]
	err := validateHeader(header)
	if err != nil {
		return nil, err
	}

	statusIndex, err := getHeaderIndex("Status", header)
	if err != nil {
		return nil, err
	}

	dateIndex, err := getHeaderIndex("Payout Date", header)
	if err != nil {
		return nil, err
	}

	totalIndex, err := getHeaderIndex("Total", header)
	if err != nil {
		return nil, err
	}

	currencyIndex, err := getHeaderIndex("Currency", header)
	if err != nil {
		return nil, err
	}

	for _, payoutRecord := range payoutRecords {
		if payoutRecord[statusIndex] == "paid" {
			date := convertDate(payoutRecord[dateIndex])
			recipient := "Shopify Auszahlung"
			total := dotToComma(payoutRecord[totalIndex])
			currency := payoutRecord[currencyIndex]

			payout := Payout{
				Date:      date,
				Recipient: recipient,
				Total:     total,
				Currency:  currency,
			}

			payouts = append(payouts, payout)
		}
	}

	return payouts, nil
}

func main() {
	payoutRecords, err := getRecords("payouts.csv")
	if err != nil {
		logError(err)
		return
	}
	printRecordLines(payoutRecords)

	payouts, err := createPayout(payoutRecords)
	if err != nil {
		logError(err)
		return
	}

	for _, payout := range payouts {
		fmt.Println(payout)
	}
}
