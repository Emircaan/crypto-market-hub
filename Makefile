.PHONY: proto-py proto-go proto-all sqlc-gen migrate docker-up docker-db help

# Generate Python protobuf files
proto-py:
	@cd backend/ccxt-server && uv run python -m grpc_tools.protoc -I../proto --python_out=src/generated --grpc_python_out=src/generated ../proto/crypto-market.proto
	@echo "✅ Python proto files generated"

# Generate Go protobuf files
proto-go:
	@cd backend/crypto-server && protoc -I../proto --go_out=. --go-grpc_out=. ../proto/crypto-market.proto
	@echo "✅ Go proto files generated"

# Generate all protobuf files
proto-all: proto-py proto-go

# Generate SQLC code
sqlc-gen:
	@cd backend/crypto-server && sqlc generate
	@echo "✅ SQLC code generated"

# Run database migrations
migrate:
	@PGPASSWORD=postgres psql -h localhost -U postgres -d crypto_exchange -f backend/crypto-server/internal/database/migrations/000001_init_schema.up.sql
	@echo "✅ Migrations applied"

# Start only PostgreSQL
docker-db:
	docker compose up -d postgres
	@echo "✅ PostgreSQL started"

# Start all services with Docker Compose
docker-up:
	docker compose up -d

# Show available commands
help:
	@echo "Available commands:"
	@echo "  make proto-py    - Generate Python protobuf files"
	@echo "  make proto-go    - Generate Go protobuf files"
	@echo "  make proto-all   - Generate all protobuf files"
	@echo "  make sqlc-gen    - Generate SQLC code"
	@echo "  make migrate     - Run database migrations"
	@echo "  make docker-db   - Start only PostgreSQL"
	@echo "  make docker-up   - Start all services with Docker Compose"
