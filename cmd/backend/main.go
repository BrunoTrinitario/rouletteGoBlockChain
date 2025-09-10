package main

import (
	"github.com/BrunoTrinintario/rouletteGoBlockChain/internal/api"
)

func main() {
	router := api.SetupRouter()
	router.Run(":8080") // levanta la API
}
