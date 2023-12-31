package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next PriceService
}

func NewLogginService(next PriceService) PriceService {

	return &loggingService{
		next: next,
	}
}

func (s loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		reqID := ctx.Value("requestID")

		logrus.WithFields(logrus.Fields{
			"requestID": reqID,
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
			"ticker":    ticker,
		}).Info("FetchPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}
