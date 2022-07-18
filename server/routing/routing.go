package routing

import (
	"errors"
	"math/big"

	"github.com/Ouroborus-Org/ouroborus-route-server/server/bigint"
	"github.com/Ouroborus-Org/ouroborus-route-server/server/ctx"
	"github.com/Ouroborus-Org/ouroborus-route-server/server/models"
)

type RouteParams struct {
	TokenIn  string
	TokenOut string
	AmountIn *bigint.BigInt
	GasPrice *bigint.BigInt // later
}

type Route struct {
	AmountOut *bigint.BigInt `json:"amountOut"`
	GasUsed   *bigint.BigInt `json:"gasUsed"`
}

type routeSegment struct {
	fromToken *string
	pair      *models.Pair
	amountOut *big.Int
	gasUsed   *big.Int
}

func (rs *routeSegment) toRoute() *Route {
	return &Route{
		AmountOut: bigint.Wrap(rs.amountOut),
		GasUsed:   bigint.Wrap(rs.gasUsed),
	}
}

func GetRoute(serverCtx *ctx.ServerContext, p *RouteParams) (*Route, error) {
	q := []string{p.TokenIn}
	s := make(map[string]bool)
	r := make(map[string]routeSegment)
	r[p.TokenIn] = routeSegment{nil, nil, p.AmountIn.Unwrap(), big.NewInt(0)}
	// maximizing each token route
	for len(q) > 0 {
		tokenIn := q[0]
		delete(s, tokenIn)
		q = q[1:]
		tail := r[tokenIn]
		for _, pair := range serverCtx.Pairs {
			tokenOut := pair.Token1
			if pair.Token1 == tokenIn {
				tokenOut = pair.Token0
			} else if pair.Token0 != tokenIn {
				continue
			}
			amountOut, gasUsed := pair.GetAmountOut(tokenOut, tail.amountOut)
			if prev, ok := r[tokenOut]; !ok || prev.amountOut.Cmp(amountOut) < 0 {
				r[tokenOut] = routeSegment{
					fromToken: &tokenIn,
					pair:      &pair,
					amountOut: amountOut,
					gasUsed:   gasUsed.Add(gasUsed, tail.gasUsed),
				}
				if !s[tokenOut] {
					s[tokenOut] = true
					q = append(q, tokenOut)
				}
			}
		}
	}
	// q is empty, all routes are maxed out
	rs, ok := r[p.TokenOut]
	if !ok {
		return nil, errors.New("route not found")
	}
	return rs.toRoute(), nil
}
