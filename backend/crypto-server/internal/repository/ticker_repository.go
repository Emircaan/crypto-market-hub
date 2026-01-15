package repository

import (
	"context"

	"github.com/Emircaan/crypto-service/internal/domain"
	"github.com/Emircaan/crypto-service/internal/generated/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TickerRepository struct {
	q *db.Queries
}

func NewTickerRepository(conn *pgxpool.Pool) *TickerRepository {
	return &TickerRepository{
		q: db.New(conn),
	}
}

func (r *TickerRepository) Save(ctx context.Context, t domain.Ticker) error {
	_, err := r.q.CreateTicker(ctx, db.CreateTickerParams{
		Exchange:      t.Exchange,
		Symbol:        t.Symbol,
		Price:         t.Price,
		Volume:        t.Volume,
		High:          t.High,
		Low:           t.Low,
		ChangePercent: t.ChangePercent,
		Timestamp:     t.Timestamp,
	})
	return err
}

func (r *TickerRepository) GetLatest(ctx context.Context, exchange, symbol string) (domain.Ticker, error) {
	t, err := r.q.GetLatestTicker(ctx, db.GetLatestTickerParams{
		Exchange: exchange,
		Symbol:   symbol,
	})
	if err != nil {
		return domain.Ticker{}, err
	}

	return domain.Ticker{
		Exchange:      t.Exchange,
		Symbol:        t.Symbol,
		Price:         t.Price,
		Volume:        t.Volume,
		High:          t.High,
		Low:           t.Low,
		ChangePercent: t.ChangePercent,
		Timestamp:     t.Timestamp,
	}, nil
}

func (r *TickerRepository) ListByExchange(ctx context.Context, exchange string) ([]domain.Ticker, error) {
	tickers, err := r.q.ListTickersByExchange(ctx, exchange)
	if err != nil {
		return nil, err
	}

	result := make([]domain.Ticker, len(tickers))
	for i, t := range tickers {
		result[i] = domain.Ticker{
			Exchange:      t.Exchange,
			Symbol:        t.Symbol,
			Price:         t.Price,
			Volume:        t.Volume,
			High:          t.High,
			Low:           t.Low,
			ChangePercent: t.ChangePercent,
			Timestamp:     t.Timestamp,
		}
	}
	return result, nil
}
