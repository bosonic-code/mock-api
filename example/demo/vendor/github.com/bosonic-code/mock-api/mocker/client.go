package mocker

import (
	"context"

	"github.com/bosonic-code/mock-api/internal/proto"
	"google.golang.org/grpc"
)

type Client struct {
	cl proto.MockerClient
}

func Create(endpoint string) (*Client, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err

	}
	cl := proto.NewMockerClient(conn)

	return &Client{
		cl: cl,
	}, nil
}

func (c *Client) Handle(req Request, resp Response) error {
	_, err := c.cl.Handle(context.TODO(), &proto.HandleRequest{
		Request: &proto.Request{
			Method: req.Method,
			Path:   req.Path,
			Query:  req.Query,
		},
		Response: &proto.Response{
			Status: resp.Status,
			Body:   resp.Body,
		},
	})

	return err
}
