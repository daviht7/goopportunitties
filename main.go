package main

import (
	"github.com/daviht7/goopportunitties/config"
	"github.com/daviht7/goopportunitties/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	//Initialize Configs
	err := config.Init()

	if err != nil {
		logger.Error(err)
		return
	}

	//Iniciando as routas
	router.Initialize()

}
