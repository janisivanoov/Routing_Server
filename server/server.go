package server
import (
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/Ouroborus-Org/ouroborus-route-server/server/ctx"
	"github.com/Ouroborus-Org/ouroborus-route-server/server/handlers"
	pairclient "github.com/Ouroborus-Org/ouroborus-route-server/server/pair_client"
	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	ListenUrl     string
	PairServerUrl string
}

func quitOnInterrupt(quit chan<- error) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	quit <- nil
}

func runGin(cfg *ServerConfig, serverCtx *ctx.ServerContext) error {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/tokens", handlers.GetTokensBuilder(serverCtx))
	r.GET("/quote", handlers.GetQuoteBuilder(serverCtx))
	return r.Run(cfg.ListenUrl)
}

func RunServer(cfg *ServerConfig) error {
	quit := make(chan error, 1)

	go quitOnInterrupt(quit)

	serverCtx := ctx.NewServerContext()

	go func() {
		if err := runGin(cfg, serverCtx); err != nil {
			quit <- errors.New(fmt.Sprintf("gin exited with error: %s", err))
		}
	}()

	go func() {
		if err := pairclient.RunPairClient(cfg.PairServerUrl, serverCtx); err != nil {
			quit <- errors.New(fmt.Sprintf("pairclient exited with error: %s", err))
		}
	}()

	return <-quit
}
