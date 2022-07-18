package handlers

import (
	"fmt"
	"net/http"

	"github.com/Ouroborus-Org/ouroborus-route-server/server/bigint"
	"github.com/Ouroborus-Org/ouroborus-route-server/server/ctx"
	"github.com/Ouroborus-Org/ouroborus-route-server/server/routing"
	"github.com/gin-gonic/gin"
)

func GetTokensBuilder(serverCtx *ctx.ServerContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, serverCtx.Tokens)
	}
}

func checkHasToken(token string, serverCtx *ctx.ServerContext, ctx *gin.Context) (ok bool) {
	if _, ok = serverCtx.Tokens[token]; !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERR",
			"message": fmt.Sprintf("no such token: %s", token),
		})
	}
	return
}

func GetQuoteBuilder(serverCtx *ctx.ServerContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		tokenIn := ctx.Query("tokenIn")
		tokenOut := ctx.Query("tokenOut")

		if tokenIn == tokenOut {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "ERR",
				"message": "input token should not equal output token",
			})
			return
		}
		if !checkHasToken(tokenIn, serverCtx, ctx) {
			return
		}
		if !checkHasToken(tokenOut, serverCtx, ctx) {
			return
		}

		amountIn, err := bigint.FromString(ctx.Query("amountIn"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "ERR",
				"message": err.Error(),
			})
			return
		}

		p := &routing.RouteParams{
			TokenIn:  tokenIn,
			TokenOut: tokenOut,
			AmountIn: amountIn,
		}
		route, err := routing.GetRoute(serverCtx, p)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "ERR",
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"result": route,
		})
	}
}
