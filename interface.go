package preev

import "context"

// PairService is the Preev pair related methods
type PairService interface {
	GetPair(ctx context.Context, pairID string) (pair *Pair, err error)
	GetPairs(ctx context.Context) (pairList *PairList, err error)
}

// TickerService is the Preev ticker related methods
type TickerService interface {
	GetTicker(ctx context.Context, pairID string) (ticker *Ticker, err error)
	GetTickerHistory(ctx context.Context, pairID string, start, end, interval int64) (tickers []*Ticker, err error)
	GetTickers(ctx context.Context) (tickerList *TickerList, err error)
}

// ClientInterface is the Preev client interface
type ClientInterface interface {
	PairService
	TickerService
	LastRequest() *LastRequest
	UserAgent() string
}
