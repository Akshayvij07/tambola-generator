package main

import (
	"log"

	"github.com/Akshayvij07/thambola-generator/pkg/config"
	dd "github.com/Akshayvij07/thambola-generator/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := dd.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
