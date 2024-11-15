package client

import (
	"context"
	"time"

	pb "github.com/SystemAlgoFund/grpc_package/proto"
	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.ServiceClient
}

func NewClient(address string) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	c := pb.NewServiceClient(conn)
	return &Client{conn: conn, client: c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) SendData(data []byte) (*pb.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.Request{Message: data}
	return c.client.Send(ctx, req)
}
