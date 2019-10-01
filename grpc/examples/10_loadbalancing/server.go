package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/fedepaol/grpcsamples/pkg/movie"
	"google.golang.org/grpc"
)

func startServer(isFaulty bool) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	movie.RegisterMoviesServiceServer(server, newServer(isFaulty))

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// MovieServer is a movie server implementation
type movieServer struct {
	mustFail       bool
	lastWasSuccess bool
}

func (b *movieServer) GetMovie(ctx context.Context, id *movie.MovieID) (*movie.Movie, error) {

	if b.mustFail {
		fmt.Println("I am the faulty one")
	} else {
		fmt.Println("I am not the faulty one")
	}
	if b.mustFail && b.lastWasSuccess {
		fmt.Println("Sending error")

		return nil, fmt.Errorf("Failure faulty")
	}
	movie := movie.Movie{
		MovieName:        "Celebration Ale",
		MovieDescription: "The long, cold nights of winter are a little brighter with Celebration Ale. Wonderfully robust and rich, Celebration Ale is dry-hopped for a lively, intense aroma. Brewed especially for the holidays, it is perfect for a festive gathering or for a quiet evening at home.",
		MovieStyle:       "American IPA",
		Mid:             1,
	}
	b.lastWasSuccess = true
	return &movie, nil
}

func (b *movieServer) QueryMovie(p *movie.MovieQueryParams, s movie.MoviesService_QueryMovieServer) error {
	panic("not implemented")
	return nil
}

// NewServer returns an instance of the server
func newServer(isFaulty bool) movie.MoviesServiceServer {
	return &movieServer{mustFail: isFaulty}
}
