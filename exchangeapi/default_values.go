package exchangeapi

const (
	BinanceEx = "binance"
	KucoinEx  = "kucoin"
)

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

type OrderStatus int

const (
	Unknown OrderStatus = iota
	Active
	Canceled
	Filled
)
