package main

import (
	"seed/backend/apiserver"
	"seed/backend/cfg"
	"seed/backend/internal/log"
)

func main() {
	//init log
	opts := log.NewOptions()
	log.Init(opts)
	defer log.Flush()

	//init config
	cfg.Init()

	seedServer := apiserver.NewSeedServer()
	seedServer.Run()
}
