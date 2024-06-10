package stream

import "github.com/17535250630/blox-client/provider"

type TraderWSPriceOpt func(s *tradeWSPrice)

func WithTraderWSClient(w *provider.WSClient) TraderWSPriceOpt {
	return func(s *tradeWSPrice) {
		s.w = w
	}
}

func WithTraderWSMint(m string) TraderWSPriceOpt {
	return func(s *tradeWSPrice) {
		s.mint = m
	}
}
