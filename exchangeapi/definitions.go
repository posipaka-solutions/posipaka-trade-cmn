package exchangeapi

import "net/http"

type ApiKey struct {
	Key        string
	Secret     string
	Passphrase string
}

type ExchangeManager struct {
	apiKey ApiKey
	client http.Client
}

type ApiConnector interface {
	SetOrder(parameters OrderParameters) (float64, error)

	GetCurrentPrice(symbol string) (float64, error)
	GetSymbolLimits(symbol string) (SymbolLimits, error)
	GetOrdersList() ([]OrderInfo, error)

	GetServerTime() (uint64, error)
}

type OrderParameters struct {
	Symbol    string
	Side      string
	OrderType int
	Quantity  float64
	Price     float64
}

type SymbolLimits struct {
	symbol string

	baseMinSize  float64
	baseMaxSize  float64
	quoteMinSize float64
	quoteMaxSize float64

	baseIncrement  float64
	quoteIncrement float64
	priceIncrement float64
}

type OrderInfo struct {
	Id        string
	Symbol    string
	OrderType string
	Price     float64
	Size      float64
}
