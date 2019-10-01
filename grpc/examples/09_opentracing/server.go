package main

import (
	"log"
	"net"
	"strconv"

	"github.com/fedepaol/grpcsamples/pkg/movie"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"

	"google.golang.org/grpc"
)

func startServer() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)))
	movie.RegisterMoviesServiceServer(server, movie.NewServer())

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

	server := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)))

	movie.RegisterMoviesServiceServer(server, newMiddle(serverAddr, port))

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
