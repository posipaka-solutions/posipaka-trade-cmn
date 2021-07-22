package exchangeapi

import "fmt"

type ExchangeError struct {
	Type    int
	Code    int
	Message string
}

const (
	HttpErr = iota
	BinanceErr
	KucoinErr
	ConnectorErr
)

var EmptyResponseErr = ExchangeError{
	Type:    ConnectorErr,
	Code:    1001,
	Message: "Exchange response is empty",
}

func (err *ExchangeError) Error() string {
	return fmt.Sprint("[ExchangeError] -> ", err.Message)
}
