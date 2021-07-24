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
	Id       string
	Status   OrderStatus
	Type     OrderType
	Price    float64
	Quantity float64
}

type Candlesticks struct {
	OpenTime                 int64  `json:"openTime"`
	Open                     string `json:"open"`
	High                     string `json:"high"`
	Low                      string `json:"low"`
	Close                    string `json:"close"`
	Volume                   string `json:"volume"`
	CloseTime                int64  `json:"closeTime"`
	QuoteAssetVolume         string `json:"quoteAssetVolume"`
	NumberOfTrade            int64  `json:"numberOfTrade"`
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
	Ignore                   string `json:"ignore"`
}
