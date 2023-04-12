package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   PricingService
}

func NewLoggingMiddleware(logger log.Logger, next PricingService) (lmw *loggingMiddleware) {
	lmw = &loggingMiddleware{
		logger: logger,
		next:   next,
	}

	return
}

func (mw loggingMiddleware) GetRetailTotal(code string, qty int) (total float64, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetRetailTotal",
			"code", code,
			"quantity", qty,
			"total", total,
			"error", err,
			"duration", time.Since(begin),
		)
	}(time.Now())

	total, err = mw.next.GetRetailTotal(code, qty)

	return
}

func (mw loggingMiddleware) GetWholesaleTotal(partner string, code string, qty int) (total float64, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetWholesaleTotal",
			"partner", partner,
			"code", code,
			"quantity", qty,
			"total", total,
			"error", err,
			"duration", time.Since(begin),
		)
	}(time.Now())
	total, err = mw.next.GetWholesaleTotal(partner, code, qty)

	return
}

func LogTotalRetailPriceEndpoint(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("endpoint", "TotalRetailPriceEndpoint", "msg", "Calling endpoint")
			defer logger.Log("endpoint", "TotalRetailPriceEndpoint", "msg", "Called endpoint")

			return next(ctx, request)
		}
	}
}

func LogTotalWholesalePriceEndpoint(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("endpoint", "TotalWholesalePriceEndpoint", "msg", "Calling endpoint")
			defer logger.Log("endpoint", "TotalWholesalePriceEndpoint", "msg", "Called endpoint")

			return next(ctx, request)
		}
	}
}
