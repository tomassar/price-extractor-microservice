package main

import (
	"context"
	"math/rand"
	"net"

	"github.com/tomassar/crypto-price-fetcher-microservice/proto"
	"google.golang.org/grpc"
)

func makeGRPCServerAndRun(listenAddr string, svc PriceFetcher) error {
	grpcPriceFetcher := NewGRPCPriceFetcherService(svc)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)

	proto.RegisterPriceFetcherServer(server, grpcPriceFetcher)
	return server.Serve(ln)
}

type GRPCPriceFetcherServer struct {
	svc PriceFetcher
	proto.UnimplementedPriceFetcherServer
}

func NewGRPCPriceFetcherService(svc PriceFetcher) *GRPCPriceFetcherServer {
	return &GRPCPriceFetcherServer{
		svc: svc,
	}
}

func (s *GRPCPriceFetcherServer) FetchPrice(ctx context.Context, req *proto.PriceRequest) (*proto.PriceResponse, error) {
	reqid := rand.Intn(10000)
	ctx = context.WithValue(ctx, "requestID", reqid)

	price, err := s.svc.FetchPrice(ctx, req.GetTicker())
	if err != nil {
		return nil, err
	}

	resp := &proto.PriceResponse{
		Ticker: req.GetTicker(),
		Price:  float32(price),
	}

	return resp, nil
}
