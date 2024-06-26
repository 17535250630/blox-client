package stream

import (
	"time"

	"github.com/17535250630/blox-client/provider"
)

type TraderHTTPPriceOpt func(s *traderHTTPPriceStream)

func WithTraderHTTPClient(h *provider.HTTPClient) TraderHTTPPriceOpt {
	return func(s *traderHTTPPriceStream) {
		s.h = h
	}
}

func WithTraderHTTPTicker(t *time.Ticker) TraderHTTPPriceOpt {
	return func(s *traderHTTPPriceStream) {
		s.ticker = t
	}
}

func WithTraderHTTPMint(m string) TraderHTTPPriceOpt {
	return func(s *traderHTTPPriceStream) {
		s.mint = m
	}
}
