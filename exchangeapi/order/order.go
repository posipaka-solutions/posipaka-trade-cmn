package order

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"time"
)

type Type int

const (
	UnknownType Type = iota
	Limit
	Market
)

type Side int

const (
	UnknownSide Side = iota
	Buy
	Sell
)

type Status int

const (
	UnknownStatus Status = iota
	Open
	Filled
	Canceled
	Expired
)

// Parameters contains all available and needed parameter for order setting at exchange
type Parameters struct {
	Assets   symbol.Assets
	Side     Side
	Type     Type
	Quantity float64
	Price    float64
}

// Info has a description about order at exchange
type Info struct {
	Id     string
	Assets symbol.Assets
	Side   Side
	Status Status
	Type   Type

	// Price match price set on Limit order
	Price float64
	// FilledPrice match resulting price when Market or Limit order is Filled
	FilledPrice   float64
	BaseQuantity  float64
	QuoteQuantity float64
	Commission    float64

	TransactionTime time.Time
}
