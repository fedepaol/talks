package main

import (
	"log"
	"net"
	"strconv"

	"github.com/fedepaol/grpcsamples/pkg/movie"
	"google.golang.org/grpc"
)

func startServer() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	movie.RegisterMoviesServiceServer(server, NewServer())

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
