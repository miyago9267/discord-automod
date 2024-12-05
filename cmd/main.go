package main

import (
	"discord-automod/config"
	"fmt"
)

func main() {
	config.InitConfig()

	fmt.Println("Hello, World!")
	fmt.Println("Config:", config.Cfg.Debug)
	fmt.Println("Config:", config.Cfg.Token)
}
