package shopify

import (
	"github.com/DerbeDotDev/butler-csv/pkg/csvutil"
)

const (
	transactionDateIndex = 0
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

	// zahlungsreferenz
	// empfaenger
	// auftragsart
	// buchungstext
	// betrag
	// verwendungszweck
	return nil
}

func (t *Transaction) isPaid(record []string) bool {
	return record[6] != "paid"
}
