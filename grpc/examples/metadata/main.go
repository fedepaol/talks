package main

import "flag"

const port = 20001

func main() {
	isServer := flag.Bool("server", false, "tells wether the component is client side or server side")
	serverAddress := flag.String("target", "", "the address of the grpc server in case of client")

	flag.Parse()

	if *isServer {
		startServer()
	} else {
		doClient(*serverAddress)
	}
}
