package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fedepaol/grpcsamples/pkg/beer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func doClient(serverAddr []string) {
	targetsWithPort := make([]string, len(serverAddr))
	for i, t := range serverAddr {
		targetsWithPort[i] = fmt.Sprintf("%s:%d", t, port)
	}
	targetsWithPort = append(targetsWithPort, "foo:1234") // appending a fake one to see what happens

	resolver.Register(&exampleResolverBuilder{targets: targetsWithPort})
	resolver.SetDefaultScheme("myResolverScheme")

	// since we are using a resolver the url we pass does not matter
	conn, err := grpc.Dial("foo:1234", grpc.WithInsecure(), grpc.WithBalancerName("round_robin"))
	if err != nil {
		log.Fatalf("Failed to dial to server %v", err)
	}
	client := beer.NewBeersServiceClient(conn)
	id := &beer.BeerID{Bid: 1}

	for i := 0; i < 10; i++ {
		res, err := client.GetBeer(context.Background(), id)
		if err != nil {
			log.Println("Got err ", err)
		} else {
			log.Println("Got beer ", res)
		}
		time.Sleep(1 * time.Second)
	}

}

type exampleResolverBuilder struct {
	targets []string
}

// No matter what the target is, we always return the two servers we have
func (b *exampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	r := &exampleResolver{targets: b.targets, cc: cc}
	r.start()
	return r, nil
}
func (*exampleResolverBuilder) Scheme() string { return "myResolverScheme" }

type exampleResolver struct {
	targets []string
	cc      resolver.ClientConn
}

func (r *exampleResolver) start() {
	addrs := make([]resolver.Address, len(r.targets))
	for i, s := range r.targets {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}
func (*exampleResolver) ResolveNow(o resolver.ResolveNowOption) {}
func (*exampleResolver) Close()                                 {}
