package mocker

import (
	"google.golang.org/grpc"
)

func NewClient(endpoint string) (MockerClient, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err

	}
	cl := NewMockerClient(conn)

	return cl, nil
}
