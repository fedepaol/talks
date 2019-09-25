package main

import (
	"context"
	"fmt"

	"github.com/fedepaol/grpcsamples/pkg/beer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

// BeerServer is a beer server implementation
type errorBeerServer struct {
	numErrors int
}

func (s *errorBeerServer) GetBeer(_ context.Context, id *beer.BeerID) (*beer.Beer, error) {

	beerErr := beer.BeerError{
		Reason:  beer.BeerErrorCode_BEER_NOT_FOUND,
		Message: fmt.Sprintf("Could not find beer %v", *id),
	}
	st, _ := status.New(codes.NotFound, "not found").WithDetails(&beerErr)

	if s.numErrors < 3 {
		s.numErrors++
		fmt.Println("Forcing error")
		return nil, st.Err()
	}

	s.numErrors = 0
	beerID := id.GetBid()
	b, ok := beers[int(beerID)]
	if !ok {
		return nil, st.Err()
	}
	return &b, nil
}

func (s *errorBeerServer) QueryBeer(_ *beer.BeerQueryParams, _ beer.BeersService_QueryBeerServer) error {
	panic("not implemented")
}

// NewServer returns an instance of the server
func NewServer() beer.BeersServiceServer {
	return &errorBeerServer{}
}
