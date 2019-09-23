package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fedepaol/grpcsamples/pkg/beer"
	"google.golang.org/grpc"
)

// BeerServer is a beer server implementation
type middleServer struct {
	client beer.BeersServiceClient
}

func (b *middleServer) GetBeer(ctx context.Context, id *beer.BeerID) (*beer.Beer, error) {
	fmt.Println("Received request")
	beer, err := b.client.GetBeer(ctx, id)

	if err != nil {
		return nil, err
	}
	fmt.Println("Sending response")
	return beer, nil
}

func (b *middleServer) QueryBeer(p *beer.BeerQueryParams, s beer.BeersService_QueryBeerServer) error {
	panic("not implemented")
}

// NewMiddle returns an instance of the middle server
func newMiddle(serverAddr string, port int) beer.BeersServiceServer {
	target := fmt.Sprintf("%s:%d", serverAddr, port)
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial to server %v", err)
	}
	client := beer.NewBeersServiceClient(conn)
	return &middleServer{client: client}
}
