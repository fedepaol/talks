package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/fedepaol/grpcsamples/pkg/beer"
	"google.golang.org/grpc"
)

func startServer(isFaulty bool) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	beer.RegisterBeersServiceServer(server, newServer(isFaulty))

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// BeerServer is a beer server implementation
type beerServer struct {
	mustFail       bool
	lastWasSuccess bool
}

func (b *beerServer) GetBeer(ctx context.Context, id *beer.BeerID) (*beer.Beer, error) {

	if b.mustFail {
		fmt.Println("I am the faulty one")
	} else {
		fmt.Println("I am not the faulty one")
	}
	if b.mustFail && b.lastWasSuccess {
		fmt.Println("Sending error")

		return nil, fmt.Errorf("Failure faulty")
	}
	beer := beer.Beer{
		BeerName:        "Celebration Ale",
		BeerDescription: "The long, cold nights of winter are a little brighter with Celebration Ale. Wonderfully robust and rich, Celebration Ale is dry-hopped for a lively, intense aroma. Brewed especially for the holidays, it is perfect for a festive gathering or for a quiet evening at home.",
		BeerStyle:       "American IPA",
		Bid:             1,
	}
	b.lastWasSuccess = true
	return &beer, nil
}

func (b *beerServer) QueryBeer(p *beer.BeerQueryParams, s beer.BeersService_QueryBeerServer) error {
	panic("not implemented")
	return nil
}

// NewServer returns an instance of the server
func newServer(isFaulty bool) beer.BeersServiceServer {
	return &beerServer{mustFail: isFaulty}
}
