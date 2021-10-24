package exchangeapi

import "fmt"

type ExchangeError struct {
	Type        int
	Code        int
	Message     string
	KeysDetails map[string]string
}

const (
	HttpErr = iota
	BinanceErr
	KucoinErr
	GateErr
)

func (err *ExchangeError) Error() string {
	return fmt.Sprint("[ExchangeError] -> ", err.Message)
}
