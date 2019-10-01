package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/fedepaol/grpcsamples/pkg/movie"
	"google.golang.org/grpc"
)

func doClient(serverAddr string) {
	target := fmt.Sprintf("%s:%d", serverAddr, port)
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial to server %v", err)
	}
	client := movie.NewMoviesServiceClient(conn)
	id := &movie.MovieID{Mid: 1}

	fmt.Println("querying movie containing and")
	queryParams := &movie.MovieQueryParams{Query: "and"}
	stream, err := client.QueryMovie(context.Background(), queryParams)
	if err != nil {
		log.Fatalf("Failed to query movie %v, error %v", id, err)
	}

	for {
		movie, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Errir while streaming movie %v %v", queryParams, err)
		}

		log.Println("Streaming got movie ", movie)
	}
}
