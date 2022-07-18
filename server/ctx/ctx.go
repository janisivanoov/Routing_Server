package ctx

import "github.com/Ouroborus-Org/ouroborus-route-server/server/models"

type ServerContext struct {
	Pairs  map[string]models.Pair
	Tokens map[string]models.Token
}

func NewServerContext() *ServerContext {
	return &ServerContext{
		Pairs:  make(map[string]models.Pair),
		Tokens: make(map[string]models.Token),
	}
}
