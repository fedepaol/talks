package main

import (
	"log"
	"net"
	"strconv"

	"github.com/fedepaol/grpcsamples/pkg/beer"
	"google.golang.org/grpc"
)

func startServer() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	beer.RegisterBeersServiceServer(server, beer.NewServer())

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startMiddle(serverAddr string) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	beer.RegisterBeersServiceServer(server, newMiddle(serverAddr, port))

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
