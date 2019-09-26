package main

import (
	"flag"
	"fmt"
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

const port = 20001

var tracer opentracing.Tracer

func main() {
	isServer := flag.Bool("server", false, "tells wether the component is client side or server side")
	serverAddress := flag.String("target", "", "the address of the grpc server in case of client")
	isMiddle := flag.Bool("middle", false, "tells wether the component is middle server")
	flag.Parse()

	var name string
	switch {
	case *isMiddle:
		name = "middle"
	case *isServer:
		name = "server"
	default:
		name = "client"
	}

	var closer io.Closer
	tracer, closer = initJaeger(name)
	defer closer.Close()

	if *isServer {
		if !*isMiddle {
			startServer()
		} else {
			startMiddle(*serverAddress)
		}
	} else {
		doClient(*serverAddress)
	}
}

// initJaeger returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: "http://jaeger:14268/api/traces",
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
