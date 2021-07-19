package exchangeapi

type OrderType int

const (
	Limit OrderType = iota
	Market
)

type OrderSide int

const (
	Buy OrderSide = iota
	Sell
)
