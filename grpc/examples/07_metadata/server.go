package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/fedepaol/grpcsamples/pkg/movie"
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
	movie.RegisterMoviesServiceServer(server, newServer())

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type movieServer struct {
	movies map[int]movie.Movie
}

func (b *movieServer) GetMovie(ctx context.Context, id *movie.MovieID) (*movie.Movie, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("Received metadata ")
		key := md.Get("key")
		if len(key) > 0 {
			fmt.Println("Value for key ", key[0])
		}
	}
	movieID := id.GetMid()
	movie, ok := b.movies[int(movieID)]
	if !ok {
		return nil, status.Error(codes.NotFound, "movie not found")
	}
	return &movie, nil
}

func (b *movieServer) QueryMovie(p *movie.MovieQueryParams, s movie.MoviesService_QueryMovieServer) error {
	panic("not implemented")
}

// NewServer returns an instance of the server
func newServer() movie.MoviesServiceServer {
	return &movieServer{movies: movie.MovieMap}
}
