package exchangeapi

type ApiKey struct {
	Key        string
	Secret     string
	Passphrase string
}

type ApiConnector interface {
	SetOrder(parameters OrderParameters) (float64, error)

	GetCurrentPrice(symbol string) (float64, error)
	GetSymbolLimits(symbol string) (SymbolLimits, error)

	GetServerTime() (uint64, error)
}

type OrderParameters struct {
}

type SymbolLimits struct {
}
