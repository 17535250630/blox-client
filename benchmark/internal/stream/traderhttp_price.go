package stream

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/17535250630/blox-client/benchmark/internal/logger"
	"github.com/17535250630/blox-client/provider"
	pb "github.com/bloXroute-Labs/solana-trader-proto/api"
	"go.uber.org/zap"
)

type traderHTTPPriceStream struct {
	h        *provider.HTTPClient
	mint     string
	ticker   *time.Ticker
	interval time.Duration
}

func NewTraderHTTPPriceStream(opts ...TraderHTTPPriceOpt) (Source[DurationUpdate[*pb.GetPriceResponse], QuoteResult], error) {
	s := &traderHTTPPriceStream{
		h:        provider.NewHTTPClient(),
		interval: defaultInterval,
	}

	for _, o := range opts {
		o(s)
	}

	if s.mint == "" {
		return nil, errors.New("mint is mandatory")
	}

	return s, nil
}

func (s traderHTTPPriceStream) log() *zap.SugaredLogger {
	return logger.Log().With("source", "traderapi/http")
}

func (s traderHTTPPriceStream) Name() string {
	return fmt.Sprintf("traderapi")
}

// Run stops when parent ctx is canceled
func (s traderHTTPPriceStream) Run(parent context.Context) ([]RawUpdate[DurationUpdate[*pb.GetPriceResponse]], error) {
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	ticker := s.ticker
	if ticker == nil {
		ticker = time.NewTicker(s.interval)
	}

	return collectOrderedUpdates(ctx, ticker, func() (*pb.GetPriceResponse, error) {
		res, err := s.h.GetPrice(ctx, []string{s.mint})
		if err != nil {
			return nil, err
		}

		filteredRes := &pb.GetPriceResponse{TokenPrices: nil}
		for _, price := range res.TokenPrices {
			if price.Project == pb.Project_P_JUPITER {
				filteredRes.TokenPrices = append(filteredRes.TokenPrices, price)
			}
		}
		return filteredRes, nil
	}, &pb.GetPriceResponse{}, func(err error) {
		s.log().Errorw("could not fetch price", "err", err)
	})
}

func (s traderHTTPPriceStream) Process(updates []RawUpdate[DurationUpdate[*pb.GetPriceResponse]], removeDuplicates bool) (results map[int][]ProcessedUpdate[QuoteResult], duplicates map[int][]ProcessedUpdate[QuoteResult], err error) {
	results = make(map[int][]ProcessedUpdate[QuoteResult])
	duplicates = make(map[int][]ProcessedUpdate[QuoteResult])

	lastBuyPrice := -1.
	lastSellPrice := -1.
	slot := -1 // no slot info is available
	for _, update := range updates {
		buyPrice := update.Data.Data.TokenPrices[0].Buy
		sellPrice := update.Data.Data.TokenPrices[0].Sell

		qr := QuoteResult{
			Elapsed:   update.Timestamp.Sub(update.Data.Start),
			BuyPrice:  buyPrice,
			SellPrice: sellPrice,
			Source:    "traderHTTP",
		}
		pu := ProcessedUpdate[QuoteResult]{
			Timestamp: update.Timestamp,
			Slot:      slot,
			Data:      qr,
		}

		if buyPrice == lastBuyPrice && sellPrice == lastSellPrice {
			duplicates[slot] = append(duplicates[slot], pu)
			if removeDuplicates {
				continue
			}
		}

		lastBuyPrice = buyPrice
		lastSellPrice = sellPrice
		results[slot] = append(results[slot], pu)
	}

	return
}
