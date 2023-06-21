package main

import (
	"context"
	"fmt"
)

type PriceService interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceService struct {
}

func (s *priceService) FetchPrice(_ context.Context, ticker string) (float64, error) {
	price, ok := priceMocker[ticker]

	if !ok {
		return 0.0, fmt.Errorf("price for ticker (%s) is not avaliable", ticker)
	}

	return price, nil
}

var priceMocker = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
	"MSR": 100_00.0,
}
