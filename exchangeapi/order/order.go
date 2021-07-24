package order

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
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
	Symbol   exchangeapi.AssetsSymbol
	Side     Side
	Type     Type
	Quantity float64
	Price    float64
}

type Info struct {
	Id     string
	Symbol string
	Status Status
	Type   Type

	Price         float64
	BaseQuantity  float64
	QuoteQuantity float64
	Commission    float64

	TransactionTime time.Time
}
