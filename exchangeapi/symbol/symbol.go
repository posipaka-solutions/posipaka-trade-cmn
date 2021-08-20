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
