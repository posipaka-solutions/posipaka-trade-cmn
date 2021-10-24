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
	New Status = iota
	Filled
	Canceled
	Rejected
	Expired
	OtherStatus
)

type Parameters struct {
	Assets   symbol.Assets
	Side     Side
	Type     Type
	Quantity float64
	Price    float64
}

type Info struct {
	Id     string
	Assets symbol.Assets
	Status Status
	Type   Type

	Price         float64
	BaseQuantity  float64
	QuoteQuantity float64
	Commission    float64

	TransactionTime time.Time
}

type OrderInfo struct {
	Id       int
	Price    float64
	Quantity float64
}
