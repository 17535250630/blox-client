# Solana Trader Golang Client

## Objective
This SDK is designed to make it easy for you to use the bloXroute Labs Solana Trader API
in Go. 

## Installation
```
go get github.com/17535250630/blox-client
```

## Usage

This library supports HTTP, websockets, and GRPC interfaces. You must use websockets or GRPC for any streaming methods, 
but any simple request/response calls are universally supported.

For any methods involving transaction creation you will need to provide your Solana private key. You can provide this 
via the environment variable `PRIVATE_KEY`, or specify it via the provider configuration if you want to load it with
some other mechanism. See samples for more information. As a general note on this: methods named `Post*` (e.g. 
`PostOrder`) typically do not sign/submit the transaction, only return the raw unsigned transaction. This isn't 
very useful to most users (unless you want to write a signer in a different language), and you'll typically want the 
similarly named `Submit*` methods (e.g. `SubmitOrder`). These methods generate, sign, and submit the
transaction all at once.

You will also need your bloXroute authorization header to use these endpoints. By default, this is loaded from the 
`AUTH_HEADER` environment variable.

## Quickstart

### Request sample:

```go
package main

import (
	"context"
	"fmt"
	"github.com/17535250630/blox-client/provider"
	pb "github.com/bloXroute-Labs/solana-trader-proto/api"
)

func main() {
	// GPRC
	g, err := provider.NewGRPCClient()
	if err != nil {
		panic(err)
	}

	orderbook, err := g.GetOrderbook(context.Background(), "ETH/USDT", 5, pb.Project_P_OPENBOOK) // in this case limit to 5 bids and asks. 0 for no limit
	if err != nil {
		panic(err)
	}
	fmt.Println(orderbook)

	// HTTP
	h := provider.NewHTTPClient()
	tickers, err := h.GetTickers(context.Background(), "ETHUSDT", pb.Project_P_OPENBOOK)
	if err != nil {
		panic(err)
	}
	fmt.Println(tickers)
	
	// WS
	w, err := provider.NewWSClient()
	if err != nil {
		panic(err)
	}
	// note that open orders is a slow function call
	openOrders, err := w.GetOpenOrders(context.Background(), "ETH/USDT", "4raJjCwLLqw8TciQXYruDEF4YhDkGwoEnwnAdwJSjcgv", "", pb.Project_P_OPENBOOK)
	if err != nil {
		panic(err)
	}
	fmt.Println(openOrders)
}

```
#### Stream (only in GRPC/WS):

```go
package main

import (
	"fmt"
	"github.com/17535250630/blox-client/provider"
	pb "github.com/bloXroute-Labs/solana-trader-proto/api"
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, err := provider.NewGRPCClient() // replace this with `NewWSClient()` to use WebSockets
	if err != nil {
		panic(err)
	}

	stream, err := g.GetOrderbookStream(ctx, []string{"SOL/USDT"}, 5, pb.Project_P_OPENBOOK)
	if err != nil {
		panic(err)
	}
	
	// wrap result in channel for easy of use
	orderbookCh := make(chan *pb.GetOrderbooksStreamResponse)
	stream.Into(orderbookCh)
	for i := 0; i < 3; i++ {
		orderbook := <-orderbookCh
		fmt.Println(orderbook)
	}
}
```

More code samples are provided in the `examples/` directory.

**A quick note on market names:**
You can use a couple of different formats, with restrictions: 
1. `A/B` (only for GRPC/WS clients) --> `ETH/USDT`
2. `A:B` --> `ETH:USDT`
3. `A-B` --> `ETH-USDT`
4. `AB` --> `ETHUSDT`


## Development

Unit tests:

```
$ make unit
```

Integration tests per provider:
```
$ make grpc-examples

$ make http-examples

$ make ws-examples
```