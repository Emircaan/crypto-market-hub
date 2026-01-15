package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Emircaan/crypto-service/internal/generated/proto/cryptomarket"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.CryptoMarketServiceClient
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewCryptoMarketServiceClient(conn)

	return &Client{
		conn:   conn,
		client: client,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) API() pb.CryptoMarketServiceClient {
	return c.client
}
