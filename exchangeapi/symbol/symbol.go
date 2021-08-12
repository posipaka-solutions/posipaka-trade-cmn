package symbol

type Assets struct {
	Base  string
	Quote string
}

func (assets Assets) IsEqual(otherAssets Assets) bool {
	return assets.Base == otherAssets.Base &&
		assets.Quote == otherAssets.Quote
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
