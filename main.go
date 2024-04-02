package main

import (
	"flag"
)

func main() {
	var (
		jsonAddr = flag.String("json", ":3000", "listen address of the json transport")
		grpcAddr = flag.String("grpc", ":4000", "listen address of the gRPC transport")
	)
	flag.Parse()

	svc := NewLoggingService(&priceFetcher{})

	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()

	go makeGRPCServerAndRun(*grpcAddr, svc)
}
