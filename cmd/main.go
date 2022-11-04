package main

import (
	"github.com/Mohsen-Khodabakhshi/gosta/services/config"
	"github.com/Mohsen-Khodabakhshi/gosta/services/server"
)

func init() {
	cfg, err := config.InitConfig("./build/config.yaml")

	if err != nil {
		panic("Error in load config.yaml")
	}

	config.SetPublicConfig(cfg)
}

func main() {
	cfg := config.Config

	server.InitMiddleware()

	server.RunServer(cfg.Server.Host, cfg.Server.Port)

}
