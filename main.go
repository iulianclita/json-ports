package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/iulianclita/json-ports/app"
)

func main() {
	var cfg app.Config
	if err := env.Parse(&cfg); err != nil {
		panic(fmt.Sprintf("failed to parse app config: %v", err))
	}

	webApp := app.New(cfg)

	webApp.Start()
	defer webApp.Stop()
}
