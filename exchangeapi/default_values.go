package exchangeapi

type CryptoExchange int
const (
	BinanceEx CryptoExchange = iota
	Kucoin
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
