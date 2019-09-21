package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/fedepaol/grpcsamples/pkg/beer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BeerServer is a beer server implementation
type BeerServer struct {
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

func (b *BeerServer) GetBeer(ctx context.Context, id *beer.BeerID) (*beer.Beer, error) {
	beerID := id.GetBid()
	beer, ok := b.beers[int(beerID)]
	if !ok {
		return nil, status.Error(codes.NotFound, "beer not found")
	}
	return &beer, nil
}

func (b *BeerServer) QueryBeer(p *beer.BeerQueryParams, s beer.BeersService_QueryBeerServer) error {
	for _, b := range b.beers {
		if strings.Contains(b.BeerDescription, p.GetQuery()) {
			s.Send(&b)
		}
	}
	return nil
}

func startServer() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	beer.RegisterBeersServiceServer(server, &BeerServer{beers: beers})

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
