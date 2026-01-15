package grpc

import (
	"context"
	"errors"
	"log"

	"github.com/Emircaan/crypto-service/internal/domain"
	pb "github.com/Emircaan/crypto-service/internal/generated/proto/cryptomarket"
	"github.com/Emircaan/crypto-service/internal/vault"
)

type TickerFetcher struct {
	client      pb.CryptoMarketServiceClient
	vaultClient *vault.Client
}

func NewTickerFetcher(client pb.CryptoMarketServiceClient, vaultClient *vault.Client) *TickerFetcher {
	return &TickerFetcher{
		client:      client,
		vaultClient: vaultClient,
	}
}

func (f *TickerFetcher) FetchTickers(ctx context.Context, exchange string, symbols []string) ([]domain.Ticker, error) {
	var apiKey, apiSecret, passphrase string

	if f.vaultClient != nil {
		k, s, p, err := f.vaultClient.GetExchangeCredentials(ctx, exchange)
		if err != nil {
			log.Printf("Warning: Failed to fetch credentials for %s: %v. Proceeding without credentials.", exchange, err)
		} else {
			apiKey = k
			apiSecret = s
			passphrase = p
		}
	}

	resp, err := f.client.FetchTickers(ctx, &pb.FetchTickersRequest{
		Exchange: exchange,
		Symbols:  symbols,
		Credentials: &pb.Credentials{
			ApiKey:     apiKey,
			ApiSecret:  apiSecret,
			Passphrase: passphrase,
		},
	})
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, errors.New(resp.ErrorMessage)
	}

	tickers := make([]domain.Ticker, len(resp.Tickers))
	for i, t := range resp.Tickers {
		tickers[i] = domain.Ticker{
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

	return tickers, nil
}

func (f *TickerFetcher) GetSupportedExchanges(ctx context.Context) ([]string, error) {
	resp, err := f.client.GetSupportedExchanges(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}
	return resp.Exchanges, nil

}

func (f *TickerFetcher) HealthCheck(ctx context.Context) (bool, error) {
	resp, err := f.client.HealthCheck(ctx, &pb.Empty{})
	if err != nil {
		return false, err
	}
	return resp.Healthy, nil
}
