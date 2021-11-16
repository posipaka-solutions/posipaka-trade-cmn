package order

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"time"
)

type Type int

const (
	Limit Type = iota
	Market
	OtherType
)

type Side int

const (
	Buy Side = iota
	Sell
	OtherSide
)

type Status int

const (
	Unknown Status = iota
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

	Price         float64
	BaseQuantity  float64
	QuoteQuantity float64
	Commission    float64

	TransactionTime time.Time
}
