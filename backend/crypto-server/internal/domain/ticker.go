package domain

import "context"

type Ticker struct {
	Exchange      string  `json:"exchange"`
	Symbol        string  `json:"symbol"`
	Price         float64 `json:"price"`
	Volume        float64 `json:"volume"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	ChangePercent float64 `json:"change_percent"`
	Timestamp     int64   `json:"timestamp"`
}

type TickerRepository interface {
	Save(ctx context.Context, ticker Ticker) error
	GetLatest(ctx context.Context, exchange, symbol string) (Ticker, error)
	ListByExchange(ctx context.Context, exchange string) ([]Ticker, error)
}
