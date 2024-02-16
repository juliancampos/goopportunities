package main

import (
	"github.com/juliancampos/goopportunities/config"
	"github.com/juliancampos/goopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("goopportunities")

	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}
	router.Initialize()
}
