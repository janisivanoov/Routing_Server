package models

import "github.com/Ouroborus-Org/ouroborus-route-server/server/bigint"

type Pair struct {
	Reserve0 bigint.BigInt `json:"reserve0"`
	Reserve1 bigint.BigInt `json:"reserve1"`
	Token0   string        `json:"token0"`
	Token1   string        `json:"token1"`
	Fees     interface{}   `json:"fees"`
}

type Token struct {
	Decimals bigint.BigInt `json:"decimals"`
	Symbol   string        `json:"symbol"`
	Name     string        `json:"name"`
}

type Updates struct {
	Pairs  map[string]Pair  `json:"pairs"`
	Tokens map[string]Token `json:"tokens"`
}
