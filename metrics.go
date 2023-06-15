package main

import "context"

type metricsService struct {
	next PriceService
}

func NewMetricService(next PriceService) PriceService {
	return &metricsService{
		next: next,
	}
}

func (s *metricsService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {

	return s.next.FetchPrice(ctx, ticker)
}
