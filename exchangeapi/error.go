package exchangeapi

import "fmt"

type ExchangeError struct {
	Type    int
	Code    int
	Message string
}

//const (
//	HttpErr = iota
//	BinanceErr
//	KucoinErr
//	ConnectorErr
//)
//
//const EmptyResponseErrorCode = 1001
//const EmptyResponseErrorMsg = "Exchange response is empty"

func (err *ExchangeError) Error() string {
	return fmt.Sprint("[ExchangeError] -> ", err.Message)
}
