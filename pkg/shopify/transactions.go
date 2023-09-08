package shopify

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/DerbeDotDev/butler-csv/pkg/csvutil"
)

const (
	transactionDateIndex = 0
	typeIndex            = 1
	orderIndex           = 2
	cardBrandIndex       = 3
	payoutStatusIndex    = 5
	amountIndex          = 8
	feeIndex             = 9
	checkoutIndex        = 11
)

type Transaction struct {
	PayoutDate       string
	PaymentReference string
	Recipient        string
	OrderType        string
	BookingText      string
	Amount           string
	Purpose          string
	Fee              string
}

func (t *Transaction) fromCsvRecord(record []string) error {
	var err error
	t.PayoutDate, err = csvutil.ConvertDate(record[transactionDateIndex])
	if err != nil {
		return err
	}

	t.PaymentReference = record[orderIndex]

	t.Fee = record[feeIndex]

	t.Recipient = record[orderIndex] + " " + record[typeIndex]

	t.OrderType = record[cardBrandIndex]

	t.BookingText = record[payoutStatusIndex]

	t.Amount, err = csvutil.DotToComma(record[amountIndex])
	if err != nil {
		return err
	}

	t.Purpose = record[checkoutIndex]

	return nil
}

func (t *Transaction) isPaid(record []string) bool {
	return record[payoutStatusIndex] == "paid"
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
	expectedHeader := []string{
		"Transaction Date",
		"Type",
		"Order",
		"Card Brand",
		"Card Source",
		"Payout Status",
		"Payout Date",
		"Available On",
		"Amount",
		"Fee",
		"Net",
		"Checkout",
		"Payment Method Name",
		"Presentment Amount",
		"Presentment Currency",
		"Currency",
	}

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

	header := []string{
		"Buchungsdatum",
		"Zahlungsreferenz",
		"Zahlungspflichtiger/Empfänger",
		"Auftragsart",
		"Buchungstext",
		"Betrag",
		"Verwendungszweck",
	}

	if err := writer.Write(header); err != nil {
		return err
	}

	var fees float64
	for _, transaction := range transactions {
		num, err := strconv.ParseFloat(transaction.Fee, 64)
		if err != nil {
			return err
		}
		fees += num
	}

	fmt.Println(fees)

	for _, transaction := range transactions {
		row := []string{
			transaction.PayoutDate,
			transaction.PaymentReference,
			transaction.Recipient,
			transaction.OrderType,
			transaction.BookingText,
			transaction.Amount,
			transaction.Purpose,
		}

		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
