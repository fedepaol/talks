package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fedepaol/grpc/examples/pkg/beer"
	"google.golang.org/grpc"
)

func doClient(serverAddr string) {
	target := fmt.Sprintf("%s:%d", serverAddr, port)
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial to server %v", err)
	}
	client := beer.NewBeersServiceClient(conn)
	id := &beer.BeerID{Bid: 1}

	beer, err := client.GetBeer(context.Background(), id)
	if err != nil {
		log.Fatalf("Failed to get beer %v, error %v", id, err)
	}

	log.Println("Got beer ", beer)
}
