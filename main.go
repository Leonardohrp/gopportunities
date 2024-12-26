package main

import (
	"github.com/Leonardohrp/gopportunities/config"
	"github.com/Leonardohrp/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", "err")
		return
	}

	router.Initialize()
}
