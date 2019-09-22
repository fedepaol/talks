package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/fedepaol/grpcsamples/pkg/beer"
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

	fmt.Println("querying beers containing and")
	queryParams := &beer.BeerQueryParams{Query: "and"}
	stream, err := client.QueryBeer(context.Background(), queryParams)
	if err != nil {
		log.Fatalf("Failed to query beer %v, error %v", id, err)
	}

	for {
		beer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Errir while streaming beer %v %v", queryParams, err)
		}

		log.Println("Streaming got beer ", beer)
	}
}
