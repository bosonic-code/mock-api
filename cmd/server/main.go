package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/bosonic-code/mock-api/internal/proto"
	"google.golang.org/grpc"
)

var (
	port    = os.Getenv("API_PORT")
	comPort = os.Getenv("COM_PORT")
)

func main() {
	mockServer := &Server{}
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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", comPort))

	if err != nil {
		log.Fatalf("Mockserver failed to listen: %v", err)
	}

	log.Printf("Mockserver listening at %v", comPort)

	s := grpc.NewServer()
	proto.RegisterMockerServer(s, mockServer)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Mockserver failed to serve: %v", err)
	}
}
