package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/fedepaol/grpcsamples/pkg/beer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func startServer() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	beer.RegisterBeersServiceServer(server, newServer())

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type beerServer struct {
	beers map[int]beer.Beer
}

var beers = map[int]beer.Beer{
	1: beer.Beer{
		BeerName:        "Celebration Ale",
		BeerDescription: "The long, cold nights of winter are a little brighter with Celebration Ale. Wonderfully robust and rich, Celebration Ale is dry-hopped for a lively, intense aroma. Brewed especially for the holidays, it is perfect for a festive gathering or for a quiet evening at home.",
		BeerStyle:       "American IPA",
		Bid:             1,
	},
	2: beer.Beer{
		BeerName:        "Punk IPA",
		BeerDescription: "Welcome to a post Punk apocalyptic mother fucker of a pale ale. A beer that spent its formative years Blitzkrieg bopping around India and the sub continent. Quintessential Empire with an anarchic twist. God save the Queen and all who sail in her. Raising a Stiff Little Finger to IPAs that have come before and those it is yet to meet. Turn up the volume Pay the man. Embrace the punked up, fucked up outlaw elite. Never Mind the Bollocks this is the real shit.",
		BeerStyle:       "American IPA",
		Bid:             2,
	},
}

func (b *beerServer) GetBeer(ctx context.Context, id *beer.BeerID) (*beer.Beer, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("Received metadata ")
		key := md.Get("key")
		if len(key) > 0 {
			fmt.Println("Value for key ", key[0])
		}
	}
	beerID := id.GetBid()
	beer, ok := b.beers[int(beerID)]
	if !ok {
		return nil, status.Error(codes.NotFound, "beer not found")
	}
	return &beer, nil
}

func (b *beerServer) QueryBeer(p *beer.BeerQueryParams, s beer.BeersService_QueryBeerServer) error {
	panic("not implemented")
}

// NewServer returns an instance of the server
func newServer() beer.BeersServiceServer {
	return &beerServer{beers: beers}
}
