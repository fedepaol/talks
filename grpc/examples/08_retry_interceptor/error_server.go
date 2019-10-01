package main

import (
	"context"
	"fmt"

	"github.com/fedepaol/grpcsamples/pkg/movie"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MovieServer is a movie server implementation
type errorMovieServer struct {
	numErrors int
}

func (s *errorMovieServer) GetMovie(_ context.Context, id *movie.MovieID) (*movie.Movie, error) {

	movieErr := movie.MovieError{
		Reason:  movie.MovieErrorCode_MOVIE_NOT_FOUND,
		Message: fmt.Sprintf("Could not find movie %v", *id),
	}
	st, _ := status.New(codes.NotFound, "not found").WithDetails(&movieErr)

	if s.numErrors < 3 {
		s.numErrors++
		fmt.Println("Forcing error")
		return nil, st.Err()
	}

	s.numErrors = 0
	movieID := id.GetMid()
	b, ok := movie.MovieMap[int(movieID)]
	if !ok {
		return nil, st.Err()
	}
	return &b, nil
}

func (s *errorMovieServer) QueryMovie(_ *movie.MovieQueryParams, _ movie.MoviesService_QueryMovieServer) error {
	panic("not implemented")
}

// NewServer returns an instance of the server
func NewServer() movie.MoviesServiceServer {
	return &errorMovieServer{}
}
