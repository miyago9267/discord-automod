package main

import (
	"discord-automod/config"
	"discord-automod/internal/bot"
	"log"
)

func main() {
	config.InitConfig()

	botToken := config.Cfg.Token

	b, err := bot.NewBot(botToken)
	if err != nil {
		log.Fatalf("Failed to initialize bot: %v", err)
	}

	err = b.Start()
	if err != nil {
		log.Fatalf("Error running bot: %v", err)
	}
}
