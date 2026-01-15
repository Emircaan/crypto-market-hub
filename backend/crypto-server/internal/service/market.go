package service

import (
	"context"
	"log"
	"time"

	"github.com/Emircaan/crypto-service/internal/domain"
)

type MarketProvider interface {
	FetchTickers(ctx context.Context, exchange string, symbols []string) ([]domain.Ticker, error)
	GetSupportedExchanges(ctx context.Context) ([]string, error)
}

type MarketService struct {
	provider MarketProvider
	repo     domain.TickerRepository
}

func NewMarketService(provider MarketProvider, repo domain.TickerRepository) *MarketService {
	return &MarketService{
		provider: provider,
		repo:     repo,
	}
}

func (s *MarketService) GetTickers(ctx context.Context, exchange string) ([]domain.Ticker, error) {
	return s.repo.ListByExchange(ctx, exchange)
}

func (s *MarketService) GetSupportedExchanges(ctx context.Context) ([]string, error) {
	return s.provider.GetSupportedExchanges(ctx)
}

func (s *MarketService) StartTickerUpdater(ctx context.Context, exchanges []string, symbols []string, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		s.updateAllTickers(ctx, exchanges, symbols)

		for {
			select {
			case <-ctx.Done():
				log.Println("Ticker updater stopped")
				return
			case <-ticker.C:
				s.updateAllTickers(ctx, exchanges, symbols)
			}
		}
	}()
}

func (s *MarketService) updateAllTickers(ctx context.Context, exchanges []string, symbols []string) {
	for _, exchange := range exchanges {
		tickers, err := s.provider.FetchTickers(ctx, exchange, symbols)
		if err != nil {
			log.Printf("Error fetching tickers for %s: %v", exchange, err)
			continue
		}

		for _, t := range tickers {
			if err := s.repo.Save(ctx, t); err != nil {
				log.Printf("Error saving ticker %s:%s: %v", t.Exchange, t.Symbol, err)
			}
		}

		log.Printf("Updated %d tickers for %s", len(tickers), exchange)
	}
}
