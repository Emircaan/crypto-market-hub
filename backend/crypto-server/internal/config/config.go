package config

import (
	"fmt"
	"os"
)

type Config struct {
	App   AppConfig
	DB    DBConfig
	Grpc  GrpcConfig
	Vault VaultConfig
}

type AppConfig struct {
	Port string
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	SSLMode  string
}

type GrpcConfig struct {
	Address string
}

type VaultConfig struct {
	Address string
	Token   string
}

func Load() (*Config, error) {
	return &Config{
		App: AppConfig{
			Port: getEnv("PORT", "3000"),
		},
		DB: DBConfig{
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Name:     getEnv("DB_NAME", "crypto_exchange"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Grpc: GrpcConfig{
			Address: getEnv("GRPC_SERVER_ADDR", "localhost:50051"),
		},
		Vault: VaultConfig{
			Address: getEnv("VAULT_ADDR", ""),
			Token:   getEnv("VAULT_TOKEN", ""),
		},
	}, nil
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode)
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
