package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Listening on port 8080")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}