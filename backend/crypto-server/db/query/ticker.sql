-- name: CreateTicker :one
INSERT INTO tickers (
  exchange, symbol, price, volume, high, low, change_percent, timestamp
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetLatestTicker :one
SELECT * FROM tickers
WHERE exchange = $1 AND symbol = $2
ORDER BY timestamp DESC
LIMIT 1;

-- name: ListTickersByExchange :many
SELECT DISTINCT ON (symbol) * FROM tickers
WHERE exchange = $1
ORDER BY symbol, timestamp DESC;
