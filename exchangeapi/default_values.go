package exchangeapi

const (
	BinanceEx = "binance"
	KucoinEx  = "kucoin"
)

type OrderType int

const (
	Limit OrderType = iota
	Market
	OtherType
)

type OrderSide int

const (
	Buy OrderSide = iota
	Sell
	OtherSide
)

type OrderStatus int

const (
	NewOrder OrderStatus = iota
	Filled
	Canceled
	Rejected
	Expired
	OtherStatus
)
