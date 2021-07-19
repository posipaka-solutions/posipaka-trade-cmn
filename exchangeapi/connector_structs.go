package exchangeapi

import "net/http"

type ApiKey struct {
	Key        string
	Secret     string
	Passphrase string
}

type AssetsSymbol struct {
	Base  string
	Quote string
}

type ApiConnector interface {
	SetOrder(parameters OrderParameters) (float64, error)

	GetCurrentPrice(symbol AssetsSymbol) (float64, error)
	GetSymbolLimits(symbol AssetsSymbol) (SymbolLimits, error)
	GetOrdersList() ([]OrderInfo, error)

	GetServerTime() (uint64, error)
}

type OrderParameters struct {
	Symbol   AssetsSymbol
	Side     OrderSide
	Type     OrderType
	Quantity float64
	Price    float64
}

type SymbolLimits struct {
	Symbol AssetsSymbol

	BaseMinSize  float64
	BaseMaxSize  float64
	QuoteMinSize float64
	QuoteMaxSize float64

	BaseIncrement  float64
	QuoteIncrement float64
	PriceIncrement float64
}

type OrderInfo struct {
	Id        string
	Symbol    string
	OrderType string
	Price     float64
	Quantity  float64
}
