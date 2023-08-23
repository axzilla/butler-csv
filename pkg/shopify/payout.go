package shopify

import (
	"encoding/csv"
	"github.com/DerbeDotDev/butler-csv/pkg/csvutil"
	"os"
)

type Payout struct {
	Date      string
	Currency  string
	Recipient string
	Total     string
}

func (p *Payout) fromCsvRecord(record []string) error {
	var err error
	p.Date, err = csvutil.ConvertDate(record[0])
	if err != nil {
		return err
	}

	p.Recipient = "Shopify Auszahlung"

	negativeTotal, err := csvutil.MakeNegative(record[0])
	if err != nil {
		return err
	}

	p.Total = csvutil.DotToComma(negativeTotal)

	p.Currency = record[9]

	return nil
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

func WriteCsv(payouts []Payout, csvPath string) error {
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
