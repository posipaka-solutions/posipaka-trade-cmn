package symbol

type Assets struct {
	Base  string
	Quote string
}

func (assets Assets) IsEqual(otherAssets Assets) bool {
	return assets.Base == otherAssets.Base &&
		assets.Quote == otherAssets.Quote
}

func (assets Assets) IsEmpty() bool {
	return len(assets.Base) == 0 &&
		len(assets.Quote) == 0
}

type Limits struct {
	Assets Assets

	Base  LimitDetail
	Quote LimitDetail
	Price LimitDetail
}

type LimitDetail struct {
	MinSize   float64
	MaxSize   float64
	Increment float64
	Precision int
}

type AssetPrice struct {
	Symbol Assets
	Price  float64
}

type ArbitrageAsset struct {
	Symbol    Assets
	Price     float64
	UsdtPrice float64
}

type AllPricesList struct {
	Symbol string
	Price  float64
}

type OrderBook struct {
	Bid []BidAsk
	Ask []BidAsk
}

type BidAsk struct {
	Price, Quantity float64
}
