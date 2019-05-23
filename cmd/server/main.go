package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/bosonic-code/mock-api/mocker"
	"google.golang.org/grpc"
)

var (
	port     = os.Getenv("API_PORT")
	grpcPort = os.Getenv("GRPC_PORT")
)

func main() {
	mockServer := &MockServer{}
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", "localhost", port),
		Handler: mockServer,
	}

	log.Printf("HTTP server listening at %v", port)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Unable to start HTTP server %v", err)
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))

	if err != nil {
		log.Fatalf("GRPC server failed to listen: %v", err)
	}

	log.Printf("GRPC server listening at %v", grpcPort)

	s := grpc.NewServer()
	mocker.RegisterMockerServer(s, mockServer)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("GRPC server failed to serve: %v", err)
	}
}
