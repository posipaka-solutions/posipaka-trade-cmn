package symbol

type Assets struct {
	Base  string
	Quote string
}

type Limits struct {
	Assets Assets

	BaseMinSize  float64
	BaseMaxSize  float64
	QuoteMinSize float64
	QuoteMaxSize float64

	BaseIncrement  float64
	QuoteIncrement float64
	PriceIncrement float64
}
