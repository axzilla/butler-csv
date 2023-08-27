package shopify

import "fmt"

type Transaction struct {
	PayoutDate       string
	PaymentReference string
	Recipient        string
	OrderType        string
	BookingText      string
	Amount           string
	Purpose          string
}

func (t *Transaction) fromCsvRecord(record []string) {
	fmt.Println("hello world")
}

func (t *Transaction) isPaid(record []string) bool {
	return record[6] != "paid"
}
