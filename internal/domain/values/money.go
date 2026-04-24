package values

type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
	// add more in the future. maybe.
)

type Amount int64

type Money struct {
	Amount   Amount
	Currency Currency
}
