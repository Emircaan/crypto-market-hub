# Crypto Exchange Dashboard

A real-time cryptocurrency dashboard tracking prices from Binance and Kraken, built with a modern microservices architecture.

## 🏗 Architecture

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│                 │     │                 │     │                 │
│    Frontend     │────▶│  Crypto Service │────▶│   CCXT Server   │
│   (Next.js)     │HTTP │   (GoFiber)     │gRPC │   (Python)      │
│                 │     │                 │     │                 │
└─────────────────┘     └────────┬────────┘     └────────┬────────┘
                                 │                       │
                    ┌────────────┼───────────────────────┘
                    │            │
                    ▼            ▼
           ┌─────────────┐  ┌─────────────┐     ┌─────────────────┐
           │  PostgreSQL │  │   Vault     │     │ Crypto Exchanges│
           │  (Storage)  │  │  (Secrets)  │     │ Binance, Kraken │
           └─────────────┘  └─────────────┘     └─────────────────┘
```

## 🚀 Tech Stack

- **Frontend:** Next.js, Tailwind CSS
- **Backend:** Go (Fiber), Python (CCXT)
- **Communication:** gRPC
- **Database:** PostgreSQL
- **Security:** HashiCorp Vault / Kubernetes Secrets

## ⚡️ Quick Start

The easiest way to run the project is using Docker Compose.

1. **Start all services:**
   ```bash
   docker-compose up -d --build
   ```

2. **Access the services:**
   - **Dashboard:** [http://localhost:3001](http://localhost:3001)
   - **API:** [http://localhost:3000](http://localhost:3000)

## � Features

- **Real-time Data:** Fetches market data every 20 seconds.
- **Multi-Exchange:** Supports Binance and Kraken.
- **Data Persistence:** Stores historical price data in PostgreSQL.
- **Secure:** Uses HashiCorp Vault for managing API keys.
