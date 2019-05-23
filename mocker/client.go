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

func (c *Client) AddHandler(req Request, resp Response) error {
	_, err := c.cl.AddHandler(context.TODO(),
		&proto.AddHandlerRequest{
			RequestMatcher: &proto.RequestMatcher{
				Method:  req.Method,
				Path:    req.Path,
				Query:   req.Query,
				Body:    req.Body,
				Headers: buildHeaderValues(req.Headers),
			},
			Response: &proto.MatcherResponse{
				Status: resp.Status,
				Body:   resp.Body,
			},
		})

	return err
}

func buildHeaderValues(h map[string][]string) map[string]*proto.HeaderValue {
	res := make(map[string]*proto.HeaderValue, len(h))
	for k, v := range h {
		res[k] = &proto.HeaderValue{Value: v}
	}
	return res
}
