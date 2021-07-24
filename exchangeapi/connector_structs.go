package exchangeapi

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
)

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
	SetOrder(parameters order.Parameters) (float64, error)

	GetCurrentPrice(symbol AssetsSymbol) (float64, error)
	GetSymbolLimits(symbol AssetsSymbol) (SymbolLimits, error)
	GetOrdersList() ([]order.Info, error)

	GetServerTime() (uint64, error)
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
