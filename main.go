package main

import (
	"log"

	"github.com/Ouroborus-Org/ouroborus-route-server/server"
)

func main() {
	cfg := &server.ServerConfig{
		PairServerUrl: "ws://127.0.0.1:8080",
		ListenUrl:     ":8090",
	}
	if err := server.RunServer(cfg); err != nil {
		log.Fatal(err)
	}
}
