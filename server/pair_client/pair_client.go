package pairclient

import (
	"encoding/json"

	"github.com/Ouroborus-Org/ouroborus-route-server/server/ctx"
	"github.com/Ouroborus-Org/ouroborus-route-server/server/models"
	"github.com/gorilla/websocket"
)

func RunPairClient(pairServerUrl string, serverCtx *ctx.ServerContext) error {
	ws, _, err := websocket.DefaultDialer.Dial(pairServerUrl, nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			return err
		}
		var updates models.Updates
		err = json.Unmarshal(msg, &updates)
		if err != nil {
			return err
		}
		for addr, token := range updates.Tokens {
			serverCtx.Tokens[addr] = token
		}
		for addr, pair := range updates.Pairs {
			serverCtx.Pairs[addr] = pair
		}
	}
}
