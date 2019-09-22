package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fedepaol/grpcsamples/pkg/beer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func doClient(serverAddr string) {
	target := fmt.Sprintf("%s:%d", serverAddr, port)
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial to server %v", err)
	}
	client := beer.NewBeersServiceClient(conn)
	id := &beer.BeerID{Bid: 1}

	res, err := client.GetBeer(context.Background(), id)
	if err != nil {
		st := status.Convert(err)
		for _, detail := range st.Details() {
			switch t := detail.(type) {
			case *beer.BeerError:
				log.Println("Received type safe error %v", t)
			}
		}
		log.Fatalf("Failed to get beer %v, error %v", id, err)
	}

	log.Println("Got beer ", res)
}
