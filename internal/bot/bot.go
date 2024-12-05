package bot

import (
	"discord-automod/internal/bot/cogs"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Discord *discordgo.Session
}

var filterFilePath = "config/banned_words.txt"

func NewBot(token string) (*Bot, error) {
	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		return nil, fmt.Errorf("failed to create Discord session: %w", err)
	}

	return &Bot{Discord: discord}, nil
}

func (b *Bot) Start() error {
	err := cogs.InitFilter(filterFilePath)
	if err != nil {
		return err
	}
	b.Discord.AddHandler(cogs.NewMessage)

	err = b.Discord.Open()
	if err != nil {
		return fmt.Errorf("failed to open Discord session: %w", err)
	}
	defer b.Stop()

	log.Println("Bot is running... Press CTRL+C to exit.")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Shutting down bot gracefully...")
	return nil
}

func (b *Bot) Stop() {
	if b.Discord != nil {
		b.Discord.Close()
	}
}
