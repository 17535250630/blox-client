package utils

import "github.com/urfave/cli/v2"

var (
	OutputFileFlag = &cli.StringFlag{
		Name:     "output",
		Usage:    "file to output CSV results to",
		Required: true,
	}
	SolanaHTTPRPCEndpointFlag = &cli.StringFlag{
		Name:  "solana-http-endpoint",
		Usage: "HTTP RPC server endpoint to make blockchain queries against",
		Value: "http://34.203.186.197:8899",
	}
	SolanaWSRPCEndpointFlag = &cli.StringFlag{
		Name:  "solana-ws-endpoint",
		Usage: "WS RPC server endpoint to make blockchain pub/sub queries against",
		Value: "ws://34.203.186.197:8900/ws",
	}
)
