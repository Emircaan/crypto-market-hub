package main

import (
	"context"
	"log"
	"time"

	"github.com/Emircaan/crypto-service/internal/config"
	"github.com/Emircaan/crypto-service/internal/database"
	"github.com/Emircaan/crypto-service/internal/grpc"
	"github.com/Emircaan/crypto-service/internal/handler/http"
	"github.com/Emircaan/crypto-service/internal/repository"
	"github.com/Emircaan/crypto-service/internal/server"
	"github.com/Emircaan/crypto-service/internal/service"
	"github.com/Emircaan/crypto-service/internal/vault"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := database.RunMigrations(cfg.DB.DSN()); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	vaultClient, err := vault.NewClient(vault.Config{
		Address: cfg.Vault.Address,
		Token:   cfg.Vault.Token,
	})
	if err != nil {
		log.Printf("Warning: Failed to initialize vault client: %v. Continuing without vault...", err)
	}

	grpcClient, err := grpc.NewClient(cfg.Grpc.Address)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer grpcClient.Close()

	dbPool, err := pgxpool.New(context.Background(), cfg.DB.DSN())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbPool.Close()

	tickerRepo := repository.NewTickerRepository(dbPool)

	tickerFetcher := grpc.NewTickerFetcher(grpcClient.API(), vaultClient)

	marketService := service.NewMarketService(tickerFetcher, tickerRepo)

	exchanges := []string{"binance", "kraken"}
	symbols := []string{"BTC/USDT", "ETH/USDT", "SOL/USDT", "XRP/USDT", "ADA/USDT", "DOGE/USDT", "AVAX/USDT", "DOT/USDT", "MATIC/USDT", "LINK/USDT"}
	marketService.StartTickerUpdater(context.Background(), exchanges, symbols, 20*time.Second)

	marketHandler := http.NewMarketHandler(marketService)

	srv := server.NewServer(marketHandler)

	log.Printf("Server starting on port %s", cfg.App.Port)
	if err := srv.Listen(":" + cfg.App.Port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
