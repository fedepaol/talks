package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	res, err := client.GetMovie(ctx, id)
	if err != nil {
		log.Fatalf("Failed to get movie %v, error %v", id, err)
	}

	log.Println("Got movie ", res)
}
