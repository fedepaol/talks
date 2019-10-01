package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fedepaol/grpcsamples/pkg/movie"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func doClient(serverAddr string) {
	target := fmt.Sprintf("%s:%d", serverAddr, port)
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial to server %v", err)
	}
	client := movie.NewMoviesServiceClient(conn)
	id := &movie.MovieID{Mid: 1}

	md := metadata.Pairs("key", "value")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	movie, err := client.GetMovie(ctx, id)
	if err != nil {
		log.Fatalf("Failed to get movie %v, error %v", id, err)
	}

	log.Println("Got movie ", movie)
}

func retryInterceptor(ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {
	var lastErr error
	for attempt := uint(0); attempt < 10; attempt++ {
		fmt.Println("Client invoking server")
		lastErr = invoker(ctx, method, req, reply, cc, opts...)
		if lastErr == nil {
			return nil
		}
		fmt.Println("Client invoking server failed")
	}
	return lastErr
}
