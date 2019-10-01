package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fedepaol/grpcsamples/pkg/movie"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"google.golang.org/grpc"
)

// MovieServer is a movie server implementation
type middleServer struct {
	client movie.MoviesServiceClient
}

func (b *middleServer) GetMovie(ctx context.Context, id *movie.MovieID) (*movie.Movie, error) {
	fmt.Println("Received request")
	time.Sleep(5 * time.Second)
	fmt.Println("Sending request")

	if ctx.Err() == context.Canceled { // this is useful to check even before sending the request
		fmt.Println("Is Canceled")
	}

	movie, err := b.client.GetMovie(ctx, id)

	if err != nil {
		fmt.Println("Got error ", err)
		return nil, err
	}
	fmt.Println("Sending response")
	return movie, nil
}

func (b *middleServer) QueryMovie(p *movie.MovieQueryParams, s movie.MoviesService_QueryMovieServer) error {
	panic("not implemented")
}

// NewMiddle returns an instance of the middle server
func newMiddle(serverAddr string, port int) movie.MoviesServiceServer {
	target := fmt.Sprintf("%s:%d", serverAddr, port)
	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)))
	if err != nil {
		log.Fatalf("Failed to dial to server %v", err)
	}
	client := movie.NewMoviesServiceClient(conn)
	return &middleServer{client: client}
}
