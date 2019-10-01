package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fedepaol/grpcsamples/pkg/movie"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func doClient(serverAddr string) {
	target := fmt.Sprintf("%s:%d", serverAddr, port)
	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithUnaryInterceptor(retryInterceptor))
	if err != nil {
		log.Fatalf("Failed to dial to server %v", err)
	}
	client := movie.NewMoviesServiceClient(conn)
	id := &movie.MovieID{Mid: 1}

	res, err := client.GetMovie(context.Background(), id)
	if err != nil {
		st := status.Convert(err)
		for _, detail := range st.Details() {
			switch t := detail.(type) {
			case *movie.MovieError:
				log.Printf("Received type safe error %v", t)
			}
		}
		log.Fatalf("Failed to get movie %v, error %v", id, err)
	}

	log.Println("Got movie ", res)
}

func retryInterceptor(ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {
	var lastErr error
	for attempt := uint(0); attempt < 10; attempt++ {
		fmt.Println("Attempt #", attempt)
		lastErr = invoker(ctx, method, req, reply, cc, opts...)
		if lastErr == nil {
			return nil
		}
	}
	return lastErr
}
