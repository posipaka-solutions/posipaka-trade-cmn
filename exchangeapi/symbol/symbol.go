package symbol

type Assets struct {
	Base  string
	Quote string
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
}
