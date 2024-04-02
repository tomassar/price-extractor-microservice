package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/tomassar/crypto-price-fetcher-microservice/client"
	"github.com/tomassar/crypto-price-fetcher-microservice/proto"
)

func main() {
	var (
		jsonAddr = flag.String("json", ":3000", "listen address of the json transport")
		grpcAddr = flag.String("grpc", ":4000", "listen address of the gRPC transport")
		svc      = NewLoggingService(&priceFetcher{})
		ctx      = context.Background()
	)
	flag.Parse()

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		time.Sleep(3 * time.Second)
		resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{
			Ticker: "ETH",
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", resp)
	}()

	go makeGRPCServerAndRun(*grpcAddr, svc)

	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()

}
