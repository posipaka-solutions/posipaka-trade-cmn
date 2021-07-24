package exchangeapi

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
)

type ApiKey struct {
	Key        string
	Secret     string
	Passphrase string
}

type ApiConnector interface {
	SetOrder(parameters order.Parameters) (float64, error)

	GetCurrentPrice(assets symbol.Assets) (float64, error)
	GetSymbolLimits(assets symbol.Assets) (symbol.Limits, error)
	GetOrdersList() ([]order.Info, error)

	GetServerTime() (uint64, error)
}
