package main

import (
	"context"
	"fmt"

	"github.com/fedepaol/grpcsamples/pkg/beer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BeerServer is a beer server implementation
type errorBeerServer struct {
}

func (s *errorBeerServer) GetBeer(_ context.Context, id *beer.BeerID) (*beer.Beer, error) {
	beerErr := beer.BeerError{
		Reason:  beer.BeerErrorCode_BEER_NOT_FOUND,
		Message: fmt.Sprintf("Could not find beer %v", *id),
	}
	st, _ := status.New(codes.NotFound, "not found").WithDetails(&beerErr)
	return nil, st.Err()
}

func (s *errorBeerServer) QueryBeer(_ *beer.BeerQueryParams, _ beer.BeersService_QueryBeerServer) error {
	panic("not implemented")
}

// NewServer returns an instance of the server
func NewServer() beer.BeersServiceServer {
	return &errorBeerServer{}
}
