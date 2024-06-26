package actor

import (
	"time"

	"github.com/17535250630/blox-client/provider"
)

type JupiterOpt func(s *jupiterSwap)

func WithJupiterTokenPair(inputMint, outputMint string) JupiterOpt {
	return func(s *jupiterSwap) {
		s.inputMint = inputMint
		s.outputMint = outputMint
	}
}

func WithJupiterAmount(amount float64) JupiterOpt {
	return func(s *jupiterSwap) {
		s.amount = amount
	}
}

func WithJupiterInitialTimeout(initialTimeout time.Duration) JupiterOpt {
	return func(s *jupiterSwap) {
		s.initialTimeout = initialTimeout
	}
}

func WithJupiterAfterTimeout(timeout time.Duration) JupiterOpt {
	return func(s *jupiterSwap) {
		s.afterTimeout = timeout
	}
}

func WithJupiterInterval(interval time.Duration) JupiterOpt {
	return func(s *jupiterSwap) {
		s.interval = interval
	}
}

func WithJupiterClient(client *provider.HTTPClient) JupiterOpt {
	return func(s *jupiterSwap) {
		s.client = client
	}
}

func WithJupiterPublicKey(pk string) JupiterOpt {
	return func(s *jupiterSwap) {
		s.publicKey = pk
	}
}

func WithJupiterAlternate(alternate bool) JupiterOpt {
	return func(s *jupiterSwap) {
		s.alternate = alternate
	}
}
