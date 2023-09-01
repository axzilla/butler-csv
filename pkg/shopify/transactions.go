package shopify

import (
	"encoding/csv"
	"os"

	"github.com/DerbeDotDev/butler-csv/pkg/csvutil"
)

const (
	transactionDateIndex = 0
	typeIndex            = 1
	orderIndex           = 3
	cardBrandIndex       = 4
	payoutStatusIndex    = 6
	amountIndex          = 9
	purposeIndex         = 12
)

type Transaction struct {
	PayoutDate       string
	PaymentReference string
	Recipient        string
	OrderType        string
	BookingText      string
	Amount           string
	Purpose          string
}

func (t *Transaction) fromCsvRecord(record []string) error {
	var err error
	t.PayoutDate, err = csvutil.ConvertDate(record[transactionDateIndex])
	if err != nil {
		return err
	}

	t.PaymentReference = record[orderIndex]

	t.Recipient = record[orderIndex] + " " + record[typeIndex]

	t.OrderType = record[cardBrandIndex]

	t.BookingText = record[payoutStatusIndex]

	t.Amount, err = csvutil.DotToComma(record[amountIndex])
	if err != nil {
		return err
	}

	t.Purpose = record[purposeIndex]

	return nil
}

func (t *Transaction) isPaid(record []string) bool {
	return record[6] != "paid"
}

func ReadTransactions(csvPath string) ([]Transaction, error) {
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

	actualHeader := records[0]
	expectedHeader := []string{"Transaction Date", "Type", "Order", "Card Brand", "Payout Status", "Payout Date", "Available On", "Amount", "Fee", "Net", "Checkout", "Payment Method Name", "Presentment Amount", "Presentment Currency", "Currency"}
	err = csvutil.ValidateCsvHeader(actualHeader, expectedHeader)
	if err != nil {
		return nil, err
	}

	var transactions []Transaction
	for _, record := range records {
		var t Transaction
		if t.isPaid(record) {
			err := t.fromCsvRecord(record)
			if err != nil {
				return nil, err
			}
			transactions = append(transactions, t)
		}
	}

	return transactions, nil
}

func WriteTransactions(transactions []Transaction, csvPath string) error {
	file, err := os.Create("new_transactions.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// header := []string{"Buchungsdatum", "Zahlungspflichtiger/Empfänger", "Betrag", "Währung"}
	// if err := writer.Write(header); err != nil {
	// 	return err
	// }
	//
	// for _, payout := range payouts {
	// 	row := []string{payout.Date, payout.Recipient, payout.Amount, payout.Currency}
	// 	if err := writer.Write(row); err != nil {
	// 		return err
	// 	}
	// }
	//
	return nil
}
