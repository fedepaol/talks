package movie

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MovieServer is a Movie server implementation
type MovieServer struct {
	Movies map[int]Movie
}

func (b *MovieServer) GetMovie(ctx context.Context, id *MovieID) (*Movie, error) {
	MovieID := id.GetMid()
	Movie, ok := b.Movies[int(MovieID)]
	if !ok {
		return nil, status.Error(codes.NotFound, "Movie not found")
	}
	return &Movie, nil
}

func (b *MovieServer) QueryMovie(p *MovieQueryParams, s MoviesService_QueryMovieServer) error {
	for _, b := range b.Movies {
		if strings.Contains(b.MovieDescription, p.GetQuery()) {
			s.Send(&b)
		}
	}
	return nil
}

// NewServer returns an instance of the server
func NewServer() MoviesServiceServer {
	return &MovieServer{Movies: MovieMap}
}
