package main

import (
	"flag"
	"strings"
)

const port = 20001

func main() {
	isServer := flag.Bool("server", false, "tells wether the component is client side or server side")
	serverAddresses := flag.String("target", "", "the address of the grpc server in case of client")
	isFaulty := flag.Bool("faulty", false, "tells if the server is faulty")
	flag.Parse()

	if *isServer {
		startServer(*isFaulty)
	} else {
		targets := strings.Split(*serverAddresses, ",")
		doClient(targets)
	}
}
