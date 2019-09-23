package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/fedepaol/grpcsamples/pkg/beer"
	"google.golang.org/grpc"
)

func doClient(serverAddr string) {
	target := fmt.Sprintf("%s:%d", serverAddr, port)
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial to server %v", err)
	}
	client := beer.NewBeersServiceClient(conn)
	id := &beer.BeerID{Bid: 1}

	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}

	wg.Add((1))
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Canceling")
		cancel()
		wg.Done()
	}()

	res, err := client.GetBeer(ctx, id)
	if err != nil {
		log.Fatalf("Failed to get beer %v, error %v", id, err)
	}

	wg.Wait()
	log.Println("Got beer ", res)
}
