package exchangeapi

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"time"
)

type ApiKey struct {
	Key        string
	Secret     string
	Passphrase string
}

type ApiConnector interface {
	SetOrder(parameters order.Parameters) (float64, error)

	GetCurrentPrice(assets symbol.Assets) (float64, error)
	GetAssetBalance(asset string) (float64, error)

	GetSymbolsLimits() ([]symbol.Limits, error)
	GetOrdersList(assets symbol.Assets) ([]order.Info, error)
	StoreSymbolsLimits([]symbol.Limits)

	GetSymbolsList() []symbol.Assets
	GetServerTime() (time.Time, error)
}

type Candlestick struct {
	OpenTime  time.Time
	CloseTime time.Time

	OpenPrice  float64
	ClosePrice float64
	MaxPrice   float64
	MinPrice   float64

	Volume       float64
	TradesNumber int
}
