package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fedepaol/grpcsamples/pkg/movie"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func doClient(serverAddr string) {
	target := fmt.Sprintf("%s:%d", serverAddr, port)

	var kap = keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithKeepaliveParams(kap))
	if err != nil {
		log.Fatalf("Failed to dial to server %v", err)
	}
	client := movie.NewMoviesServiceClient(conn)
	id := &movie.MovieID{Mid: 1}

	movie, err := client.GetMovie(context.Background(), id)
	if err != nil {
		log.Fatalf("Failed to get movie %v, error %v", id, err)
	}

	log.Println("Got movie ", movie)
}
