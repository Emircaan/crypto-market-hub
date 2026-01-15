package vault

import (
	"context"
	"errors"
	"fmt"
	"os"

	vault "github.com/hashicorp/vault/api"
)

type Config struct {
	Address string
	Token   string
}

type Client struct {
	client *vault.Client
}

func NewClient(cfg Config) (*Client, error) {
	if cfg.Address == "" {
		cfg.Address = os.Getenv("VAULT_ADDR")
	}
	if cfg.Token == "" {
		cfg.Token = os.Getenv("VAULT_TOKEN")
	}

	if cfg.Address == "" || cfg.Token == "" {
		return nil, errors.New("vault address and token are required")
	}

	config := vault.DefaultConfig()
	config.Address = cfg.Address

	client, err := vault.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize vault client: %w", err)
	}

	client.SetToken(cfg.Token)

	return &Client{
		client: client,
	}, nil
}

func (c *Client) GetSecret(ctx context.Context, path string) (map[string]interface{}, error) {
	secret, err := c.client.KVv2("secret").Get(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("unable to read secret at path %s: %w", path, err)
	}

	if secret == nil || secret.Data == nil {
		return nil, fmt.Errorf("secret not found at path %s", path)
	}

	return secret.Data, nil
}

func (c *Client) GetExchangeCredentials(ctx context.Context, exchange string) (string, string, string, error) {
	path := fmt.Sprintf("exchanges/%s", exchange)
	secret, err := c.GetSecret(ctx, path)
	if err != nil {
		return "", "", "", err
	}

	apiKey, _ := secret["api_key"].(string)
	apiSecret, _ := secret["api_secret"].(string)
	passphrase, _ := secret["passphrase"].(string)

	return apiKey, apiSecret, passphrase, nil
}
