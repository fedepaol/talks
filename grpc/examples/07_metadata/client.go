package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fedepaol/grpcsamples/pkg/beer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func doClient(serverAddr string) {
	target := fmt.Sprintf("%s:%d", serverAddr, port)
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial to server %v", err)
	}
	client := beer.NewBeersServiceClient(conn)
	id := &beer.BeerID{Bid: 1}

	md := metadata.Pairs("key", "value")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	beer, err := client.GetBeer(ctx, id)
	if err != nil {
		log.Fatalf("Failed to get beer %v, error %v", id, err)
	}

	log.Println("Got beer ", beer)
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
