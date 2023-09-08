package shopify

import (
	"encoding/csv"
	"os"

	"github.com/DerbeDotDev/butler-csv/pkg/csvutil"
)

const (
	payoutDateIndex = 0
	statusIndex     = 1
	totalIndex      = 8
	currencyIndex   = 9
)

type Payout struct {
	Date      string
	Currency  string
	Recipient string
	Amount    string
}

func (p *Payout) fromCsvRecord(record []string) error {
	var err error
	p.Date, err = csvutil.ConvertDate(record[payoutDateIndex])
	if err != nil {
		return err
	}

	p.Recipient = "Shopify Auszahlung"

	negativeTotal, err := csvutil.MakeNegative(record[totalIndex])
	if err != nil {
		return err
	}

	p.Amount, err = csvutil.DotToComma(negativeTotal)
	if err != nil {
		return err
	}

	p.Currency = record[currencyIndex]

	return nil
}

func (p *Payout) isPaid(record []string) bool {
	return record[statusIndex] == "paid"
}

func ReadPayouts(csvPath string) ([]Payout, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	actualHeaders := records[0]
	expectedHeaders := []string{
		"Payout Date",
		"Status",
		"Charges",
		"Refunds",
		"Adjustments",
		"Reserved Funds",
		"Fees",
		"Retried Amount",
		"Total",
		"Currency",
	}

	err = csvutil.ValidateCsvHeader(actualHeaders, expectedHeaders)
	if err != nil {
		return nil, err
	}

	var payouts []Payout
	for _, record := range records {
		var p Payout
		if p.isPaid(record) {
			err := p.fromCsvRecord(record)
			if err != nil {
				return nil, err
			}
			payouts = append(payouts, p)
		}
	}

	return payouts, nil
}

func WriteCsv(payouts []Payout, csvPath string) error {
	file, err := os.Create(csvPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{
		"Buchungsdatum",
		"Zahlungspflichtiger/Empfänger",
		"Betrag",
		"Währung",
	}

	if err := writer.Write(header); err != nil {
		return err
	}

	for _, payout := range payouts {
		row := []string{
			payout.Date,
			payout.Recipient,
			payout.Amount,
			payout.Currency,
		}

		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
