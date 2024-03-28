package main

import (
	"context"
	"fmt"
)

// Interface that can fetch a price
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// Implements PriceFetcher interface
type priceFetcher struct {
}

func (s priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC": 60_000.0,
	"ETH": 600.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("Ticket (%s) is not supported", ticker)
	}

	return price, nil
}
