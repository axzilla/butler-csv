package shopify

import (
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
