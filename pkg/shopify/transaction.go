package shopify

type Transaction struct {
	Date        string
	Reference   string
	Recipient   string
	OrderType   string
	BookingText string
	Amount      string
	Purpose     string
}

func (t *Transaction) fromCsvRecord() error {
	return nil
}
