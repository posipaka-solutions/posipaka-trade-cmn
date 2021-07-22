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

const emptyResponseErrorCode = 1001
const emptyResponseErrorMsg = "Exchange response is empty"

func (err *ExchangeError) Error() string {
	return fmt.Sprint("[ExchangeError] -> ", err.Message)
}
