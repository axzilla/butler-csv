package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

type Payout struct {
	Date      string
	Currency  string
	Recipient string
	Total     string
}

func (p *Payout) fromCsvRecord(record []string) error {
	var err error
	p.Date, err = convertDate(record[0])
	if err != nil {
		return err
	}
	p.Recipient = "Shopify Auszahlung"
	p.Total = makeNegative(dotToComma(record[8]))
	p.Currency = record[9]
	return nil
}

func makeNegative(s string) string {
	return "-" + s
}

func dotToComma(s string) string {
	return strings.ReplaceAll(s, ".", ",")
}

func convertDate(s string) (string, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return "", err
	}
	return t.Format("02-01-2006"), nil
}

func readPayouts() ([]Payout, error) {
	file, err := os.Open("payouts.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var payouts []Payout
	for _, record := range records {
		if record[1] == "paid" {
			var p Payout
			err := p.fromCsvRecord(record)
			if err != nil {
				return nil, err
			}
			payouts = append(payouts, p)
		}
	}

	return payouts, nil
}

func writeCsv(payouts []Payout) error {
	file, err := os.Create("new_payouts.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Buchungsdatum", "Zahlungspflichtiger/Empfänger", "Betrag", "Währung"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, payout := range payouts {
		row := []string{payout.Date, payout.Recipient, payout.Total, payout.Currency}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	payouts, err := readPayouts()
	if err != nil {
		fmt.Println(err)
	}

	err = writeCsv(payouts)
	if err != nil {
		fmt.Println(err)
	}
}
